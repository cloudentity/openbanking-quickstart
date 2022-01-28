package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/clients/openbanking/client/cdr"
	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type CDRAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *CDRAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *cdr.GetCDRArrangementSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Openbanking.Cdr.GetCDRArrangementSystem(
		cdr.NewGetCDRArrangementSystemParamsWithContext(c).
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

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), s.GetCDRAccountAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *CDRAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *cdr.GetCDRArrangementSystemOK
		accept  *cdr.AcceptCDRArrangementSystemOK
		err     error
	)

	if consent, err = s.Client.Openbanking.Cdr.GetCDRArrangementSystem(
		cdr.NewGetCDRArrangementSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.Cdr.AcceptCDRArrangementSystem(
		cdr.NewAcceptCDRArrangementSystemParamsWithContext(c).
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

func (s *CDRAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *cdr.RejectCDRArrangementSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.Cdr.RejectCDRArrangementSystem(
		cdr.NewRejectCDRArrangementSystemParamsWithContext(c).
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

var _ ConsentHandler = &CDRAccountAccessConsentHandler{}
