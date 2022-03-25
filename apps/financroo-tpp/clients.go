package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client"
	obbrPayments "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/client"
	obc "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client"
	payments_client "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/client"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type Clients struct {
	AcpAccountsClient acpclient.Client
	AcpPaymentsClient acpclient.Client
	BankClient        BankClient
	ConsentClient     ConsentClient
}

type BankClient interface {
	GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error)
	GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error)
	GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error)
	CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error)
}

type BankClientCreationFn func(Config) (BankClient, error)

type ConsentClient interface {
	CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error)
	GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error)
	CreateAccountConsent(c *gin.Context) (string, error)
	Signer
}

func (c *Clients) RenewAccountsToken(ctx context.Context, bank ConnectedBank) (*oauth2Models.TokenResponse, error) {
	var (
		resp *oauth2.TokenOK
		err  error
	)

	if resp, err = c.AcpAccountsClient.Oauth2.Oauth2.Token(
		oauth2.NewTokenParamsWithContext(ctx).
			WithClientID(&c.AcpAccountsClient.Config.ClientID).
			WithGrantType("refresh_token").
			WithRefreshToken(&bank.RefreshToken)); err != nil {
		return nil, errors.Wrapf(err, "can't renew access token for a bank: %s", bank.BankID)
	}

	return resp.Payload, nil
}

func InitClients(config Config,
	signerCreateFn SignerCreationFn,
	bankClientCreateFn BankClientCreationFn,
	consentClientCreateFn ConsentClientCreationFn,
) (Clients, error) {
	var (
		clients              = Clients{}
		acpAccountsWebClient acpclient.Client
		acpPaymentsWebClient acpclient.Client
		bankClient           BankClient
		signer               Signer
		consentClient        ConsentClient
		err                  error
	)

	if acpAccountsWebClient, err = NewAcpClient(config, "/api/callback"); err != nil {
		return clients, errors.Wrapf(err, "failed to create acp accounts client")
	}

	if acpPaymentsWebClient, err = NewAcpClient(config, "/api/domestic/callback"); err != nil {
		return clients, errors.Wrapf(err, "failed to create acp payments client")
	}

	if signer, err = signerCreateFn(config.KeyFile); err != nil {
		return clients, errors.Wrapf(err, "failed to create consent message signer for %s", config.Spec)
	}

	if bankClient, err = bankClientCreateFn(config); err != nil {
		return clients, errors.Wrapf(err, "failed to create bank client for %s", config.Spec)
	}

	consentClient = consentClientCreateFn(acpAccountsWebClient, acpPaymentsWebClient, signer)

	return Clients{
		AcpAccountsClient: acpAccountsWebClient,
		AcpPaymentsClient: acpPaymentsWebClient,
		BankClient:        bankClient,
		ConsentClient:     consentClient,
	}, nil
}

func NewAcpClient(cfg Config, redirect string) (acpclient.Client, error) {
	var (
		authorizeURL, issuerURL, redirectURL *url.URL
		client                               acpclient.Client
		err                                  error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", cfg.ACPInternalURL, cfg.Tenant, cfg.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", cfg.ACPURL, cfg.Tenant, cfg.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s%s", cfg.UIURL, redirect)); err != nil {
		return client, err
	}

	requestObjectExpiration := time.Minute * 10
	config := acpclient.Config{
		ClientID:                    cfg.ClientID,
		IssuerURL:                   issuerURL,
		AuthorizeURL:                authorizeURL,
		RedirectURL:                 redirectURL,
		RequestObjectSigningKeyFile: cfg.KeyFile,
		RequestObjectExpiration:     &requestObjectExpiration,
		Scopes:                      cfg.ClientScopes,
		Timeout:                     time.Second * 5,
		CertFile:                    cfg.CertFile,
		KeyFile:                     cfg.KeyFile,
		RootCA:                      cfg.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
}

type OBUKClient struct {
	*obc.OpenbankingAccountsClient
	*payments_client.OpenbankingPaymentsClient
}

func NewOBUKClient(config Config) (BankClient, error) {
	var (
		c   = &OBUKClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.BankURL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		"/",
		[]string{u.Scheme},
		hc,
	)

	c.OpenbankingAccountsClient = obc.New(tr, nil)
	c.OpenbankingPaymentsClient = payments_client.New(tr, nil)

	return c, nil
}

type OBBRClient struct {
	*obbrAccounts.Accounts
	*obbrPayments.PaymentConsentsBrasil
}

func NewOBBRClient(config Config) (BankClient, error) {
	var (
		c   = &OBBRClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.BankURL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Accounts = obbrAccounts.New(NewHTTPRuntimeWithClient(
		u.Host,
		"/accounts/v1",
		[]string{u.Scheme},
		hc,
	), nil)
	c.PaymentConsentsBrasil = obbrPayments.New(NewHTTPRuntimeWithClient(
		u.Host,
		"/payments/v1",
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
