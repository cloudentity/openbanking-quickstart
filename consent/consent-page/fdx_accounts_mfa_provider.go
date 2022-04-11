package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type FDXAccountAccessMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *FDXAccountAccessMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *f_d_x.GetFDXConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.OpenbankingFDX.GetFDXConsentSystem(f_d_x.NewGetFDXConsentSystemParamsWithContext(c).
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
		&obModels.GetAccountAccessConsentResponse{
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
