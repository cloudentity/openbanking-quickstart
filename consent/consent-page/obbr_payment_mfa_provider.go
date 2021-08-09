package main

import (
	"fmt"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
)

type OBBRPaymentMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *OBBRPaymentMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *openbanking.GetOBBRCustomerPaymentConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.GetOBBRCustomerPaymentConsentSystem(
		openbanking.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return data, err
	}

	data.ConsentID = response.Payload.ConsentID
	data.AuthenticationContext = response.Payload.AuthenticationContext
	data.ClientName = s.GetClientName(response.Payload.ClientInfo)
	data.Amount = fmt.Sprintf(
		"%s%s",
		string(response.Payload.CustomerDataAccessConsent.Payment.Amount),
		string(response.Payload.CustomerDataAccessConsent.Payment.Currency),
	)
	data.Account = string(response.Payload.CustomerDataAccessConsent.DebtorAccount.Number)

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
		&models.GetOBBRCustomerPaymentConsentResponse{

			CustomerDataAccessConsent: &models.OBBRCustomerPaymentConsent{
				Creditor: &models.OpenbankingBrasilIdentification{
					Name: "ACME Inc",
				},
				DebtorAccount: &models.OpenbankingBrasilDebtorAccount{
					Number: account,
				},
				Payment: &models.OpenbankingBrasilPaymentConsent{
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
