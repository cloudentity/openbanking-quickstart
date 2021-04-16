package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type AccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *AccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *openbanking.GetAccountAccessConsentSystemOK
		err      error
	)

	if response, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
		openbanking.NewGetAccountAccessConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, errors.Wrapf(err, "failed to get account access consent"))
		return
	}

	if accounts, err = s.BankClient.GetInternalAccounts(response.Payload.Subject); err != nil {
		RenderInternalServerError(c, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	Render(c, "account-consent.tmpl", s.GetAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *AccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *openbanking.GetAccountAccessConsentSystemOK
		accept  *openbanking.AcceptAccountAccessConsentSystemOK
		err     error
	)

	if consent, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
		openbanking.NewGetAccountAccessConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.AcceptAccountAccessConsentSystem(
		openbanking.NewAcceptAccountAccessConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID).
			WithAcceptAccountAccessConsent(&models.AcceptAccountAccessConsentRequest{
				GrantedScopes: s.GrantScopes(consent.Payload.RequestedScopes),
				AccountIDs:    c.PostFormArray("account_ids"),
				LoginState:    loginRequest.State,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return accept.Payload.RedirectTo, nil
}

func (s *AccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *openbanking.RejectAccountAccessConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.RejectAccountAccessConsentSystem(
		openbanking.NewRejectAccountAccessConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID).
			WithRejectAccountAccessConsent(&models.RejectAccountAccessConsentRequest{
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

var _ SpecificConsentHandler = &AccountAccessConsentHandler{}
