package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	fdx "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
	"github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type FDXAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *FDXAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *fdx.GetFDXConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Openbanking.Fdx.GetFDXConsentSystem(
		fdx.NewGetFDXConsentSystemParamsWithContext(c).
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

	data := s.GetFDXAccountAccessConsentTemplateData(loginRequest, response.Payload, accounts)

	Render(c, s.GetTemplateNameForSpec("account-consent.tmpl"), data)
}

func (s *FDXAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent *fdx.GetFDXConsentSystemOK
		accept  *fdx.AcceptFDXConsentSystemOK
		err     error
	)

	if consent, err = s.Client.Openbanking.Fdx.GetFDXConsentSystem(
		fdx.NewGetFDXConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.Fdx.AcceptFDXConsentSystem(
		fdx.NewAcceptFDXConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&models.AcceptFDXConsentRequest{
				GrantedScopes: s.GrantScopes(consent.Payload.RequestedScopes),
				// AccountIds:    c.PostFormArray("account_ids"),// TODO resources
				LoginState: loginRequest.State,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return accept.Payload.RedirectTo, nil
}

func (s *FDXAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *fdx.RejectFDXConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.Fdx.RejectFDXConsentSystem(
		fdx.NewRejectFDXConsentSystemParamsWithContext(c).
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

var _ ConsentHandler = &FDXAccountAccessConsentHandler{}
