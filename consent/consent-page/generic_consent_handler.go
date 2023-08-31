package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/clients/system/client/logins"
	"github.com/cloudentity/acp-client-go/clients/system/models"
)

type GenericAccountAccessConsentHandler struct {
	*Server
	GenericConsentTools
	ConsentStorage
}

func (s *GenericAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		response *logins.GetScopeGrantRequestOK
		accounts InternalAccounts
		err      error
	)

	if response, err = s.Client.System.Logins.GetScopeGrantRequest(
		logins.NewGetScopeGrantRequestParamsWithContext(c).WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get login session"))
		return
	}

	id := s.GenericConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), s.GetAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *GenericAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		response          *logins.GetScopeGrantRequestOK
		accept            *logins.AcceptScopeGrantRequestOK
		grantedScopes     = []string{}
		externalConsentID ConsentID
		err               error
	)

	if response, err = s.Client.System.Logins.GetScopeGrantRequest(
		logins.NewGetScopeGrantRequestParamsWithContext(c).WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", errors.Wrapf(err, "failed to get login session")
	}

	if externalConsentID, err = s.Store(c, Subject(response.Payload.Subject), Data(*response.Payload)); err != nil {
		return "", errors.Wrapf(err, "failed to store consent")
	}

	for _, scp := range response.Payload.RequestedScopes {
		grantedScopes = append(grantedScopes, scp.RequestedName)
	}

	if accept, err = s.Client.System.Logins.AcceptScopeGrantRequest(
		logins.NewAcceptScopeGrantRequestParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptScopeGrant(&models.AcceptScopeGrant{
				GrantedScopes: grantedScopes,
				LoginState:    loginRequest.State,
				ConsentID:     string(externalConsentID),
			}),
		nil,
	); err != nil {
		return "", err
	}

	return accept.Payload.RedirectTo, nil
}

func (s *GenericAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		response *logins.RejectScopeGrantRequestOK
		err      error
	)

	if response, err = s.Client.System.Logins.RejectScopeGrantRequest(
		logins.NewRejectScopeGrantRequestParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectScopeGrant(&models.RejectScopeGrant{
				ID:               loginRequest.ID,
				LoginState:       loginRequest.State,
				Error:            "rejected",
				ErrorCause:       "consent_rejected",
				ErrorDescription: "The user rejected the authentication.",
				StatusCode:       403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return response.Payload.RedirectTo, nil
}

var _ ConsentHandler = &GenericAccountAccessConsentHandler{}
