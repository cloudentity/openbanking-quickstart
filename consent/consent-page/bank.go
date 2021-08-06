package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudentity/openbanking-quickstart/models"
)

type BankClient struct {
	baseURL string
	*http.Client
}

func NewBankClient(config Config) BankClient {
	c := BankClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankURL.String()

	return c
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance Balance `json:"balance"`
}

type Balance struct {
	AccountID string        `json:"AccountId"`
	Amount    BalanceAmount `json:"Amount"`
}

type BalanceAmount struct {
	Amount   string
	Currency string
}

type BalanceResponse struct {
	Data BalanceData `json:"Data"`
}

type BalanceData struct {
	Balance []Balance `json:"Balance"`
}

func (c *BankClient) GetInternalAccounts(subject string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = models.OBReadAccount6{}
		err      error
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/internal/accounts/%s", c.baseURL, subject), nil); err != nil {
		return InternalAccounts{}, nil
	}

	if response, err = c.Client.Do(request); err != nil {
		return InternalAccounts{}, nil
	}
	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		return InternalAccounts{}, nil
	}

	if response.StatusCode != 200 {
		return InternalAccounts{}, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bytes))
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return InternalAccounts{}, nil
	}

	return ToInternalAccounts(resp), nil
}

func ToInternalAccounts(data models.OBReadAccount6) InternalAccounts {
	accounts := make([]InternalAccount, len(data.Data.Account))
	for i, account := range data.Data.Account {
		accounts[i] = InternalAccount{
			ID:   string(*account.AccountID),
			Name: string(account.Nickname),
		}
	}
	return InternalAccounts{Accounts: accounts}
}

// TODO: map response to InternalBalances
func (c *BankClient) GetInternalBalances(subject string) (BalanceResponse, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = BalanceResponse{}
		err      error
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/internal/balances/%s", c.baseURL, subject), nil); err != nil {
		return resp, nil
	}

	if response, err = c.Client.Do(request); err != nil {
		return resp, nil
	}
	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		return resp, nil
	}

	if response.StatusCode != 200 {
		return resp, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bytes))
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return resp, nil
	}

	return resp, nil
}
