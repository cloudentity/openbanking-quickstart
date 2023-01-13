package main

import (
	"fmt"
	"strconv"

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

	if consent, err = s.Client.Openbanking.Fdx.GetFDXConsentSystem(
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

	if accept, err = s.Client.Openbanking.Fdx.AcceptFDXConsentSystem(
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

	if reject, err = s.Client.Openbanking.Fdx.RejectFDXConsentSystem(
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
