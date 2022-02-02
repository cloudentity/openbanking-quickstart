package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/models"
)

type OBBRBankClient struct {
	baseURL string
	*http.Client
}

var _ BankClient = &OBBRBankClient{}

func NewOBBRBankClient(config Config) *OBBRBankClient {
	c := OBBRBankClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankURL.String()

	return &c
}

func (c *OBBRBankClient) GetInternalAccounts(id string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = models.ResponseAccountList{}
		err      error
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/internal/accounts?id=%s", c.baseURL, id), nil); err != nil {
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

	return ToInternalOBBRAccounts(resp), nil
}

func ToInternalOBBRAccounts(data models.ResponseAccountList) InternalAccounts {
	accounts := make([]InternalAccount, len(data.Data))
	for i, account := range data.Data {
		accounts[i] = InternalAccount{
			ID:   *account.AccountID,
			Name: *account.Number,
		}
	}
	return InternalAccounts{Accounts: accounts}
}

// TODO: map response to InternalBalances
func (c *OBBRBankClient) GetInternalBalances(id string) (BalanceResponse, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = BalanceResponse{}
		err      error
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/internal/balances?id=%s", c.baseURL, id), nil); err != nil {
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
