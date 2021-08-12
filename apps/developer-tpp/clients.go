package main

import (
	"net/http"

	obbr "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client"
	obuk "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client"
	httptransport "github.com/go-openapi/runtime/client"
)

type OpenbankingClient struct {
	OBUK *obuk.OpenbankingAccountsClient
	OBBR *obbr.Accounts
}

func NewOpenbankingClient(config Config) OpenbankingClient {
	var (
		c  = OpenbankingClient{}
		hc = &http.Client{}
	)

	c.OBUK = obuk.New(httptransport.NewWithClient(
		config.BankURL.Host,
		"/",
		[]string{config.BankURL.Scheme},
		hc,
	), nil)

	c.OBBR = obbr.New(httptransport.NewWithClient(
		config.BankURL.Host,
		"/",
		[]string{config.BankURL.Scheme},
		hc,
	), nil)

	return c
}
