package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	obbrModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type OBBRPaymentMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *OBBRPaymentMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
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
		response.Payload.CustomerPaymentConsent.Payment.Amount,
		response.Payload.CustomerPaymentConsent.Payment.Currency,
	)
	data.Account = response.Payload.CustomerPaymentConsent.DebtorAccount.Number

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
