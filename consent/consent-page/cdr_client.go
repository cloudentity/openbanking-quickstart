package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type CDRBankClient struct {
	httpClient *http.Client
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
	}
}

// TODO: expose endpoint on mock data holder instead of redundant hardcoded mocking in every application
// TODO: this id will be customer id and will be passed in request body
func (c *CDRBankClient) GetInternalAccounts(id string) (InternalAccounts, error) {
	// test token retrieval
	token, err := c.cc.Token(context.WithValue(context.Background(), oauth2.HTTPClient, c.httpClient))
	if err != nil {
		logrus.Infof("error retrieving cc token %v", err)
	} else {
		logrus.Infof("retrieved cc token with no errors %+v", token)
	}

	if id == "user" {
		return InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   "96534987",
					Name: "Digital banking account",
					Balance: Balance{
						AccountID: "96534987",
						Amount: BalanceAmount{
							Amount:   "100",
							Currency: "USD",
						},
					},
				},
				{
					ID:   "1000001",
					Name: "Savings",
					Balance: Balance{
						AccountID: "1000001",
						Amount: BalanceAmount{
							Amount:   "150",
							Currency: "USD",
						},
					},
				},
				{
					ID:   "1000002",
					Name: "Savings 2",
					Balance: Balance{
						AccountID: "1000002",
						Amount: BalanceAmount{
							Amount:   "175",
							Currency: "USD",
						},
					},
				},
			},
		}, nil
	}

	return InternalAccounts{
		Accounts: []InternalAccount{
			{
				ID:   "96565987",
				Name: "Credit",
				Balance: Balance{
					AccountID: "96565987",
					Amount: BalanceAmount{
						Amount:   "100",
						Currency: "USD",
					},
				},
			},
			{
				ID:   "1122334455",
				Name: "Savings",
				Balance: Balance{
					AccountID: "1122334455",
					Amount: BalanceAmount{
						Amount:   "150",
						Currency: "USD",
					},
				},
			},
		},
	}, nil
}

// TODO: mock data holder cdr app doesn't even have this data yet
func (c *CDRBankClient) GetInternalBalances(id string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
