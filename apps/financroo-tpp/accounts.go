package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"

	"github.com/gin-gonic/gin"
)

type Account struct {
	models.OBAccount6
	/*AccountID string `json:"AccountId"`
	Amount    string `json:"Amount"`
	Nickname  string `json:"Nickname"`*/
	BankID string `json:"BankId"`
}

func (o *OBUKClient) GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error) {
	var (
		resp         *accounts.GetAccountsOK
		accountsData = []Account{}
		err          error
	)

	if resp, err = o.Accounts.GetAccounts(accounts.NewGetAccountsParamsWithContext(c).WithAuthorization(accessToken), nil); err != nil {
		return accountsData, err
	}

	for _, a := range resp.Payload.Data.Account {
		accountsData = append(accountsData, Account{
			OBAccount6: *a,
			//	AccountID:  string(*a.AccountID),
			//	Nickname:   string(a.Nickname),
			BankID: bank.BankID,
		})
	}

	return accountsData, nil
}

func (o *OBBRClient) GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error) {
	return []Account{}, nil
}
