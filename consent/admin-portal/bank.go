package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *BankClient) GetInternalAccounts(subject string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = InternalAccounts{}
		err      error
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/internal/accounts/%s", c.baseURL, subject), nil); err != nil {
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
