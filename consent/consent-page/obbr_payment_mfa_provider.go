package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type OBBRPaymentMFAConsentProvider struct {
	*Server
	OBBRConsentTools
}

func (s *OBBRPaymentMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return data, err
	}

	wrapper := OBBRConsentWrapper{v1: response.Payload.CustomerPaymentConsent}

	data.ConsentID = response.Payload.ConsentID
	data.AuthenticationContext = response.Payload.AuthenticationContext
	data.ClientName = s.GetClientName(response.Payload.ClientInfo)
	data.Amount = fmt.Sprintf(
		"%s%s",
		response.Payload.CustomerPaymentConsent.Payment.Amount,
		response.Payload.CustomerPaymentConsent.Payment.Currency,
	)
	data.Account = wrapper.GetDebtorAccountNumber()

	return data, nil
}

func (s *OBBRPaymentMFAConsentProvider) GetSMSBody(data MFAData, otp OTP) string {
	return fmt.Sprintf(
		"%s is requesting to initiate a payment of %s to %s, please pre-authorize the consent %s using following code %s to proceed.",
		data.ClientName,
		data.Amount,
		data.Account,
		data.ConsentID,
		otp.OTP,
	)
}

func (s *OBBRPaymentMFAConsentProvider) GetTemplateName() string {
	return s.GetTemplateNameForSpec("payment-consent.tmpl")
}

func (s *OBBRPaymentMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
	account := "08080021325698"

	return s.GetOBBRPaymentConsentTemplateData(
		loginRequest,
		&obModels.GetOBBRCustomerPaymentConsentResponse{
			CustomerPaymentConsent: &obModels.BrazilCustomerPaymentConsent{
				Creditor: &obModels.OpenbankingBrasilPaymentIdentification{
					Name: "ACME Inc",
				},
				DebtorAccount: &obModels.OpenbankingBrasilPaymentDebtorAccount{
					Number: account,
				},
				Payment: &obModels.OpenbankingBrasilPaymentPaymentConsent{
					Currency: "BRL",
					Amount:   "100",
				},
			},
		},
		InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   account,
					Name: "ACME Savings",
				},
			},
		},
		BalanceData{
			Balance: []Balance{
				{
					AccountID: account,
					Amount: BalanceAmount{
						Amount:   "12000",
						Currency: "GBP",
					},
				},
			},
		},
	)
}
