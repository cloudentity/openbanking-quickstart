package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels2 "github.com/cloudentity/acp-client-go/clients/obuk/models"
)

type OBBRAccountAccessMFAConsentProvider struct {
	*Server
	OBBRConsentTools
}

func (s *OBBRAccountAccessMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *obbrModels.GetOBBRCustomerDataAccessConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Obbr.Consentpage.GetOBBRCustomerDataAccessConsentSystem(
		obbrModels.NewGetOBBRCustomerDataAccessConsentSystemParamsWithContext(c).
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
		&obModels2.GetAccountAccessConsentResponse{ // UK model in BR spec!?
			AccountAccessConsent: &obModels2.AccountAccessConsent{
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
