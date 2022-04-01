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
	*Server
}

func (h *FDXLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	return nil, nil
}

func (h *FDXLogic) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		resp   *a2.PushedAuthorizationRequestCreated
		client acpclient.Client
		err    error
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

	clientConfig := h.Config.ClientConfig()
	clientConfig.ClientSecret = ""
	clientConfig.AuthMethod = acpclient.NoneAuthnMethod

	if client, err = acpclient.New(clientConfig); err != nil {
		return nil, errors.Wrapf(err, "failed to create acp client")
	}

	if resp, err = client.Oauth2.Oauth2.PushedAuthorizationRequest(
		a2.NewPushedAuthorizationRequestParams().
			WithContext(c.Request.Context()).
			WithClientID(h.Config.ClientID).
			WithClientSecret(&h.Config.ClientSecret).
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
		return m["request_uri"].(string)
	}

	return ""
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
	var (
		grantID         string
		ok              bool
		client          acpclient.Client
		resp            *f_d_x.GetFDXConsentOK
		consentResponse []byte
		err             error
	)

	if grantID, ok = data["grant_id"].(string); !ok {
		return nil, errors.New("grant_id is missing")
	}

	clientConfig := h.Config.ClientConfig()
	clientConfig.ClientSecret = h.Config.ClientSecret
	clientConfig.Scopes = []string{"READ_CONSENTS"}
	clientConfig.AuthMethod = acpclient.ClientSecretPostAuthnMethod

	if client, err = acpclient.New(clientConfig); err != nil {
		return nil, errors.Wrapf(err, "failed to create acp client")
	}

	if resp, err = client.GetFDXConsent(
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
