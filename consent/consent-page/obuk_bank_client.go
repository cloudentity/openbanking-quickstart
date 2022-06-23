package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
)

type OBUKBankClient struct {
	baseURL string
	*http.Client
}

var _ BankClient = &OBUKBankClient{}

func NewOBUKBankClient(config Config) *OBUKBankClient {
	c := OBUKBankClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankClientConfig.URL.String()

	return &c
}

func (c *OBUKBankClient) GetInternalAccounts(ctx context.Context, id string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = models.OBReadAccount6{}
		err      error
	)

	if request, err = http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/internal/accounts?id=%s", c.baseURL, id), nil); err != nil {
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

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return InternalAccounts{}, err
	}

	return ToInternalOBUKAccounts(resp), nil
}

func ToInternalOBUKAccounts(data models.OBReadAccount6) InternalAccounts {
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
func (c *OBUKBankClient) GetInternalBalances(ctx context.Context, id string) (BalanceResponse, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = BalanceResponse{}
		err      error
	)

	if request, err = http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/internal/balances?id=%s", c.baseURL, id), nil); err != nil {
		return resp, err
	}

	if response, err = c.Client.Do(request); err != nil {
		return resp, err
	}
	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		return resp, err
	}

	if response.StatusCode != 200 {
		return resp, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bytes))
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
