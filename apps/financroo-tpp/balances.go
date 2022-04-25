package main

import (
	"fmt"

	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/balances"
	"github.com/gin-gonic/gin"
)

type Balance struct {
	AccountID string `json:"AccountId"`
	Amount    string `json:"Amount"`
	Currency  string `json:"Currency"`
	BankID    string `json:"BankId"`
}

func (o *OBUKClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	var (
		resp         *balances.GetBalancesOK
		balancesData = []Balance{}
		err          error
	)

	if resp, err = o.Balances.GetBalances(balances.NewGetBalancesParamsWithContext(c).WithAuthorization(accessToken), nil); err != nil {
		return balancesData, err
	}

	for _, a := range resp.Payload.Data.Balance {
		amount := a.Amount
		balancesData = append(balancesData, Balance{
			AccountID: string(*a.AccountID),
			Amount:    string(*amount.Amount),
			Currency:  string(*amount.Currency),
			BankID:    bank.BankID,
		})
	}

	return balancesData, nil
}

func (o *CDRClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	return []Balance{}, nil
}

func (o *OBBRClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	var (
		resp         *obbrAccounts.AccountsGetAccountsAccountIDBalancesOK
		balancesData = []Balance{}
		err          error
		accounts     []Account
	)

	if accounts, err = o.GetAccounts(c, accessToken, bank); err != nil {
		return nil, err
	}

	for _, acc := range accounts {
		accountID := string(*acc.AccountID)

		if resp, err = o.Accounts.Accounts.AccountsGetAccountsAccountIDBalances(obbrAccounts.NewAccountsGetAccountsAccountIDBalancesParamsWithContext(c).
			WithAccountID(accountID).
			WithAuthorization(accessToken), nil); err != nil {
			return balancesData, err
		}

		data := resp.Payload.Data
		balancesData = append(balancesData, Balance{
			AccountID: accountID,
			Amount:    fmt.Sprintf("%f", *data.AvailableAmount),
			Currency:  *data.AvailableAmountCurrency,
			BankID:    bank.BankID,
		})
	}

	return balancesData, nil
}
