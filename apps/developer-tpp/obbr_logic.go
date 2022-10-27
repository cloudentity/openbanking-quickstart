package main

import (
	"time"

	"github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/client/accounts"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	acpclient "github.com/cloudentity/acp-client-go"
	obbrModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
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
		registerResponse *obbrModels.CreateDataAccessConsentCreated
		perms            []obModels.OpenbankingBrasilConsentPermission
		err              error
	)

	for _, p := range c.PostFormArray("permissions") {
		perms = append(perms, obModels.OpenbankingBrasilConsentPermission(p))
	}

	if registerResponse, err = h.Client.Openbanking.Openbankingbr.CreateDataAccessConsent(
		obbrModels.NewCreateDataAccessConsentParamsWithContext(c).
			WithRequest(&obModels.BrazilCustomerDataAccessConsentRequest{
				Data: &obModels.OpenbankingBrasilConsentData{
					ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24)),
					BusinessEntity: &obModels.OpenbankingBrasilConsentBusinessEntity{
						Document: &obModels.OpenbankingBrasilConsentDocument1{
							Identification: "11111111111111",
							Rel:            "CNPJ",
						},
					},
					LoggedUser: &obModels.OpenbankingBrasilConsentLoggedUser{
						Document: &obModels.OpenbankingBrasilConsentDocument{
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
		registerResponse *obbrModels.CreateDataAccessConsentCreated
		ok               bool
	)

	if registerResponse, ok = data.(*obbrModels.CreateDataAccessConsentCreated); !ok {
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
			acpclient.WithResponseType("code"),
			acpclient.WithOpenbankingIntentID(consentID, []string{"urn:brasil:openbanking:loa2"}),
			acpclient.WithRequestObjectEncryption(key),
			acpclient.WithPKCE(),
			acpclient.WithResponseMode("jwt"),
		)
	}

	return client.AuthorizeURL(
		acpclient.WithResponseType("code"),
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:brasil:openbanking:loa2"}),
		acpclient.WithPKCE(),
		acpclient.WithResponseMode("jwt"),
	)
}

func (h *OBBRLogic) PostAuthenticationAction(c *gin.Context, d map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

var _ SpecLogicHandler = &OBBRLogic{}
