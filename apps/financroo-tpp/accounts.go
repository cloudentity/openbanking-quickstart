package main

import (
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/gin-gonic/gin"
)

type Account struct {
	AccountID string `json:"AccountId"`
	Nickname  string `json:"Nickname"`
	BankID    string `json:"BankId"`
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
			AccountID: string(*a.AccountID),
			Nickname:  string(a.Nickname),
			BankID:    bank.BankID,
		})
	}

	return accountsData, nil
}

func (o *OBBRClient) GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error) {
	var (
		resp         *obbrAccounts.AccountsGetAccountsOK
		accountsData = []Account{}
		err          error
	)
	if resp, err = o.Accounts.Accounts.AccountsGetAccounts(obbrAccounts.NewAccountsGetAccountsParamsWithContext(c).WithAuthorization(accessToken), nil); err != nil {
		return accountsData, err
	}

	for _, a := range resp.Payload.Data {
		accountsData = append(accountsData, Account{
			AccountID: *a.Number,
			Nickname:  *a.AccountID,
		})
	}

	return accountsData, nil
}
