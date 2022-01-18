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
	"github.com/cloudentity/acp-client-go/client/oauth2"
	"github.com/cloudentity/acp-client-go/models"
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

type ConsentClient interface {
	CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error)
	GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error)
	CreateAccountConsent(c *gin.Context) (string, error)
	Signer
}

func (c *Clients) RenewAccountsToken(ctx context.Context, bank ConnectedBank) (*models.TokenResponse, error) {
	var (
		resp *oauth2.TokenOK
		err  error
	)

	if resp, err = c.AcpAccountsClient.Acp.Oauth2.Token(
		oauth2.NewTokenParamsWithContext(ctx).
			WithAid(c.AcpAccountsClient.ServerID).
			WithTid(c.AcpAccountsClient.TenantID).
			WithClientID(&c.AcpAccountsClient.Config.ClientID).
			WithGrantType("refresh_token").
			WithRefreshToken(&bank.RefreshToken)); err != nil {
		return nil, errors.Wrapf(err, "can't renew access token for a bank: %s", bank.BankID)
	}

	return resp.Payload, nil
}

func InitClients(config Config) (map[BankID]Clients, error) {
	var (
		clients              = map[BankID]Clients{}
		acpAccountsWebClient acpclient.Client
		acpPaymentsWebClient acpclient.Client
		bankClient           BankClient
		consentClient        ConsentClient
		signer               Signer
		err                  error
	)

	for _, bank := range config.Banks {
		if acpAccountsWebClient, err = NewAcpClient(config, bank, "/api/callback"); err != nil {
			return clients, errors.Wrapf(err, "failed to init acp web client for bank: %s", bank.ID)
		}

		if acpPaymentsWebClient, err = NewAcpClient(config, bank, "/api/domestic/callback"); err != nil {
			return clients, errors.Wrapf(err, "failed to init acp web client for bank: %s", bank.ID)
		}

		switch bank.BankType {
		case "obuk":
			if signer, err = NewOBUKSigner(config.KeyFile); err != nil {
				return clients, errors.Wrapf(err, "failed to init consent message signer for oguk bank: %s", bank.ID)
			}
			consentClient = &OBUKConsentClient{acpAccountsWebClient, acpPaymentsWebClient, signer}
			if bankClient, err = NewOBUKClient(bank); err != nil {
				return clients,
					errors.Wrapf(err, "failed to init client for obuk bank: %s", bank.ID)
			}
		case "obbr":
			if signer, err = NewOBBRSigner(config.KeyFile); err != nil {
				return clients, errors.Wrapf(err, "failed to init consent message signer for obbr bank: %s", bank.ID)
			}
			consentClient = &OBBRConsentClient{acpAccountsWebClient, acpPaymentsWebClient, signer}
			if bankClient, err = NewOBBRClient(bank); err != nil {
				return clients,
					errors.Wrapf(err, "failed to init client for obbr bank: %s", bank.ID)
			}
		}

		clients[bank.ID] = Clients{
			AcpAccountsClient: acpAccountsWebClient,
			AcpPaymentsClient: acpPaymentsWebClient,
			BankClient:        bankClient,
			ConsentClient:     consentClient,
		}
	}

	return clients, nil
}

func NewAcpClient(c Config, cfg BankConfig, redirect string) (acpclient.Client, error) {
	var (
		issuerURL, authorizeURL, redirectURL *url.URL
		client                               acpclient.Client
		err                                  error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, c.Tenant, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", c.ACPURL, c.Tenant, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s%s", c.UIURL, redirect)); err != nil {
		return client, err
	}

	requestObjectExpiration := time.Minute * 10
	config := acpclient.Config{
		ClientID:                    cfg.AcpClient.ClientID,
		IssuerURL:                   issuerURL,
		AuthorizeURL:                authorizeURL,
		RedirectURL:                 redirectURL,
		RequestObjectSigningKeyFile: cfg.AcpClient.KeyFile,
		RequestObjectExpiration:     &requestObjectExpiration,
		Scopes:                      cfg.AcpClient.Scopes,
		Timeout:                     cfg.AcpClient.Timeout,
		CertFile:                    cfg.AcpClient.CertFile,
		KeyFile:                     cfg.AcpClient.KeyFile,
		RootCA:                      cfg.AcpClient.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
}

func NewLoginClient(c Config) (acpclient.Client, error) {
	var (
		issuerURL *url.URL
		client    acpclient.Client
		err       error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, c.Tenant, c.Login.ServerID)); err != nil {
		return client, err
	}

	config := acpclient.Config{
		ClientID:  c.Login.ClientID,
		IssuerURL: issuerURL,
		Timeout:   c.Login.Timeout,
		RootCA:    c.Login.RootCA,
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

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		"/accounts/v1",
		[]string{u.Scheme},
		hc,
	)

	c.Accounts = obbrAccounts.New(tr, nil)
	c.PaymentConsentsBrasil = obbrPayments.New(tr, nil)

	return c, nil
}

type OBUKConsentClient struct {
	Accounts acpclient.Client
	Payments acpclient.Client
	Signer
}

type OBBRConsentClient struct {
	Accounts acpclient.Client
	Payments acpclient.Client
	Signer
}
