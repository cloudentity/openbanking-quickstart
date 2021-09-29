package main

import (
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/client/accounts"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type OBBRLogic struct {
	*Server
}

func (h *OBBRLogic) GetAccounts(c *gin.Context, token string) (interface{}, error) {
	var (
		accountsResp *accounts.AccountsGetAccountsOK
		err          error
	)

	if accountsResp, err = h.BankClient.OBBR.Accounts.AccountsGetAccounts(
		accounts.NewAccountsGetAccountsParamsWithContext(c).
			WithAuthorization(token),
		nil,
	); err != nil {
		return nil, err
	}
	return accountsResp.Payload, nil
}

func (h *OBBRLogic) CreateConsent(c *gin.Context) (interface{}, error) {
	var (
		registerResponse *openbanking.CreateDataAccessConsentCreated
		perms            []models.OpenbankingBrasilPermission
		err              error
	)

	for _, p := range c.PostFormArray("permissions") {
		perms = append(perms, models.OpenbankingBrasilPermission(p))
	}

	if registerResponse, err = h.Client.Openbanking.CreateDataAccessConsent(
		openbanking.NewCreateDataAccessConsentParamsWithContext(c).
			WithTid(h.Client.TenantID).
			WithAid(h.Client.ServerID).
			WithRequest(&models.OBBRCustomerDataAccessConsentRequest{
				Data: &models.OpenbankingBrasilData{
					ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24)),
					LoggedUser: &models.OpenbankingBrasilLoggedUser{
						Document: &models.OpenbankingBrasilDocument1{
							Identification: "11111111111",
							Rel:            "CPF",
						},
					},
					Permissions: perms,
				},
			}),
		nil,
	); err != nil {
		return nil, err
	}

	return registerResponse, nil
}

func (h *OBBRLogic) GetConsentID(data interface{}) string {
	var (
		registerResponse *openbanking.CreateDataAccessConsentCreated
		ok               bool
	)

	if registerResponse, ok = data.(*openbanking.CreateDataAccessConsentCreated); !ok {
		return ""
	}
	return registerResponse.Payload.Data.ConsentID
}

func (h *OBBRLogic) DoRequestObjectEncryption() bool {
	return true
}

func (h *OBBRLogic) BuildLoginURL(c *gin.Context, consentID string, doRequestObjectEncryption bool) (string, acpclient.CSRF, error) {
	var (
		key    jose.JSONWebKey
		client acpclient.Client
		config = h.Config.ExtendConsentScope(consentID).ClientConfig()
		err    error
	)

	if key, err = h.GetEncryptionKey(c); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to retrieve encryption key")
	}

	if client, err = acpclient.New(config); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create new acp client")
	}

	if doRequestObjectEncryption {
		return client.AuthorizeURL(
			acpclient.WithOpenbankingIntentID(consentID, []string{"urn:brasil:openbanking:loa2"}),
			acpclient.WithRequestObjectEncryption(key),
			acpclient.WithPKCE(),
		)
	}

	return client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:brasil:openbanking:loa2"}),
		acpclient.WithPKCE(),
	)
}
