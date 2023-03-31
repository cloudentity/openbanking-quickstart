package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	obukModels "github.com/cloudentity/acp-client-go/clients/obuk/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obuk/models"
)

type OBUKAccountAccessConsentHandler struct {
	*Server
	OBUKConsentTools
}

func (s *OBUKAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *obukModels.GetAccountAccessConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Obuk.Consentpage.GetAccountAccessConsentSystem(
		obukModels.NewGetAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get account access consent"))
		return
	}

	id = s.OBUKConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

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

	if consent, err = s.Client.Obuk.Consentpage.GetAccountAccessConsentSystem(
		obukModels.NewGetAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Obuk.Consentpage.AcceptAccountAccessConsentSystem(
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

	if reject, err = s.Client.Obuk.Consentpage.RejectAccountAccessConsentSystem(
		obukModels.NewRejectAccountAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&obModels.RejectConsentRequest{
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

	return reject.Payload.RedirectTo, nil
}

var _ ConsentHandler = &OBUKAccountAccessConsentHandler{}
