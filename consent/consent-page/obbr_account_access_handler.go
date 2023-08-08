package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type OBBRAccountAccessConsentHandler struct {
	*Server
	OBBRConsentTools
}

func (s *OBBRAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *obbrModels.GetOBBRCustomerDataAccessConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Obbr.Consentpage.GetOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithDefaults().
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get data access consent consent")) //nolint
		return
	}

	id = s.OBBRConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), s.GetOBBRDataAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *OBBRAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *obbrModels.GetOBBRCustomerDataAccessConsentSystemOK
		accept  *obbrModels.AcceptOBBRCustomerDataAccessConsentSystemOK
		err     error
	)

	if consent, err = s.Client.Obbr.Consentpage.GetOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Obbr.Consentpage.AcceptOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewAcceptOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
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

func (s *OBBRAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *obbrModels.RejectOBBRCustomerDataAccessConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Obbr.Consentpage.RejectOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewRejectOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
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

var _ ConsentHandler = &OBBRAccountAccessConsentHandler{}
