package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type FDXBankClient struct {
	httpClient       *http.Client
	bankClientConfig BankClientConfig
	cc               clientcredentials.Config
}

func NewFDXBankClient(config Config) (BankClient, error) {
	var (
		pool  *x509.CertPool
		cert  tls.Certificate
		data  []byte
		certs []tls.Certificate
		err   error
	)

	if pool, err = x509.SystemCertPool(); err != nil {
		return &FDXBankClient{}, errors.Wrap(err, "failed to read system root CAs")
	}

	if data, err = os.ReadFile(config.RootCA); err != nil {
		return &FDXBankClient{}, errors.Wrap(err, "failed to read http client root ca")
	}
	pool.AppendCertsFromPEM(data)

	if config.BankClientConfig.CertFile != "" && config.BankClientConfig.KeyFile != "" {
		if cert, err = tls.LoadX509KeyPair(config.BankClientConfig.CertFile, config.BankClientConfig.KeyFile); err != nil {
			return &FDXBankClient{}, errors.Wrap(err, "failed to read certificate and private key")
		}
		certs = append(certs, cert)
	}
	return &FDXBankClient{
		bankClientConfig: config.BankClientConfig,
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
			ClientID: config.BankClientConfig.ClientID,
			TokenURL: config.BankClientConfig.TokenURL,
			Scopes:   config.BankClientConfig.Scopes,
		},
	}, nil
}

func (c *FDXBankClient) GetInternalAccounts(ctx context.Context, id string) (InternalAccounts, error) {
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

	if request, err = http.NewRequestWithContext(ctx, http.MethodGet, accountsEndpointPath, strings.NewReader(
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
		return InternalAccounts{}, fmt.Errorf("internal bank accounts api returned unexpected status code: %d, body: %s", response.StatusCode, string(body))
	}

	return c.accountsResponseToInternalAccounts(body)
}

func (c *FDXBankClient) accountsResponseToInternalAccounts(body []byte) (accounts InternalAccounts, err error) {
	var accountListResponse models.Accountsentity

	if err = json.Unmarshal(body, &accountListResponse); err != nil {
		return accounts, err
	}

	for _, acc := range accountListResponse.Accounts {
		accMap := acc.(map[string]interface{})
		var keys []string
		for k := range accMap {
			keys = append(keys, k)
		}

		accMapVals := (accMap[keys[0]]).(map[string]interface{})
		accounts.Accounts = append(accounts.Accounts, InternalAccount{
			ID:   accMapVals["accountId"].(string),
			Name: accMapVals["nickname"].(string),
		})
	}

	return accounts, nil
}
