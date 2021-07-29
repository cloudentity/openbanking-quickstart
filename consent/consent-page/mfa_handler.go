package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type MFAData struct {
	ClientName            string
	ConsentID             string
	AuthenticationContext map[string]interface{}

	// domestic payment specific
	Amount  string
	Account string
}

type MFAConsentProvider interface {
	GetMFAData(LoginRequest) (MFAData, error)
	GetSMSBody(MFAData, OTP) string
	GetTemplateName() string
	GetConsentMockData(LoginRequest) map[string]interface{}
}

func (s *Server) GetMFAConsentProvider(loginRequest LoginRequest) (MFAConsentProvider, bool) {
	var handler MFAConsentProvider

	switch loginRequest.ConsentType {
	case "domestic_payment":
		handler = &DomesticPaymentMFAConsentProvider{s, ConsentTools{Trans: s.Trans}}
	case "account_access":
		handler = &AccountAccessMFAConsentProvider{s, ConsentTools{Trans: s.Trans}}
	default:
		return nil, false
	}
	return handler, true
}

type AccountAccessMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *AccountAccessMFAConsentProvider) GetMFAData(loginRequest LoginRequest) (MFAData, error) {
	var (
		response *openbanking.GetAccountAccessConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
		openbanking.NewGetAccountAccessConsentSystemParams().
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

func (s *AccountAccessMFAConsentProvider) GetSMSBody(data MFAData, otp OTP) string {
	return fmt.Sprintf(
		"%s is requesting access to your accounts, please pre-authorize the consent %s using following code: %s to proceed.",
		data.ClientName,
		data.ConsentID,
		otp.OTP,
	)
}

func (s *AccountAccessMFAConsentProvider) GetTemplateName() string {
	return "account-consent.tmpl"
}

func (s *AccountAccessMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
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

type DomesticPaymentMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *DomesticPaymentMFAConsentProvider) GetMFAData(loginRequest LoginRequest) (MFAData, error) {
	var (
		response *openbanking.GetDomesticPaymentConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.GetDomesticPaymentConsentSystem(
		openbanking.NewGetDomesticPaymentConsentSystemParams().
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
		string(*response.Payload.DomesticPaymentConsent.Initiation.InstructedAmount.Amount),
		string(*response.Payload.DomesticPaymentConsent.Initiation.InstructedAmount.Currency),
	)
	data.Account = string(*response.Payload.DomesticPaymentConsent.Initiation.DebtorAccount.Identification)

	return data, nil
}

func (s *DomesticPaymentMFAConsentProvider) GetSMSBody(data MFAData, otp OTP) string {
	return fmt.Sprintf(
		"%s is requesting to initiate a payment of %s to %s, please pre-authorize the consent %s using following code %s to proceed.",
		data.ClientName,
		data.Amount,
		data.Account,
		data.ConsentID,
		otp.OTP,
	)
}

func (s *DomesticPaymentMFAConsentProvider) GetTemplateName() string {
	return "payment-consent.tmpl"
}

func (s *DomesticPaymentMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
	amount := models.OBActiveCurrencyAndAmountSimpleType("100")
	currency := models.ActiveOrHistoricCurrencyCode("GBP")
	creditorAccountName := "ACME Inc"
	debtorAccount := models.Identification0("08080021325698")

	return s.GetDomesticPaymentTemplateData(
		loginRequest,
		&models.GetDomesticPaymentConsentResponse{
			DomesticPaymentConsent: &models.DomesticPaymentConsent{
				OBWriteDomesticConsentResponse5Data: models.OBWriteDomesticConsentResponse5Data{
					Initiation: &models.OBWriteDomesticConsentResponse5DataInitiation{
						CreditorAccount: &models.OBWriteDomesticConsentResponse5DataInitiationCreditorAccount{
							Name: &creditorAccountName,
						},
						DebtorAccount: &models.OBWriteDomesticConsentResponse5DataInitiationDebtorAccount{
							Identification: &debtorAccount,
						},
						InstructedAmount: &models.OBWriteDomesticConsentResponse5DataInitiationInstructedAmount{
							Amount:   &amount,
							Currency: &currency,
						},
						RemittanceInformation: &models.OBWriteDomesticConsentResponse5DataInitiationRemittanceInformation{
							Reference: "FRESCO-101",
						},
					},
				},
			},
		},
		InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   string(debtorAccount),
					Name: "ACME Savings",
				},
			},
		},
		BalanceData{
			Balance: []Balance{
				{
					AccountID: string(debtorAccount),
					Amount: BalanceAmount{
						Amount:   "12000",
						Currency: "GBP",
					},
				},
			},
		},
	)
}

func (s *Server) MFAHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			r        = NewLoginRequest(c)
			provider MFAConsentProvider
			data     MFAData
			ok       bool
			valid    bool
			mobile   string
			err      error
		)

		if err = r.Validate(); err != nil {
			RenderInvalidRequestError(c, s.Trans, err)
			return
		}

		if provider, ok = s.GetMFAConsentProvider(r); !ok {
			RenderInvalidRequestError(c, s.Trans, fmt.Errorf("invalid consent type %s", r.ConsentType))
			return
		}

		if data, err = provider.GetMFAData(r); err != nil {
			RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to get authn context"))
			return
		}
		logrus.Debugf("authentication context: %+v", data.AuthenticationContext)

		claimData, ok := data.AuthenticationContext[s.Config.MFAClaim]

		if !ok {
			RenderInvalidRequestError(c, s.Trans, fmt.Errorf("user does not have %s configured", s.Config.MFAClaim))
			return
		}

		if mobile, ok = claimData.(string); !ok {
			RenderInternalServerError(c, s.Trans,
				fmt.Errorf(
					"failed to get %s from authn context: %+v, type: %T",
					s.Config.MFAClaim,
					data.AuthenticationContext,
					data.AuthenticationContext[s.Config.MFAClaim],
				),
			)
			return
		}

		action := c.PostForm("action")

		if action == "" {
			action = s.OTPHandler.GetDefaultAction()
		}

		logrus.Debugf("action: %s, mobile: %s", action, mobile)

		switch action {
		case "request", "resend":
			if err = s.OTPHandler.Send(r, provider, mobile, data); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to send sms otp"))
				return
			}

			isResend := action == "resend"
			templateData := map[string]interface{}{
				"mobile":          MaskMobile(mobile),
				"mfaConfirmation": true,
				"resend":          isResend,
				"mfaTrans": map[string]interface{}{
					"title": s.Trans.OrDefault("mfa.postRequest.title", "Additional verification required"),
					"subTitle": s.Trans.WithDataOrDefault("mfa.postRequest.subTitle",map[string]interface{}{
						"id": mobile,
					}, "Enter the security code sent to {{ .id }}"),
					"caption1": s.Trans.WithDataOrDefault("mfa.postRequest.caption1",map[string]interface{}{
						"resend": isResend,
					}, "We just {{ if .resend }}re-{{ end }}send you a verification code. Enter the code to proceed."),
					"resend": s.Trans.OrDefault("mfa.postRequest.resend", "Re-send the code"),
					"authenticationCode": s.Trans.OrDefault("mfa.postRequest.authenticationCode", "Authentication code"),
					"errorInfo": s.Trans.OrDefault("mfa.postRequest.errorInfo", "Invalid code. Try again or re-send the code."),
				},

			}

			if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to merge template data"))
				return
			}

			Render(c, provider.GetTemplateName(), templateData)
			return
		case "verify":
			otpStr := c.PostForm("otp")
			logrus.Debugf("check otp: %s", otpStr)

			if valid, err = s.OTPHandler.Verify(r, mobile, otpStr); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to validate otp"))
				return
			}

			if !valid {
				templateData := map[string]interface{}{
					"mobile":          MaskMobile(mobile),
					"mfaConfirmation": true,
					"invalid_otp":     true,
					"mfaTrans": map[string]interface{}{
						"title": s.Trans.OrDefault("mfa.postRequest.title", "Additional verification required"),
						"subTitle": s.Trans.WithDataOrDefault("mfa.postRequest.subTitle",map[string]interface{}{
							"id": mobile,
						}, "Enter the security code sent to {{ .id }}"),
						"caption1": s.Trans.WithDataOrDefault("mfa.postRequest.caption1",map[string]interface{}{
							"resend": false,
						}, "We just {{ if .resend }}re-{{ end }}send you a verification code. Enter the code to proceed."),
						"resend": s.Trans.OrDefault("mfa.postRequest.resend", "Re-send the code"),
						"authenticationCode": s.Trans.OrDefault("mfa.postRequest.authenticationCode", "Authentication code"),
						"errorInfo": s.Trans.OrDefault("mfa.postRequest.errorInfo", "Invalid code. Try again or re-send the code."),
					},
				}

				if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
					RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to merge template data"))
					return
				}

				Render(c, provider.GetTemplateName(), templateData)
				return
			}

			redirect := fmt.Sprintf("?%s", c.Request.URL.Query().Encode())
			logrus.Debugf("otp is valid, redirect: %s", redirect)

			c.Redirect(http.StatusMovedPermanently, redirect)
			return
		default:
			templateData := map[string]interface{}{
				"mobile":     MaskMobile(mobile),
				"mfaRequest": true,
				"mfaTrans": map[string]interface{}{
					"title": s.Trans.OrDefault("mfa.init.title", "Additional verification required"),
					"subTitle": s.Trans.OrDefault("mfa.init.subTitle", "To proceed, select where to send your security code"),
					"caption1": s.Trans.OrDefault("mfa.init.caption1", "By choosing text message you authorize <strong>GoBank</strong> to text and call you.", AsHTML),
					"caption2": s.Trans.OrDefault("mfa.init.caption2", "Message rates may apply"),
					"sms": s.Trans.OrDefault("mfa.init.sms", "Text message"),
					"email": s.Trans.OrDefault("mfa.init.email", "Email"),
				},
			}

			if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to validate otp"))
				return
			}

			Render(c, provider.GetTemplateName(), templateData)
		}
	}
}
