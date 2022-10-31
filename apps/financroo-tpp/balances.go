package main

import (
	"encoding/json"
	"fmt"

	"github.com/cloudentity/openbanking-quickstart/generated/cdr/client/banking"
	fdxAccounts "github.com/cloudentity/openbanking-quickstart/generated/fdx/client/account_information"
	fdxModels "github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/client/balances"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
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

func (o *FDXBankClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	var (
		resp        *fdxAccounts.SearchForAccountsOK
		balance     *fdxAccounts.GetAccountOK
		balanceData = []Balance{}
		err         error
	)

	if resp, err = o.AccountInformation.SearchForAccounts(fdxAccounts.NewSearchForAccountsParamsWithContext(c), httptransport.BearerToken(accessToken)); err != nil {
		return balanceData, err
	}

	for _, acct := range resp.Payload.Accounts {
		var (
			depositAccount fdxModels.AccountWithDetailsentity
			acctBalance    fdxModels.DepositAccountentity2
			jsonStr        []byte
		)
		if jsonStr, err = json.Marshal(acct); err != nil {
			return balanceData, err
		}

		if err = json.Unmarshal(jsonStr, &depositAccount); err != nil {
			return balanceData, err
		}

		if balance, err = o.AccountInformation.GetAccount(fdxAccounts.NewGetAccountParamsWithContext(c).WithAccountID(depositAccount.DepositAccount.AccountID), httptransport.BearerToken(accessToken)); err != nil {
			return balanceData, err
		}

		if jsonStr, err = json.Marshal(balance.Payload); err != nil {
			return balanceData, err
		}

		if err = json.Unmarshal(jsonStr, &acctBalance); err != nil {
			return balanceData, err
		}

		balanceData = append(balanceData, Balance{
			AccountID: acctBalance.AccountID,
			Amount:    fmt.Sprint(depositAccount.DepositAccount.CurrentBalance),
			Currency:  string(acctBalance.Currency.CurrencyCode),
			BankID:    bank.BankID,
		})
	}

	return balanceData, nil
}
