package main

import (
	"fmt"

	"github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/client/banking"
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/balances"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
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

func (o *CDRClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) (balancesData []Balance, err error) {
	var resp *banking.ListBalancesBulkOK

	if resp, err = o.Banking.Banking.ListBalancesBulk(
		banking.NewListBalancesBulkParams().
			WithDefaults(),
		runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
			return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		}),
	); err != nil {
		return []Balance{}, err
	}

	for _, balance := range resp.Payload.Data.Balances {
		balancesData = append(balancesData, Balance{
			AccountID: *balance.AccountID,
			Amount:    *balance.AvailableBalance,
			Currency:  balance.Currency,
			BankID:    bank.BankID,
		})
	}

	return balancesData, nil
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
