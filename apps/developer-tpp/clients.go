package main

import (
	"net/http"

	obbr "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/client"
	obuk "github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/client"
	httptransport "github.com/go-openapi/runtime/client"
)

type OpenbankingClient struct {
	OBUK *obuk.Accounts
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
		"/accounts/v1",
		[]string{config.BankURL.Scheme},
		hc,
	), nil)

	return c
}
