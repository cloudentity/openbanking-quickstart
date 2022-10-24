package main

import (
	"encoding/json"
	"fmt"

	cdrModels "github.com/cloudentity/openbanking-quickstart/openbanking/cdr/client/banking"
	fdxAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/account_information"
	fdxModels "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/models"
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accounts/models"
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

	if resp, err = o.Accounts.Accounts.GetAccounts(accounts.NewGetAccountsParamsWithContext(c).WithAuthorization(accessToken), nil); err != nil {
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

	if resp, err = o.AccountInformation.SearchForAccounts(fdxAccounts.NewSearchForAccountsParamsWithContext(c), httptransport.BearerToken(accessToken)); err != nil {
		return accountsData, err
	}

	for _, acct := range resp.Payload.Accounts {
		var (
			depositAccount fdxModels.AccountWithDetailsentity
			jsonStr        []byte
		)
		if jsonStr, err = json.Marshal(acct); err != nil {
			return accountsData, err
		}

		if err = json.Unmarshal(jsonStr, &depositAccount); err != nil {
			return accountsData, err
		}

		accountsData = append(accountsData, Account{
			OBAccount6: models.OBAccount6{
				AccountID: (*models.AccountID)(&depositAccount.DepositAccount.AccountID),
				Nickname:  models.Nickname(depositAccount.DepositAccount.Nickname),
				Account: []*models.OBAccount6AccountItems0{
					{
						Name:           models.Name0(depositAccount.DepositAccount.ProductName),
						Identification: (*models.Identification0)(&depositAccount.DepositAccount.AccountNumber),
					},
				},
			},
			BankID: bank.BankID,
		})
	}

	return accountsData, nil
}
