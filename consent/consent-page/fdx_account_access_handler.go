package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	fdx "github.com/cloudentity/acp-client-go/clients/fdx/client/c_o_n_s_e_n_t_p_a_g_e"
	"github.com/cloudentity/acp-client-go/clients/fdx/models"
)

type FDXAccountAccessConsentHandler struct {
	*Server
	FDXConsentTools
}

func (s *FDXAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *fdx.GetFDXConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Fdx.Consentpage.GetFDXConsentSystem(
		fdx.NewGetFDXConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get account access consent"))
		return
	}

	id = s.FDXConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
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

	if consent, err = s.Client.Fdx.Consentpage.GetFDXConsentSystem(
		fdx.NewGetFDXConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	grantedResources := []*models.FDXResource{}

	dataClusters := []string{}
	for _, r := range consent.Payload.FdxConsent.Resources {
		if r.ResourceType == "ACCOUNT" {
			dataClusters = r.DataClusters
			break
		}
	}

	// accept ACCOUNT resources based on user account selection
	for _, a := range c.PostFormArray("account_ids") {
		grantedResources = append(grantedResources, &models.FDXResource{
			DataClusters: dataClusters,
			ID:           a,
			ResouceType:  "ACCOUNT",
		})
	}

	// accept other resources types
	for i, r := range consent.Payload.FdxConsent.Resources {
		if r.ResourceType != "ACCOUNT" {
			grantedResources = append(grantedResources, &models.FDXResource{
				DataClusters: r.DataClusters,
				ID:           fmt.Sprintf("id-%s", strconv.Itoa(i)),
				ResouceType:  r.ResourceType,
			})
		}
	}

	if accept, err = s.Client.Fdx.Consentpage.AcceptFDXConsentSystem(
		fdx.NewAcceptFDXConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&models.AcceptFDXConsentRequest{
				GrantedScopes: s.GrantScopes(consent.Payload.RequestedScopes),
				LoginState:    loginRequest.State,
				Resources:     grantedResources,
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

	if reject, err = s.Client.Fdx.Consentpage.RejectFDXConsentSystem(
		fdx.NewRejectFDXConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&models.RejectConsentRequest{
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

var _ ConsentHandler = &FDXAccountAccessConsentHandler{}
