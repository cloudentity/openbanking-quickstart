package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	cdr "github.com/cloudentity/acp-client-go/clients/cdr/client/c_o_n_s_e_n_t_p_a_g_e"
	clientmodels "github.com/cloudentity/acp-client-go/clients/cdr/models"
)

type CDRAccountAccessConsentHandler struct {
	*Server
	CDRConsentTools
}

func (s *CDRAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *cdr.GetCDRArrangementSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Cdr.Consentpage.GetCDRArrangementSystem(
		cdr.NewGetCDRArrangementSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get account access consent"))
		return
	}

	id = s.CDRConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	data := s.GetCDRAccountAccessConsentTemplateData(loginRequest, response.Payload, accounts)

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), data)
}

func (s *CDRAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *cdr.GetCDRArrangementSystemOK
		accept  *cdr.AcceptCDRArrangementSystemOK
		err     error
	)

	if consent, err = s.Client.Cdr.Consentpage.GetCDRArrangementSystem(
		cdr.NewGetCDRArrangementSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Cdr.Consentpage.AcceptCDRArrangementSystem(
		cdr.NewAcceptCDRArrangementSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&clientmodels.AcceptCDRConsentRequest{
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

	if reject, err = s.Client.Cdr.Consentpage.RejectCDRArrangementSystem(
		cdr.NewRejectCDRArrangementSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&clientmodels.RejectCDRConsentRequest{
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
