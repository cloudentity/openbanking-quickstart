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
		wrapper OBBRConsentWrapper
		data    = MFAData{}
		err     error
	)

	if wrapper, err = GetOBBRPaymentsSystemConsent(c, s.Client, loginRequest); err != nil {
		return data, err
	}

	data.ConsentID = wrapper.GetConsentID()
	data.AuthenticationContext = wrapper.GetAuthenticationContext()
	data.ClientName = s.GetClientName(wrapper.GetClientInfo())
	data.Amount = fmt.Sprintf(
		"%s%s",
		wrapper.GetPaymentAmount(),
		wrapper.GetPaymentCurrency(),
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

	mockResponse := &obbrModels.GetOBBRCustomerPaymentConsentSystemOK{
		Payload: &obModels.GetOBBRCustomerPaymentConsentResponse{
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
	}

	return s.GetOBBRPaymentConsentTemplateData(
		loginRequest,
		OBBRConsentWrapper{
			GetOBBRCustomerPaymentConsentSystemOK: mockResponse,
			SystemConsent:                         OBBRPaymentsV1SystemConsent{mockResponse.Payload.CustomerPaymentConsent},
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
