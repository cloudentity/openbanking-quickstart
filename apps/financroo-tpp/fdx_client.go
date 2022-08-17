package main

import (
	"encoding/json"

	acpclient "github.com/cloudentity/acp-client-go"
	a2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type FDXConsentClient interface {
	CreateConsent(c *gin.Context) (interface{}, error)
}

type FDXClient struct {
	ClientID                string
	ClientSecret            string
	PublicClient            acpclient.Client
	ClientCredentialsClient acpclient.Client
}

type FDXConsentClientFn func(publicClient, clientCredentialsClient acpclient.Client) FDXConsentClient

func NewFDXClient(publicClient, clientCredentialsClient acpclient.Client) FDXConsentClient {
	return &FDXClient{
		ClientID:                clientCredentialsClient.Config.ClientID,
		ClientSecret:            clientCredentialsClient.Config.ClientSecret,
		PublicClient:            clientCredentialsClient,
		ClientCredentialsClient: publicClient,
	}
}

func (f *FDXClient) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		resp *a2.PushedAuthorizationRequestCreated
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
               },
               {
                  "resourceType":"CUSTOMER",
                  "dataClusters":[
                     "CUSTOMER_CONTACT"
                  ]
               }
            ]
         }
      }
   ]`

	if resp, err = f.PublicClient.Oauth2.Oauth2.PushedAuthorizationRequest(
		a2.NewPushedAuthorizationRequestParams().
			WithContext(c.Request.Context()).
			WithClientID(f.ClientID).
			WithClientSecret(&f.ClientSecret).
			WithResponseType(responseType).
			WithAuthorizationDetails(&authorizationDetails),
	); err != nil {
		return "", errors.Wrapf(err, "failed to register par request")
	}

	strResp, err := json.Marshal(&map[string]interface{}{
		"request_uri": resp.Payload.RequestURI,
	})

	return string(strResp), err
}

func (f *FDXClient) GetConsentID(data interface{}) string {
	if m, ok := data.(*map[string]interface{}); ok {
		m := *m
		return m["request_uri"].(string) // nolint
	}

	return ""
}

func (f *FDXClient) DoRequestObjectEncryption() bool {
	return false
}

func (f *FDXClient) BuildLoginURL(c *gin.Context, consentID string, _ bool) (string, acpclient.CSRF, error) {
	var (
		u   string
		err error
	)

	if u, err = f.PublicClient.AuthorizeURLWithPAR(consentID); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create authorize url with par")
	}

	return u, acpclient.CSRF{}, nil
}
