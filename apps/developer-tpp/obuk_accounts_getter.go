package main

import (
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/accounts"
	"github.com/gin-gonic/gin"
)

type OBUKAccountsGetter struct {
	*Server
}

func (h *OBUKAccountsGetter) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	var (
		accountsResp *accounts.GetAccountsOK
		err          error
	)

	if accountsResp, err = h.BankClient.OBUK.Accounts.GetAccounts(accounts.NewGetAccountsParamsWithContext(c).WithAuthorization(token), nil); err != nil {
		return nil, err
	}

	return accountsResp.Payload, nil
}

type OBUKAccountConsentCreator struct {
	*Server
}

func (h *OBUKAccountConsentCreator) CreateConsent(c *gin.Context) (interface{}, error) {
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

func (h *OBUKAccountConsentCreator) GetConsentID(data interface{}) string {
	var (
		registerResponse *openbanking.CreateAccountAccessConsentRequestCreated
		ok               bool
	)

	if registerResponse, ok = data.(*openbanking.CreateAccountAccessConsentRequestCreated); !ok {
		return ""
	}
	return registerResponse.Payload.Data.ConsentID
}
