package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type OBUKAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *OBUKAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *openbanking.GetAccountAccessConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
		openbanking.NewGetAccountAccessConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get account access consent"))
		return
	}

	id = s.ConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), s.GetAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *OBUKAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *openbanking.GetAccountAccessConsentSystemOK
		accept  *openbanking.AcceptAccountAccessConsentSystemOK
		err     error
	)

	if consent, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
		openbanking.NewGetAccountAccessConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.AcceptAccountAccessConsentSystem(
		openbanking.NewAcceptAccountAccessConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&models.AcceptConsentRequest{
				GrantedScopes: s.GrantScopes(consent.Payload.RequestedScopes),
				AccountIds:    c.PostFormArray("account_ids"),
				LoginState:    loginRequest.State,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return accept.Payload.RedirectTo, nil
}

func (s *OBUKAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *openbanking.RejectAccountAccessConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.RejectAccountAccessConsentSystem(
		openbanking.NewRejectAccountAccessConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&models.RejectConsentRequest{
				ID:         loginRequest.ID,
				LoginState: loginRequest.State,
				Error:      "rejected",
				StatusCode: 403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return reject.Payload.RedirectTo, nil
}

var _ ConsentHandler = &OBUKAccountAccessConsentHandler{}
