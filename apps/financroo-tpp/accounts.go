package main

import (
	cdrModels "github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/client/banking"
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
)

type Account struct {
	models.OBAccount6
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
			BankID:     bank.BankID,
		},
		)
	}

	return accountsData, nil
}

// TODO: AUT-5813
func (o *CDRClient) GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) (accountsData []Account, err error) {
	var (
		resp *cdrModels.ListAccountsOK
	)

	if resp, err = o.Banking.Banking.ListAccounts(
		cdrModels.NewListAccountsParamsWithContext(c).
			WithDefaults(),
	); err != nil {
		return accountsData, err
	}

	for _, a := range resp.Payload.Data.Accounts {
		accountsData = append(accountsData, Account{
			OBAccount6: models.OBAccount6{
				AccountID: (*models.AccountID)(a.AccountID),
				Nickname:  models.Nickname(a.Nickname),
				Account: []*models.OBAccount6AccountItems0{
					{
						Name:           models.Name0(*a.AccountID),
						Identification: (*models.Identification0)(a.MaskedNumber),
					},
				},
			},
			BankID: bank.BankID,
		})
	}

	return accountsData, err
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
			OBAccount6: models.OBAccount6{
				AccountID: (*models.AccountID)(a.AccountID),
				Nickname:  models.Nickname(*a.AccountID),
				Account: []*models.OBAccount6AccountItems0{
					{
						Name:           models.Name0(*a.AccountID),
						Identification: (*models.Identification0)(a.Number),
					},
				},
			},
			BankID: bank.BankID,
		})
	}

	return accountsData, nil
}
