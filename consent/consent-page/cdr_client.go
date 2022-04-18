package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type CDRBankClient struct {
	httpClient *http.Client
	bankURL    string
	cc         clientcredentials.Config
}

func NewCDRBankClient(config Config) *CDRBankClient {
	var (
		pool  *x509.CertPool
		cert  tls.Certificate
		certs = []tls.Certificate{}
		data  []byte
		err   error
	)

	if cert, err = tls.LoadX509KeyPair(config.BankClientConfig.CertFile, config.BankClientConfig.KeyFile); err != nil {
		logrus.Fatalf("failed to read certificate and private key %v", err)
	}
	certs = append(certs, cert)

	if pool, err = x509.SystemCertPool(); err != nil {
		logrus.Fatalf("failed to read system root CAs %v", err)
	}

	if data, err = os.ReadFile(config.RootCA); err != nil {
		logrus.Fatalf("failed to read http client root ca: %v", err)
	}
	pool.AppendCertsFromPEM(data)

	return &CDRBankClient{
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
		bankURL: config.BankClientConfig.URL.String(),
	}
}

func (c *CDRBankClient) GetInternalAccounts(id string) (InternalAccounts, error) {
	var (
		token    *oauth2.Token
		request  *http.Request
		response *http.Response
		err      error
	)

	if token, err = c.cc.Token(context.WithValue(context.Background(), oauth2.HTTPClient, c.httpClient)); err != nil {
		return InternalAccounts{}, errors.Wrapf(err, "failed to get client credentials token for internal bank api call")
	}

	if request, err = http.NewRequest(http.MethodPost, c.bankURL+"/internal/accounts", strings.NewReader(
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

	if response.StatusCode >= http.StatusBadRequest {
		var raw []byte
		if raw, err = io.ReadAll(response.Body); err != nil {
			return InternalAccounts{}, errors.Wrap(err, "internal bank accounts api call failed")
		}
		return InternalAccounts{}, errors.Wrap(errors.New(string(raw)), "internal bank accounts api call failed")
	}

	return c.accountsResponseToInternalAccounts(response)
}

func (c *CDRBankClient) accountsResponseToInternalAccounts(response *http.Response) (accounts InternalAccounts, err error) {
	var (
		responseBytes       []byte
		accountListResponse models.ResponseBankingAccountList
	)
	defer response.Body.Close()

	if responseBytes, err = ioutil.ReadAll(response.Body); err != nil {
		return accounts, err
	}

	if err = json.Unmarshal(responseBytes, &accountListResponse); err != nil {
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

func (c *CDRBankClient) GetInternalBalances(id string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
