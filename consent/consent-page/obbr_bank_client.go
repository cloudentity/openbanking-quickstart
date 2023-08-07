package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/models"
)

type OBBRBankClient struct {
	baseURL string
	*http.Client
}

var _ BankClient = &OBBRBankClient{}

func NewOBBRBankClient(config Config) *OBBRBankClient {
	c := OBBRBankClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankClientConfig.URL.String()

	return &c
}

func (c *OBBRBankClient) GetInternalAccounts(ctx context.Context, id string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = models.ResponseAccountList{}
		err      error
	)

	if request, err = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/internal/accounts?id=%s", c.baseURL, id), http.NoBody); err != nil {
		return InternalAccounts{}, err
	}

	if response, err = c.Client.Do(request); err != nil {
		return InternalAccounts{}, err
	}
	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		return InternalAccounts{}, err
	}

	if response.StatusCode != http.StatusOK {
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

type OBBRBalances struct {
	Data []struct {
		models.AccountBalancesData
		AccountID string `json:"accountId"`
	} `json:"data"`
}

// TODO: map response to InternalBalances
func (c *OBBRBankClient) GetInternalBalances(ctx context.Context, id string) (BalanceResponse, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     OBBRBalances
		err      error
	)

	if request, err = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/internal/balances?id=%s", c.baseURL, id), http.NoBody); err != nil {
		return BalanceResponse{}, err
	}

	if response, err = c.Client.Do(request); err != nil {
		return BalanceResponse{}, err
	}
	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		return BalanceResponse{}, err
	}

	if response.StatusCode != http.StatusOK {
		return BalanceResponse{}, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bytes))
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return BalanceResponse{}, err
	}

	return c.ToBalanceResponse(resp), nil
}

func (c *OBBRBankClient) ToBalanceResponse(data OBBRBalances) BalanceResponse {
	var balances []Balance

	for _, r := range data.Data {
		balances = append(balances, Balance{
			AccountID: r.AccountID,
			Amount: BalanceAmount{
				Amount:   fmt.Sprintf("%.2f", *r.AvailableAmount),
				Currency: *r.AvailableAmountCurrency,
			},
		})
	}
	return BalanceResponse{
		Data: BalanceData{
			Balance: balances,
		},
	}
}
