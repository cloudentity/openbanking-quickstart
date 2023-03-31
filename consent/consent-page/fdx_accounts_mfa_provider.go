package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp-client-go/clients/fdx/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obuk/models"

)

type FDXAccountAccessMFAConsentProvider struct {
	*Server
	FDXConsentTools
}

func (s *FDXAccountAccessMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *c_o_n_s_e_n_t_p_a_g_e.GetFDXConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Fdx.Consentpage.GetFDXConsentSystem(c_o_n_s_e_n_t_p_a_g_e.NewGetFDXConsentSystemParamsWithContext(c).
		WithLogin(loginRequest.ID), nil); err != nil {
		return data, err
	}

	data.ClientName = s.GetClientName(response.Payload.ClientInfo)
	data.ConsentID = response.Payload.ConsentID
	data.AuthenticationContext = response.Payload.AuthenticationContext

	return data, nil
}

func (s *FDXAccountAccessMFAConsentProvider) GetSMSBody(data MFAData, otp OTP) string {
	return fmt.Sprintf(
		"%s is requesting access to your accounts, please pre-authorize the consent %s using following code: %s to proceed.",
		data.ClientName,
		data.ConsentID,
		otp.OTP,
	)
}

func (s *FDXAccountAccessMFAConsentProvider) GetTemplateName() string {
	return s.GetTemplateNameForSpec("account-consent.tmpl")
}

func (s *FDXAccountAccessMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
	return s.GetAccessConsentTemplateData(
		loginRequest,
		&obModels.GetAccountAccessConsentResponse{ //UK model in FDX spec!?
			AccountAccessConsent: &obModels.AccountAccessConsent{
				Permissions: []string{"ReadAccountsBasic"},
			},
		},
		InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   "08080021325698",
					Name: "ACME Savings",
				},
				{
					ID:   "08080016225921",
					Name: "ACME Credit Card",
				},
			},
		},
	)
}
