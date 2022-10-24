package main

import (
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accounts/client/accounts"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpclient "github.com/cloudentity/acp-client-go"
	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type OBUKLogic struct {
	*Server
}

func (h *OBUKLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	var (
		accountsResp *accounts.GetAccountsOK
		err          error
	)

	if accountsResp, err = h.BankClient.OBUK.Accounts.GetAccounts(accounts.NewGetAccountsParamsWithContext(c).WithAuthorization(token), nil); err != nil {
		return nil, err
	}

	return accountsResp.Payload, nil
}

func (h *OBUKLogic) DoRequestObjectEncryption() bool {
	return false
}

func (h *OBUKLogic) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		registerResponse *obukModels.CreateAccountAccessConsentRequestCreated
		err              error
	)

	if registerResponse, err = h.Client.Openbanking.Openbankinguk.CreateAccountAccessConsentRequest(
		obukModels.NewCreateAccountAccessConsentRequestParamsWithContext(c).
			WithRequest(&obModels.AccountAccessConsentRequest{
				Data: &obModels.OBReadConsent1Data{
					Permissions:        c.PostFormArray("permissions"),
					ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24 * 30)),
				},
				Risk: map[string]interface{}{},
			}),
		nil,
	); err != nil {
		return nil, err
	}

	return registerResponse, nil
}

func (h *OBUKLogic) GetConsentID(data interface{}) string {
	var (
		registerResponse *obukModels.CreateAccountAccessConsentRequestCreated
		ok               bool
	)

	if registerResponse, ok = data.(*obukModels.CreateAccountAccessConsentRequestCreated); !ok {
		return ""
	}
	return registerResponse.Payload.Data.ConsentID
}

func (h *OBUKLogic) BuildLoginURL(c *gin.Context, consentID string, doRequestObjectEncryption bool) (string, acpclient.CSRF, error) {
	return h.Client.AuthorizeURL(
		acpclient.WithResponseType("code"),
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:openbanking:psd2:sca"}),
		acpclient.WithPKCE(),
		acpclient.WithResponseMode("jwt"),
	)
}

func (h *OBUKLogic) PostAuthenticationAction(c *gin.Context, d map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

var _ SpecLogicHandler = &OBUKLogic{}
