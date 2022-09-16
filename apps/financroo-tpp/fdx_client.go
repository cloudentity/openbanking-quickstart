package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	a2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

type FDXClient struct {
	ClientID     string
	ClientSecret string
	PublicClient acpclient.Client
}

func NewFDXConsentClient(publicClient, clientCredentialsClient acpclient.Client, _ Signer) ConsentClient {
	return &FDXClient{
		ClientID:     clientCredentialsClient.Config.ClientID,
		ClientSecret: clientCredentialsClient.Config.ClientSecret,
		PublicClient: publicClient,
	}
}

func (f *FDXClient) CreateAccountConsent(c *gin.Context) (string, error) {
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

	scopes := "offline_access ACCOUNT_DETAILED READ_CONSENTS ACCOUNT_BASIC TRANSACTIONS"
	if resp, err = f.PublicClient.Oauth2.Oauth2.PushedAuthorizationRequest(
		a2.NewPushedAuthorizationRequestParams().
			WithContext(c.Request.Context()).
			WithRedirectURI(f.PublicClient.Config.RedirectURL.String()).
			WithClientID(f.ClientID).
			WithScope(&scopes).
			WithResponseType(responseType).
			WithAuthorizationDetails(&authorizationDetails),
	); err != nil {
		return "", errors.Wrapf(err, "failed to register par request")
	}

	return resp.Payload.RequestURI, err
}

func (f *FDXClient) DoRequestObjectEncryption() bool {
	return false
}

func (f *FDXClient) GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error) {
	return "", nil
}

func (f *FDXClient) CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error) {
	return "", nil
}

func (f *FDXClient) Sign([]byte) (string, error) {
	return "", nil
}
