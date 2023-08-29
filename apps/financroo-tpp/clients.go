package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	cdrBank "github.com/cloudentity/openbanking-quickstart/generated/cdr/client"
	cdrModels "github.com/cloudentity/openbanking-quickstart/generated/cdr/client/banking"
	fdxBank "github.com/cloudentity/openbanking-quickstart/generated/fdx/client"
	obbrPayments "github.com/cloudentity/openbanking-quickstart/generated/obbr/payments/client"
	obukAccounts "github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/client"
	"github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/models"
	payments_client "github.com/cloudentity/openbanking-quickstart/generated/obuk/payments/client"
	"github.com/cloudentity/openbanking-quickstart/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"

	obbrAccounts "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/client"

	"gopkg.in/square/go-jose.v2"
)

type Clients struct {
	AcpAccountsClients map[BankID]acpclient.Client
	AcpPaymentsClients map[BankID]acpclient.Client
	BankClients        map[BankID]BankClient
	ConsentClients     map[BankID]ConsentClient

	SignatureVerificationKey map[BankID]jose.JSONWebKey
}

func (c *Clients) GetConsentClient(id BankID) (ConsentClient, error) {
	if client, ok := c.ConsentClients[id]; ok {
		return client, nil
	}

	return nil, fmt.Errorf("consent client not configured for bank %s", id)
}

func (c *Clients) GetAccountsClient(id BankID) (acpclient.Client, error) {
	if client, ok := c.AcpAccountsClients[id]; ok {
		return client, nil
	}

	return acpclient.Client{}, fmt.Errorf("acp accounts client not configured for bank %s", id)
}

func (c *Clients) GetPaymentsClient(id BankID) (acpclient.Client, error) {
	if client, ok := c.AcpPaymentsClients[id]; ok {
		return client, nil
	}

	return acpclient.Client{}, fmt.Errorf("acp payments client not configured for bank %s", id)
}

func (c *Clients) GetBankClient(id BankID) (BankClient, error) {
	if client, ok := c.BankClients[id]; ok {
		return client, nil
	}

	return nil, fmt.Errorf("bank client not configured for bank %s", id)
}

type BankClient interface {
	GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error)
	GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error)
	GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error)
	CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error)
}

type BankClientCreationFn func(BankConfig) (BankClient, error)

type ConsentClient interface {
	CreateConsentExplicitly() bool
	CreateAccountConsent(c *gin.Context) (string, error)
	CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error)
	GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error)

	UsePAR() bool
	DoPAR(c *gin.Context) (string, acpclient.CSRF, error)
	Signer
}

func InitClients(
	config Config,
	signerCreateFn SignerCreationFn,
	bankClientCreateFn BankClientCreationFn,
	consentClientCreateFn ConsentClientCreationFn,
) (Clients, error) {
	var (
		clients = Clients{
			AcpAccountsClients:       make(map[BankID]acpclient.Client),
			AcpPaymentsClients:       make(map[BankID]acpclient.Client),
			BankClients:              make(map[BankID]BankClient),
			ConsentClients:           make(map[BankID]ConsentClient),
			SignatureVerificationKey: make(map[BankID]jose.JSONWebKey),
		}
		acpAccountsWebClient acpclient.Client
		acpPaymentsWebClient acpclient.Client
		signer               Signer
		err                  error
	)

	for _, b := range config.Banks {
		if acpAccountsWebClient, err = NewAcpClient(config, b, "/api/callback"); err != nil {
			return clients, errors.Wrapf(err, "failed to create acp accounts client")
		}

		clients.AcpAccountsClients[b.ID] = acpAccountsWebClient

		if acpPaymentsWebClient, err = NewAcpClient(config, b, "/api/domestic/callback"); err != nil {
			return clients, errors.Wrapf(err, "failed to create acp payments client")
		}

		clients.AcpPaymentsClients[b.ID] = acpPaymentsWebClient

		if signerCreateFn != nil {
			if signer, err = signerCreateFn(config.KeyFile); err != nil {
				return clients, errors.Wrapf(err, "failed to create consent message signer for %s", config.Spec)
			}
		}

		if clients.BankClients[b.ID], err = bankClientCreateFn(b); err != nil {
			return clients, errors.Wrapf(err, "failed to create bank client for %s", config.Spec)
		}

		if consentClientCreateFn != nil {
			clients.ConsentClients[b.ID] = consentClientCreateFn(acpAccountsWebClient, acpPaymentsWebClient, signer)
		}

		if clients.SignatureVerificationKey[b.ID], err = utils.GetServerKey(&acpPaymentsWebClient, utils.SIG); err != nil {
			return clients, errors.Wrapf(err, "failed to get signature verification key for %s", config.Spec)
		}
	}

	return clients, nil
}

func NewAcpClient(config Config, bankConfig BankConfig, redirect string) (acpclient.Client, error) {
	var (
		authorizeURL, issuerURL, redirectURL *url.URL
		client                               acpclient.Client
		err                                  error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", bankConfig.ACPInternalURL, bankConfig.Tenant, bankConfig.Server)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", bankConfig.ACPURL, bankConfig.Tenant, bankConfig.Server)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s%s", config.UIURL, redirect)); err != nil {
		return client, err
	}

	requestObjectExpiration := time.Minute * 10
	acpConfig := acpclient.Config{
		ClientID:                      bankConfig.ClientID,
		IssuerURL:                     issuerURL,
		AuthorizeURL:                  authorizeURL,
		RedirectURL:                   redirectURL,
		RequestObjectSigningKeyFile:   config.KeyFile,
		RequestObjectExpiration:       &requestObjectExpiration,
		Scopes:                        config.ClientScopes,
		Timeout:                       time.Second * 5,
		CertFile:                      config.CertFile,
		KeyFile:                       config.KeyFile,
		RootCA:                        config.RootCA,
		ClientAssertionSigningKeyFile: config.AssertionSigningKeyFile,
		ClientAssertionSigningAlg:     config.AssertionSigningAlg,
	}

	if config.Spec == CDR {
		acpConfig.SkipClientCredentialsAuthn = true
		acpConfig.AuthMethod = acpclient.PrivateKeyJwtAuthnMethod
	}

	if config.Spec == FDX {
		acpConfig.SkipClientCredentialsAuthn = true
		acpConfig.AuthMethod = acpclient.TLSClientAuthnMethod
	}

	if config.Spec == GENERIC {
		acpConfig.SkipClientCredentialsAuthn = true
	}

	if client, err = acpclient.New(acpConfig); err != nil {
		return client, err
	}

	return client, nil
}

func RenewAccountsToken(ctx context.Context, bank ConnectedBank, client acpclient.Client) (*oauth2Models.TokenResponse, error) {
	var (
		resp      oauth2.TokenOK
		request   *http.Request
		response  *http.Response
		assertion string
		body      []byte
		err       error
	)

	values := url.Values{
		"client_id":     {client.Config.ClientID},
		"grant_type":    {"refresh_token"},
		"refresh_token": {bank.RefreshToken},
	}

	if client.Config.AuthMethod == acpclient.ClientSecretPostAuthnMethod && client.Config.ClientSecret != "" {
		values.Add("client_secret", client.Config.ClientSecret)
	}

	if client.Config.AuthMethod == acpclient.PrivateKeyJwtAuthnMethod {
		if assertion, err = client.GenerateClientAssertion(); err != nil {
			return nil, err
		}
		values.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
		values.Add("client_assertion", assertion)
	}

	if request, err = http.NewRequest(http.MethodPost, client.Config.GetTokenURL(), strings.NewReader(values.Encode())); err != nil {
		return nil, errors.Wrapf(err, "failed to create token request")
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if response, err = client.DoRequest(request); err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &resp.Payload); err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

type OBUKClient struct {
	*obukAccounts.Accounts
	*payments_client.Payments
}

func NewOBUKClient(config BankConfig) (BankClient, error) {
	var (
		c   = &OBUKClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		hc,
	)

	c.Accounts = obukAccounts.New(tr, nil)
	c.Payments = payments_client.New(tr, nil)

	return c, nil
}

type CDRClient struct {
	*cdrBank.Banking
}

func NewCDRClient(config BankConfig) (BankClient, error) {
	var (
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return nil, err
	}

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		http.DefaultClient,
	)
	return &CDRClient{
		cdrBank.New(tr, nil),
	}, nil
}

type OBBRClient struct {
	*obbrAccounts.Accounts
	*obbrPayments.Payments
}

func NewOBBRClient(config BankConfig) (BankClient, error) {
	var (
		c   = &OBBRClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Accounts = obbrAccounts.New(NewHTTPRuntimeWithClient(
		u.Host,
		u.Path+"/accounts/v1",
		[]string{u.Scheme},
		hc,
	), nil)
	c.Payments = obbrPayments.New(NewHTTPRuntimeWithClient(
		u.Host,
		u.Path+"/payments/v1",
		[]string{u.Scheme},
		hc,
	), nil)

	return c, nil
}

type ConsentClientCreationFn func(acpclient.Client, acpclient.Client, Signer) ConsentClient

type OBUKConsentClient struct {
	Accounts acpclient.Client
	Payments acpclient.Client
	Signer
}

func NewOBUKConsentClient(accountsClient, paymentsClient acpclient.Client, signer Signer) ConsentClient {
	return &OBUKConsentClient{accountsClient, paymentsClient, signer}
}

type OBBRConsentClient struct {
	Accounts acpclient.Client
	Payments acpclient.Client
	Signer
}

func NewOBBRConsentClient(accountsClient, paymentsClient acpclient.Client, signer Signer) ConsentClient {
	return &OBBRConsentClient{accountsClient, paymentsClient, signer}
}

type FDXBankClient struct {
	*fdxBank.Client
}

func NewFDXBankClient(config BankConfig) (BankClient, error) {
	var (
		c   = &FDXBankClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Client = fdxBank.New(NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		hc,
	), nil)

	return c, nil
}

type CDRConsentClient struct {
	ClientID     string
	ClientSecret string
	PublicClient acpclient.Client
}

func NewCDRConsentClient(publicClient, clientCredentialsClient acpclient.Client, _ Signer) ConsentClient {
	return &CDRConsentClient{
		ClientID:     clientCredentialsClient.Config.ClientID,
		ClientSecret: clientCredentialsClient.Config.ClientSecret,
		PublicClient: publicClient,
	}
}

func (c *CDRConsentClient) CreateConsentExplicitly() bool {
	return false
}

func (c *CDRConsentClient) UsePAR() bool {
	return true
}

func (c *CDRConsentClient) DoPAR(ctx *gin.Context) (string, acpclient.CSRF, error) {
	var (
		csrf acpclient.CSRF
		resp acpclient.PARResponse
		err  error
	)

	if resp, csrf, err = c.PublicClient.DoPAR(
		acpclient.WithResponseType("code id_token"),
		acpclient.WithPKCE(),
		acpclient.WithOpenbankingACR([]string{"urn:cds.au:cdr:3"}),
	); err != nil {
		return "", acpclient.CSRF{}, err
	}
	return resp.RequestURI, csrf, err
}

func (c *CDRConsentClient) CreateAccountConsent(ctx *gin.Context) (string, error) {
	return "", nil
}

func (c *CDRConsentClient) DoRequestObjectEncryption() bool {
	return false
}

func (c *CDRConsentClient) GetPaymentConsent(ctx *gin.Context, consentID string) (interface{}, error) {
	return "", nil
}

func (c *CDRConsentClient) CreatePaymentConsent(ctx *gin.Context, req CreatePaymentRequest) (string, error) {
	return "", nil
}

func (c *CDRConsentClient) Sign([]byte) (string, error) {
	return "", nil
}

type GenericBankClient struct {
	*cdrBank.Banking
}

func NewGenericBankClient(config BankConfig) (BankClient, error) {
	var (
		c   = &GenericBankClient{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		http.DefaultClient,
	)
	return &GenericBankClient{
		cdrBank.New(tr, nil),
	}, nil
}

var _ BankClient = &GenericBankClient{}

func (c *GenericBankClient) GetAccounts(ctx *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error) {
	var (
		resp         *cdrModels.ListAccountsOK
		accountsData = []Account{}
		err          error
	)

	if resp, err = c.Banking.Banking.ListAccounts(
		cdrModels.NewListAccountsParamsWithContext(ctx).
			WithDefaults(),
		runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
			return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}),
	); err != nil {
		return accountsData, err
	}

	for _, a := range resp.Payload.Data.Accounts {
		accountsData = append(accountsData, Account{
			OBAccount6: models.OBAccount6{
				AccountID: (*models.AccountID)(a.AccountID),
				Nickname:  models.Nickname(a.Nickname),
				Account: []*models.OBAccount6AccountItems0{
					{
						Name:           models.Name0(*a.AccountID),
						Identification: (*models.Identification0)(a.MaskedNumber),
					},
				},
			},
			BankID: bank.BankID,
		})
	}

	return accountsData, err
}

func (c *GenericBankClient) GetTransactions(ctx *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error) {
	var (
		resp             *cdrModels.GetTransactionsOK
		accounts         []Account
		transactionsData []Transaction
		err              error
	)

	if accounts, err = c.GetAccounts(ctx, accessToken, bank); err != nil {
		return transactionsData, errors.Wrap(err, "failed to get account ids for transactions")
	}

	for _, account := range accounts {
		if resp, err = c.Banking.Banking.GetTransactions(
			cdrModels.NewGetTransactionsParams().
				WithDefaults().
				WithAccountID(string(*account.AccountID)),
			runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
				return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", accessToken))
			}),
		); err != nil {
			return transactionsData, err
		}

		for _, cdrTransaction := range resp.Payload.Data.Transactions {
			if transaction, err := cdrTransactionToInternalTransaction(cdrTransaction, bank); err != nil {
				logrus.Infof("failed to map cdr transaction to internal transaction: %+v", err)
			} else {
				transactionsData = append(transactionsData, transaction)
			}
		}
	}

	return transactionsData, nil
}

func (c *GenericBankClient) GetBalances(ctx *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	var (
		resp         *cdrModels.ListBalancesBulkOK
		balancesData []Balance
		err          error
	)

	if resp, err = c.Banking.Banking.ListBalancesBulk(
		cdrModels.NewListBalancesBulkParams().
			WithDefaults(),
		runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
			return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}),
	); err != nil {
		return []Balance{}, err
	}

	for _, balance := range resp.Payload.Data.Balances {
		balancesData = append(balancesData, Balance{
			AccountID: *balance.AccountID,
			Amount:    *balance.AvailableBalance,
			Currency:  balance.Currency,
			BankID:    bank.BankID,
		})
	}

	return balancesData, nil
}

func (c *GenericBankClient) CreatePayment(ctx *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	return PaymentCreated{}, errors.New("not implemented")
}

type GenericConsentClient struct {
	PublicClient acpclient.Client
}

var _ ConsentClient = &GenericConsentClient{}

func NewGenericConsentClient(publicClient, clientCredentialsClient acpclient.Client, _ Signer) ConsentClient {
	return &GenericConsentClient{
		PublicClient: publicClient,
	}
}

func (c *GenericConsentClient) CreateConsentExplicitly() bool {
	return false
}

func (c *GenericConsentClient) UsePAR() bool {
	return true
}

func (c *GenericConsentClient) DoPAR(ctx *gin.Context) (string, acpclient.CSRF, error) {
	var (
		csrf acpclient.CSRF
		resp acpclient.PARResponse
		err  error
	)

	if resp, csrf, err = c.PublicClient.DoPAR(
		acpclient.WithResponseType("code"),
		acpclient.WithPKCE(),
		acpclient.WithOpenbankingACR([]string{"generic:acr:3"}),
		acpclient.WithResponseMode("jwt"),
	); err != nil {
		return "", acpclient.CSRF{}, err
	}
	return resp.RequestURI, csrf, err
}

func (c *GenericConsentClient) CreateAccountConsent(ctx *gin.Context) (string, error) {
	return "", errors.New("not implemented")
}

func (c *GenericConsentClient) DoRequestObjectEncryption() bool {
	return false
}

func (c *GenericConsentClient) GetPaymentConsent(ctx *gin.Context, consentID string) (interface{}, error) {
	return "", errors.New("not implemented")
}

func (c *GenericConsentClient) CreatePaymentConsent(ctx *gin.Context, req CreatePaymentRequest) (string, error) {
	return "", errors.New("not implemented")
}

func (c *GenericConsentClient) Sign([]byte) (string, error) {
	return "", errors.New("not implemented")
}
