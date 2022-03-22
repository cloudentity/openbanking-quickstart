package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	oauth2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

type FDXLogic struct {
	*Server
}

func (h *FDXLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	return nil, nil
}

func (h *FDXLogic) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		client acpclient.Client
		resp   *oauth2.PushedAuthorizationRequestCreated
		err    error
	)

	config := h.Client.Config
	config.ClientSecret = h.Config.ClientSecret
	config.AuthMethod = acpclient.ClientSecretPostAuthnMethod

	if client, err = acpclient.New(config); err != nil {
		return nil, errors.Wrapf(err, "failed to create acp client")
	}

	// TODO change authn method
	if h.Config.ClientSecret == "" {
		return nil, errors.New("client secret must be set")
	}

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

	if resp, err = client.Oauth2.Oauth2.PushedAuthorizationRequest(oauth2.NewPushedAuthorizationRequestParams().
		WithContext(c).
		WithClientID(h.Config.ClientID).
		WithClientSecret(&h.Config.ClientSecret).
		WithResponseType(responseType).
		WithAuthorizationDetails(&authorizationDetails)); err != nil {
		return nil, errors.Wrapf(err, "failed to register par request")
	}

	return &map[string]interface{}{
		"request_uri": resp.Payload.RequestURI,
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

	return u, acpclient.CSRF{}, nil // TODO handle csrf
}
