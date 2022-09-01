package main

import (
	"errors"
	"fmt"
	"log"

	cdrModels "github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/client/banking"
	fdxAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/client/account_information"
	fdxModel "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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

func (o *CDRClient) GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) (accountsData []Account, err error) {
	var resp *cdrModels.ListAccountsOK

	if resp, err = o.Banking.Banking.ListAccounts(
		cdrModels.NewListAccountsParamsWithContext(c).
			WithDefaults(),
		runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
			return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}),
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

func (o *FDXBankClient) GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error) {
	var (
		resp         *fdxAccounts.SearchForAccountsOK
		accountsData = []Account{}
		err          error
	)

	log.Printf("Passed in token from financroo '%s' for bank %+v", accessToken, bank.BankID)
	if resp, err = o.AccountInformation.SearchForAccounts(fdxAccounts.NewSearchForAccountsParamsWithContext(c).WithAccountIds([]string{"10001"}), httptransport.BearerToken(accessToken)); err != nil {
		log.Printf("Failed to search for accounts %v", err)
		return accountsData, err
	}

	log.Printf("Resp from bank %+v", resp.Payload)
	for _, a := range resp.Payload.Accounts {
		v, ok := a.(fdxModel.Accountentity)
		if !ok {
			return accountsData, errors.New("unable to get account from accounts data")
		}

		accountsData = append(accountsData, Account{
			OBAccount6: models.OBAccount6{
				AccountID: (*models.AccountID)(&v.AccountID),
				Nickname:  models.Nickname(v.Nickname),
				Account: []*models.OBAccount6AccountItems0{
					{
						Name:           models.Name0(v.ProductName),
						Identification: (*models.Identification0)(&v.AccountNumber),
					},
				},
			},
			BankID: bank.BankID,
		})
	}

	return accountsData, nil
}
