package main

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/fdx/client/f_d_x"
	a2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

type FDXLogic struct {
	ClientID     string
	ClientSecret string
	ACPClient    acpclient.Client
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
	publicConfig.AuthMethod = acpclient.TLSClientAuthnMethod
	publicConfig.Scopes = serverConfig.ClientScopes

	if logic.ACPClient, err = acpclient.New(publicConfig); err != nil {
		return logic, errors.Wrapf(err, "failed to create public acp client")
	}

	return logic, nil
}

func (h *FDXLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	return nil, nil // nolint
}

func (h *FDXLogic) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		resp   *a2.PushedAuthorizationRequestCreated
		err    error
		scopes string
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
                     "fdx:accountdetailed:read",
                     "fdx:transactions:read",
                     "fdx:statements:read"
                  ]
               },
               {
                  "resourceType":"CUSTOMER",
                  "dataClusters":[
                     "fdx:customercontact:read"
                  ]
               }
            ]
         }
      }
   ]`

	scopes = strings.Join(h.ACPClient.Config.Scopes, " ")
	if resp, err = h.ACPClient.Oauth2.Oauth2.PushedAuthorizationRequest(
		a2.NewPushedAuthorizationRequestParams().
			WithScope(&scopes).
			WithContext(c.Request.Context()).
			WithClientID(h.ClientID).
			WithRedirectURI(h.ACPClient.Config.RedirectURL.String()).
			WithResponseType(responseType).
			WithAuthorizationDetails(&authorizationDetails),
		nil,
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

	if u, err = h.ACPClient.AuthorizeURLWithPAR(consentID); err != nil {
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

	if resp, err = h.ACPClient.Fdx.Fdx.GetFDXConsent(
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
