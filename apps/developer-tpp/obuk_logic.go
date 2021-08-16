package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/gin-gonic/gin"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
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
		registerResponse *openbanking.CreateAccountAccessConsentRequestCreated
		err              error
	)

	if registerResponse, err = h.Client.Openbanking.CreateAccountAccessConsentRequest(
		openbanking.NewCreateAccountAccessConsentRequestParamsWithContext(c).
			WithTid(h.Client.TenantID).
			WithAid(h.Client.ServerID).
			WithRequest(&models.AccountAccessConsentRequest{
				Data: &models.OBReadConsent1Data{
					Permissions: c.PostFormArray("permissions"),
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
		registerResponse *openbanking.CreateAccountAccessConsentRequestCreated
		ok               bool
	)

	if registerResponse, ok = data.(*openbanking.CreateAccountAccessConsentRequestCreated); !ok {
		return ""
	}
	return registerResponse.Payload.Data.ConsentID
}

func (h *OBUKLogic) BuildLoginURL(c *gin.Context, consentID string, doRequestObjectEncryption bool) (string, acpclient.CSRF, error) {
	return h.Client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:openbanking:psd2:sca"}),
		acpclient.WithPKCE(),
	)
}
