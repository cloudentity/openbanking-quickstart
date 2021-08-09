package main

import (
	"fmt"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
)

type OBBRAccountAccessMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *OBBRAccountAccessMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *openbanking.GetOBBRCustomerDataAccessConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.GetOBBRCustomerDataAccessConsentSystem(
		openbanking.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return data, err
	}

	data.ClientName = s.GetClientName(response.Payload.ClientInfo)
	data.ConsentID = response.Payload.ConsentID
	data.AuthenticationContext = response.Payload.AuthenticationContext

	return data, nil
}

func (s *OBBRAccountAccessMFAConsentProvider) GetSMSBody(data MFAData, otp OTP) string {
	return fmt.Sprintf(
		"%s is requesting access to your accounts, please pre-authorize the consent %s using following code: %s to proceed.",
		data.ClientName,
		data.ConsentID,
		otp.OTP,
	)
}

func (s *OBBRAccountAccessMFAConsentProvider) GetTemplateName() string {
	return s.GetTemplateNameForSpec("account-consent.tmpl")
}

func (s *OBBRAccountAccessMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
	return s.GetAccessConsentTemplateData(
		loginRequest,
		&models.GetAccountAccessConsentResponse{
			AccountAccessConsent: &models.AccountAccessConsent{
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
