package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	a2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	"github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
)

type FDXLogic struct {
	ClientID                string
	ClientSecret            string
	PublicClient            acpclient.Client
	ClientCredentialsClient acpclient.Client
}

func NewFDXLogic(serverConfig Config) (*FDXLogic, error) {
	var (
		logic = &FDXLogic{
			ClientID:     serverConfig.ClientID,
			ClientSecret: serverConfig.ClientSecret,
		}
		err error
	)

	publicConfig := serverConfig.ClientConfig()
	publicConfig.SkipClientCredentialsAuthn = true

	if logic.PublicClient, err = acpclient.New(publicConfig); err != nil {
		return logic, errors.Wrapf(err, "failed to create public acp client")
	}

	ccConfig := serverConfig.ClientConfig()
	ccConfig.ClientSecret = serverConfig.ClientSecret
	ccConfig.Scopes = []string{"READ_CONSENTS"}
	ccConfig.AuthMethod = acpclient.ClientSecretPostAuthnMethod

	if logic.ClientCredentialsClient, err = acpclient.New(ccConfig); err != nil {
		return logic, errors.Wrapf(err, "failed to create client credentials acp client")
	}

	return logic, nil
}

func (h *FDXLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	return nil, nil // nolint
}

func (h *FDXLogic) CreateConsent(c *gin.Context) (interface{}, error) {
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

	if resp, err = h.PublicClient.Oauth2.Oauth2.PushedAuthorizationRequest(
		a2.NewPushedAuthorizationRequestParams().
			WithContext(c.Request.Context()).
			WithClientID(h.ClientID).
			WithRedirectURI(h.PublicClient.Config.RedirectURL.String()).
			WithClientSecret(&h.ClientSecret).
			WithResponseType(responseType).
			WithAuthorizationDetails(&authorizationDetails),
	); err != nil {
		return nil, errors.Wrapf(err, "failed to register par request")
	}

	return &map[string]interface{}{
		"request_uri": resp.Payload.RequestURI,
	}, nil
}

func (h *FDXLogic) GetConsentID(data interface{}) string {
	if m, ok := data.(*map[string]interface{}); ok {
		m := *m
		return m["request_uri"].(string) // nolint
	}

	return ""
}

func (h *FDXLogic) DoRequestObjectEncryption() bool {
	return false
}

func (h *FDXLogic) BuildLoginURL(c *gin.Context, consentID string, _ bool) (string, acpclient.CSRF, error) {
	var (
		u   string
		err error
	)

	if u, err = h.PublicClient.AuthorizeURLWithPAR(consentID); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create authorize url with par")
	}

	return u, acpclient.CSRF{}, nil
}

func (h *FDXLogic) PostAuthenticationAction(c *gin.Context, data map[string]interface{}) (map[string]interface{}, error) {
	var (
		grantID         string
		ok              bool
		resp            *f_d_x.GetFDXConsentOK
		consentResponse []byte
		err             error
	)

	if grantID, ok = data["grant_id"].(string); !ok {
		return nil, errors.New("grant_id is missing")
	}

	if resp, err = h.ClientCredentialsClient.GetFDXConsent(
		f_d_x.NewGetFDXConsentParams().
			WithContext(c.Request.Context()).
			WithConsentID(grantID),
	); err != nil {
		return nil, errors.Wrapf(err, "failed to get consent")
	}

	if consentResponse, err = json.MarshalIndent(&resp.Payload, "", "  "); err != nil {
		return nil, errors.Wrapf(err, "failed to marshal consent response")
	}

	return map[string]interface{}{
		"consent":          resp.Payload,
		"consent_response": string(consentResponse),
	}, nil
}

var _ SpecLogicHandler = &FDXLogic{}
