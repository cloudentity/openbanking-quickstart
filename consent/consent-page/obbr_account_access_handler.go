package main

import (
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type OBBRAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *OBBRAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *openbanking.GetOBBRCustomerDataAccessConsentSystemOK
		err      error
	)

	if response, err = s.Client.Openbanking.GetOBBRCustomerDataAccessConsentSystem(
		openbanking.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithDefaults().
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get data access consent consent"))
		return
	}

	// TODO: converting internal accounts to ob model is currently in PR
	if accounts, err = s.BankClient.GetInternalAccounts(response.Payload.Subject); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), s.GetOBBRDataAccessConsentTemplateData(loginRequest, response.Payload, accounts))
}

func (s *OBBRAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *openbanking.GetOBBRCustomerDataAccessConsentSystemOK
		accept  *openbanking.AcceptOBBRCustomerDataAccessConsentSystemOK
		err     error
	)

	if consent, err = s.Client.Openbanking.GetOBBRCustomerDataAccessConsentSystem(
		openbanking.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.AcceptOBBRCustomerDataAccessConsentSystem(
		openbanking.NewAcceptOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
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

func (s *OBBRAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *openbanking.RejectOBBRCustomerDataAccessConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.RejectOBBRCustomerDataAccessConsentSystem(
		openbanking.NewRejectOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
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

var _ ConsentHandler = &OBBRAccountAccessConsentHandler{}
