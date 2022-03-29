package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/oauth2/models"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type FDXLogic struct {
	*Server
}

func (h *FDXLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	return nil, nil
}

func (h *FDXLogic) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		resp *http.Response
		bs   []byte
		err  error
	)

	responseType := "code"
	authorizationDetails := `[
      {
         "type":"fdx_v1.0",
         "consentRequest":{
            "durationType":"ONE_TIME",
            "lookbackPeriod":60,
            "resources":[
               {
                  "resourceType":"ACCOUNT",
                  "dataClusters":[
                     "ACCOUNT_DETAILED",
                     "TRANSACTIONS",
                     "STATEMENTS"
                  ]
               }
            ]
         }
      }
   ]`

	u := h.Client.Config.IssuerURL.String() + "/par"
	contentType := "application/x-www-form-urlencoded"

	params := url.Values{
		"client_id":             {h.Config.ClientID},
		"client_secret":         {h.Config.ClientSecret},
		"response_type":         {responseType},
		"authorization_details": {authorizationDetails},
	}

	if resp, err = h.hc().Post(u, contentType, strings.NewReader(params.Encode())); err != nil {
		return nil, errors.Wrapf(err, "failed to register par request")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Wrapf(err, "par endpoint returned unexpected status code")
	}

	if bs, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, errors.Wrapf(err, "par endpoint did not return the response body")
	}

	var parResponse models.PARResponse

	if err = json.Unmarshal(bs, &parResponse); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal par response")
	}

	return &map[string]interface{}{
		"request_uri": parResponse.RequestURI,
	}, nil
}

func (h *FDXLogic) GetConsentID(data interface{}) string {
	if m, ok := data.(*map[string]interface{}); ok {
		m := *m
		return m["request_uri"].(string)
	}

	return "n/a"
}

func (h *FDXLogic) DoRequestObjectEncryption() bool {
	return false
}

func (h *FDXLogic) BuildLoginURL(c *gin.Context, consentID string, _ bool) (string, acpclient.CSRF, error) {
	var (
		client acpclient.Client
		config = h.Config.ExtendConsentScope(consentID).ClientConfig()
		u      string
		err    error
	)

	if client, err = acpclient.New(config); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create new acp client")
	}

	if u, err = client.AuthorizeURLWithPAR(consentID); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create authorize url with par")
	}

	return u, acpclient.CSRF{}, nil
}

func (h *FDXLogic) PostAuthenticationAction(c *gin.Context, data map[string]interface{}) (map[string]interface{}, error) {
	grantID, ok := data["grant_id"].(string)
	if !ok {
		return nil, errors.New("grant_id is missing")
	}

	cc := clientcredentials.Config{
		ClientID:     h.Config.ClientID,
		ClientSecret: h.Config.ClientSecret,
		TokenURL:     h.Config.TokenURL.String(),
		Scopes:       []string{"READ_CONSENTS"},
	}

	hc := cc.Client(context.WithValue(c.Request.Context(), oauth2.HTTPClient, h.hc()))

	u := fmt.Sprintf("%s/consents/%s", h.Client.Config.IssuerURL.String(), grantID)

	var (
		resp *http.Response
		bs   []byte
		err  error
	)

	if resp, err = hc.Get(u); err != nil {
		return nil, errors.Wrapf(err, "failed to call url"+u)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Wrapf(err, "consent endpoint returned unexpected status code")
	}

	if bs, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, errors.Wrapf(err, "consent endpoint did not return the response body")
	}

	var consent obModels.GetFDXConsent

	if err = json.Unmarshal(bs, &consent); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal consent response")
	}

	consentResponse, _ := json.MarshalIndent(&consent, "", "  ")

	return map[string]interface{}{
		"consent":          consent,
		"consent_response": string(consentResponse),
	}, nil
}

func (h *FDXLogic) hc() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
