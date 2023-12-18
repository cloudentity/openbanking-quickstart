package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/generated/cdr/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type GenericBankClient struct {
	httpClient       *http.Client
	bankClientConfig BankClientConfig
	cc               clientcredentials.Config
}

func NewGenericBankClient(config Config) *GenericBankClient {
	var (
		pool  *x509.CertPool
		cert  tls.Certificate
		data  []byte
		certs []tls.Certificate
		err   error
	)

	if pool, err = x509.SystemCertPool(); err != nil {
		logrus.Fatalf("failed to read system root CAs %v", err)
	}

	if data, err = os.ReadFile(config.RootCA); err != nil {
		logrus.Fatalf("failed to read http client root ca: %v", err)
	}
	pool.AppendCertsFromPEM(data)

	if config.BankClientConfig.CertFile != "" && config.BankClientConfig.KeyFile != "" {
		if cert, err = tls.LoadX509KeyPair(config.BankClientConfig.CertFile, config.BankClientConfig.KeyFile); err != nil {
			logrus.Fatalf("failed to read certificate and private key %v", err)
		}
		certs = append(certs, cert)
	}

	return &GenericBankClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					RootCAs:      pool,
					MinVersion:   tls.VersionTLS12,
					Certificates: certs,
				},
			},
		},
		cc: clientcredentials.Config{
			ClientID:     config.BankClientConfig.ClientID,
			ClientSecret: config.BankClientConfig.ClientSecret,
			TokenURL:     config.BankClientConfig.TokenURL,
			Scopes:       config.BankClientConfig.Scopes,
		},
		bankClientConfig: config.BankClientConfig,
	}
}

func (c *GenericBankClient) GetInternalAccounts(ctx context.Context, id string) (InternalAccounts, error) {
	var (
		token                *oauth2.Token
		request              *http.Request
		response             *http.Response
		accountsEndpointPath string
		body                 []byte
		err                  error
	)

	if c.bankClientConfig.AccountsURL != nil {
		accountsEndpointPath = c.bankClientConfig.AccountsURL.String()
	} else {
		accountsEndpointPath = c.bankClientConfig.URL.String() + "/internal/accounts"
	}

	if token, err = c.cc.Token(context.WithValue(ctx, oauth2.HTTPClient, c.httpClient)); err != nil {
		return InternalAccounts{}, errors.Wrapf(err, "failed to get client credentials token for internal bank api call")
	}

	if request, err = http.NewRequestWithContext(ctx, http.MethodPost, accountsEndpointPath, strings.NewReader(
		url.Values{
			"customer_id": []string{id},
		}.Encode(),
	)); err != nil {
		return InternalAccounts{}, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Bearer "+token.AccessToken)

	if response, err = c.httpClient.Do(request); err != nil {
		return InternalAccounts{}, errors.Wrapf(err, "internal bank accounts api call failed")
	}
	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return InternalAccounts{}, errors.Wrap(err, "internal bank accounts api call failed")
	}

	if response.StatusCode >= http.StatusBadRequest {
		return InternalAccounts{}, errors.Wrap(errors.New(string(body)), "internal bank accounts api call failed")
	}

	return c.accountsResponseToInternalAccounts(body)
}

func (c *GenericBankClient) accountsResponseToInternalAccounts(body []byte) (accounts InternalAccounts, err error) {
	var accountListResponse models.ResponseBankingAccountList

	if err = json.Unmarshal(body, &accountListResponse); err != nil {
		return accounts, err
	}

	for _, acc := range accountListResponse.Data.Accounts {
		accounts.Accounts = append(accounts.Accounts, InternalAccount{
			ID:   *acc.AccountID,
			Name: acc.Nickname,
		})
	}

	return accounts, nil
}

func (c *GenericBankClient) GetInternalBalances(_ context.Context, _ string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
