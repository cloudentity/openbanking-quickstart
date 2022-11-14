package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type OBUKAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *OBUKAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *obukModels.GetAccountAccessConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Openbanking.Openbankinguk.GetAccountAccessConsentSystem(
		obukModels.NewGetAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get account access consent"))
		return
	}

	id = s.ConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), s.GetAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *OBUKAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *obukModels.GetAccountAccessConsentSystemOK
		accept  *obukModels.AcceptAccountAccessConsentSystemOK
		err     error
	)

	if consent, err = s.Client.Openbanking.Openbankinguk.GetAccountAccessConsentSystem(
		obukModels.NewGetAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.Openbankinguk.AcceptAccountAccessConsentSystem(
		obukModels.NewAcceptAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
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
		reject *obukModels.RejectAccountAccessConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.Openbankinguk.RejectAccountAccessConsentSystem(
		obukModels.NewRejectAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&obModels.RejectConsentRequest{
				ID:         loginRequest.ID,
				LoginState: loginRequest.State,
				Error:      "rejected",
				ErrorCause:       "consent_rejected",
				ErrorDescription: "The user rejected the authentication.",
				StatusCode: 403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return reject.Payload.RedirectTo, nil
}

var _ ConsentHandler = &OBUKAccountAccessConsentHandler{}
