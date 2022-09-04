package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FDXClient struct {
	baseURL string
	*http.Client
}

var _ BankClient = &FDXClient{}

func NewFDXClient(config Config) *FDXClient {
	c := FDXClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankClientConfig.URL.String()

	return &c
}

type FdxAccounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	DepositAccount DepositAccount `json:"depositAccount"`
}

type DepositAccount struct {
	AccountID      string  `json:"accountId"`
	BalanceAsOf    string  `json:"balanceAsOf"`
	CurrentBalance float64 `json:"currentBalance"`
	Nickname       string  `json:"nickname"`
	Status         string  `json:"status"`
}

func (c *FDXClient) GetInternalAccounts(ctx context.Context, id string) (InternalAccounts, error) {
	if id == "user" {
		var (
			request     *http.Request
			response    *http.Response
			bytes       []byte
			fdxAccounts FdxAccounts
			err         error
		)

		if request, err = http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/internal/accounts", c.baseURL), http.NoBody); err != nil {
			return InternalAccounts{}, err
		}

		if response, err = c.Client.Do(request); err != nil {
			return InternalAccounts{}, err
		}
		defer response.Body.Close()

		if bytes, err = ioutil.ReadAll(response.Body); err != nil {
			return InternalAccounts{}, err
		}

		if response.StatusCode != 200 {
			return InternalAccounts{}, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bytes))
		}

		if err = json.Unmarshal(bytes, &fdxAccounts); err != nil {
			return InternalAccounts{}, err
		}

		return ToInternalFDXAccounts(fdxAccounts), nil
	}

	return InternalAccounts{
		Accounts: []InternalAccount{
			{
				ID:   "96565987",
				Name: "Credit",
				Balance: Balance{
					AccountID: "96565987",
					Amount: BalanceAmount{
						Amount:   "100",
						Currency: "USD",
					},
				},
			},
			{
				ID:   "1122334455",
				Name: "Savings",
				Balance: Balance{
					AccountID: "1122334455",
					Amount: BalanceAmount{
						Amount:   "150",
						Currency: "USD",
					},
				},
			},
		},
	}, nil
}

func (c *FDXClient) GetInternalBalances(ctx context.Context, id string) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}

func ToInternalFDXAccounts(data FdxAccounts) InternalAccounts {
	accounts := make([]InternalAccount, len(data.Accounts))
	for i, account := range data.Accounts {
		accounts[i] = InternalAccount{
			ID:   string(account.DepositAccount.AccountID),
			Name: string(account.DepositAccount.Nickname),
		}
	}
	return InternalAccounts{Accounts: accounts}
}
