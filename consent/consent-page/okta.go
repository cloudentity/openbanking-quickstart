package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OktaHandler interface {
	GetAPIToken() string
	GetOktaID(token, user string) (string, error)
	HasFactorType(token, id, factorType string) bool
	GetVerifyURL(token, id string) (string, error)
	SendVerify(token, id, verifyURL string) (string, error)
	GetVerificationStatus(token, pollURL string) (string, error)
}

type OktaAPIResponseData []struct {
	ID         string `json:"id"`
	FactorType string `json:"factorType"`
	Links      []struct {
		Verify struct {
		} `json:"verify"`
	} `json:"_links"`
}

type DefaultOktaHandler struct {
	HostURL  string
	Client   *http.Client
	APIToken string
}

func NewDefaultOktaHandler(host string, apiToken string) OktaHandler {
	return &DefaultOktaHandler{
		HostURL: host,
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		APIToken: apiToken,
	}
}

func (o *DefaultOktaHandler) GetAPIToken() string {
	return o.APIToken
}

type OKTAUsersResponsData struct {
}

func (o *DefaultOktaHandler) GetOktaID(token, user string) (string, error) {
	var (
		endpoint = fmt.Sprintf("/api/v1/users?q=%s&limit=1", user)
		req      *http.Request
		resp     *http.Response
		err      error
	)

	if req, err = http.NewRequest("GET", o.HostURL+endpoint, nil); err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("SSWS %s", token))

	if resp, err = o.Client.Do(req); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// response data
	data := []struct {
		ID string `json:"id"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if len(data) == 0 {
		return fmt.Errorf("okta /user endpoint returned no data")
	}

	return data[0].ID, nil
}

func (o *DefaultOktaHandler) HasFactorType(token, id, factorType string) bool {
	var (
		endpoint = fmt.Sprintf("/api/v1/users/%s/factors", id)
		req      *http.Request
		resp     *http.Response
		err      error
	)

	if req, err = http.NewRequest("GET", o.HostURL+endpoint, nil); err != nil {
		return false
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("SSWS %s", token))

	if resp, err = o.Client.Do(req); err != nil {
		return false
	}
	defer resp.Body.Close()

	data := []struct {
		FactorType string `json:"factorType"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false
	}

	for _, d := range data {
		if d.FactorType == "push" {
			return true
		}
	}

	return false
}

func (o *DefaultOktaHandler) GetVerifyURL(token, id string) (string, error) {
	var (
		endpoint = fmt.Sprintf("/api/v1/users/%s/factors", id)
		req      *http.Request
		resp     *http.Response
		err      error
	)

	if req, err = http.NewRequest("GET", o.HostURL+endpoint, nil); err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("SSWS %s", token))

	if resp, err = o.Client.Do(req); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := []struct {
		FactorType string `json:"factorType"`
		Links      struct {
			Verify struct {
				Href string `json:"href"`
			} `json:"verify"`
		} `json:"_links"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	for _, d := range data {
		if d.FactorType == "push" {
			return d.Links.Verify.Href, nil
		}
	}

	return "", nil
}

func (o *DefaultOktaHandler) SendVerify(token, id, verifyURL string) (string, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	if req, err = http.NewRequest("POST", verifyURL, nil); err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("SSWS %s", token))

	if resp, err = o.Client.Do(req); err != nil {
		return "", err
	}

	data := struct {
		FactorResult string `json:"factorResult"`
		Links        struct {
			Poll struct {
				Href string `json:"href"`
			} `json:"poll"`
		} `json:"_links"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Links.Poll.Href, nil
}

func (o *DefaultOktaHandler) GetVerificationStatus(token, pollURL string) (string, error) {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	if req, err = http.NewRequest("GET", pollURL, nil); err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("SSWS %s", token))

	if resp, err = o.Client.Do(req); err != nil {
		return "", err
	}

	data := struct {
		FactorResult string `json:"factorResult"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.FactorResult, nil
}
