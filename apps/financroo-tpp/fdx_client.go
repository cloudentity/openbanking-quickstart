package main

import (
	"github.com/gin-gonic/gin"

	acpclient "github.com/cloudentity/acp-client-go"
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

func (f *FDXClient) CreateConsentExplicitly() bool {
	return false
}

func (f *FDXClient) UsePAR() bool {
	return true
}

func (f *FDXClient) DoPAR(c *gin.Context) (string, acpclient.CSRF, error) {
	var (
		csrf acpclient.CSRF
		resp acpclient.PARResponse
		err  error
	)

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

	if resp, csrf, err = f.PublicClient.DoPAR(
		acpclient.WithResponseType("code"),
		acpclient.WithAuthorizationDetails(authorizationDetails),
	); err != nil {
		return "", acpclient.CSRF{}, err
	}
	return resp.RequestURI, csrf, err
}

func (f *FDXClient) CreateAccountConsent(c *gin.Context) (string, error) {
	return "", nil
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
