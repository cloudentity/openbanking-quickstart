package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/gin-gonic/gin"
)

type OBUKAccountsGetter struct {
	*Server
}

func (h *OBUKAccountsGetter) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	var (
		accountsResp *accounts.GetAccountsOK
		err          error
	)

	if accountsResp, err = h.BankClient.OBUK.Accounts.GetAccounts(accounts.NewGetAccountsParamsWithContext(c).WithAuthorization(token), nil); err != nil {
		return nil, err
	}

	return accountsResp.Payload, nil
}
