package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	obbrModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type OBBRAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *OBBRAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *obbrModels.GetOBBRCustomerDataAccessConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithDefaults().
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get data access consent consent"))
		return
	}

	id = s.ConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

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

	if consent, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.Openbankingbr.AcceptOBBRCustomerDataAccessConsentSystem(
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

	if reject, err = s.Client.Openbanking.Openbankingbr.RejectOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewRejectOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&obModels.RejectConsentRequest{
				ID:               loginRequest.ID,
				LoginState:       loginRequest.State,
				Error:            "access_denied",
				ErrorDescription: "rejected",
				StatusCode:       403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return reject.Payload.RedirectTo, nil
}

var _ ConsentHandler = &OBBRAccountAccessConsentHandler{}
