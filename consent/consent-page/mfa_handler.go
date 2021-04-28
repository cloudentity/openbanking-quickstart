package main

import (
	"fmt"
	"net/http"
	"time"

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
		handler = &DomesticPaymentMFAConsentProvider{s, ConsentTools{}}
	case "account_access":
		handler = &AccountAccessMFAConsentProvider{s, ConsentTools{}}
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
			WithLoginID(loginRequest.ID),
		nil,
	); err != nil {
		return data, err
	}

	data.ClientName = s.GetClientName(response.Payload.Client)
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
			Permissions: []string{"ReadAccountsBasic"},
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
			WithLoginID(loginRequest.ID),
		nil,
	); err != nil {
		return data, err
	}

	data.ConsentID = response.Payload.ConsentID
	data.AuthenticationContext = response.Payload.AuthenticationContext
	data.ClientName = response.Payload.Client.Name
	data.Amount = fmt.Sprintf(
		"%s%s",
		*response.Payload.Initiation.InstructedAmount.Amount,
		*response.Payload.Initiation.InstructedAmount.Currency,
	)
	data.Account = *response.Payload.Initiation.DebtorAccount.Identification

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
	var (
		amount              = "100"
		currency            = "GBP"
		creditorAccountName = "ACME Inc"
		debtorAccount       = "08080021325698"
	)

	return s.GetDomesticPaymentTemplateData(
		loginRequest,
		&models.GetDomesticPaymentConsentResponse{
			Initiation: &models.DomesticPaymentConsentDataInitiation{
				CreditorAccount: &models.DomesticPaymentConsentCreditorAccount{
					Name: &creditorAccountName,
				},
				DebtorAccount: &models.DomesticPaymentConsentDebtorAccount{
					Identification: &debtorAccount,
				},
				InstructedAmount: &models.DomesticPaymentConsentInstructedAmount{
					Amount:   &amount,
					Currency: &currency,
				},
				RemittanceInformation: &models.DomesticPaymentConsentRemittanceInformation{
					Reference: "FRESCO-101",
				},
			},
		},
		InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   debtorAccount,
					Name: "ACME Savings",
				},
			},
		},
		BalanceData{
			Balance: []Balance{
				{
					AccountID: debtorAccount,
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
			otp      OTP
			ok       bool
			valid    bool
			mobile   string
			subject  string
			err      error
		)

		if err = r.Validate(); err != nil {
			RenderInvalidRequestError(c, err)
			return
		}

		if provider, ok = s.GetMFAConsentProvider(r); !ok {
			RenderInvalidRequestError(c, fmt.Errorf("invalid consent type %s", r.ConsentType))
			return
		}

		if data, err = provider.GetMFAData(r); err != nil {
			RenderInternalServerError(c, errors.Wrapf(err, "failed to get authn context"))
			return
		}

		if mobile, ok = data.AuthenticationContext[s.Config.MobileClaim].(string); !ok {
			RenderInternalServerError(c,
				fmt.Errorf(
					"failed to get mobile from authn context: %+v, mobile claim: %s, type: %T",
					data.AuthenticationContext,
					s.Config.MobileClaim,
					data.AuthenticationContext[s.Config.MobileClaim],
				),
			)
			return
		}

		if subject, ok = data.AuthenticationContext["sub"].(string); !ok {
			RenderInternalServerError(c, fmt.Errorf("subject not retrieved from authn context"))
			return
		}

		action := c.PostForm("action")

		logrus.Debugf("action: %s, mobile: %s, subject: %s", action, mobile, subject)

		switch action {
		case "generate", "resend":
			if otp, err = s.OTPHandler.Generate(r); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to generate otp"))
				return
			}

			if err = s.OTPHandler.Store(otp); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to store otp"))
				return
			}

			if err = s.OTPHandler.Send(mobile, provider.GetSMSBody(data, otp)); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to send sms otp"))
				return
			}

			templateData := map[string]interface{}{
				"mobile":          MaskMobile(mobile),
				"mfaConfirmation": true,
				"resend":          action == "resend",
			}

			if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to merge template data"))
				return
			}

			Render(c, provider.GetTemplateName(), templateData)
			return
		case "verify":
			otpStr := c.PostForm("otp")
			logrus.Debugf("check otp: %s", otpStr)

			if valid, err = s.OTPHandler.Verify(r, otpStr); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to validate otp"))
				return
			}

			if !valid {
				templateData := map[string]interface{}{
					"mobile":          MaskMobile(mobile),
					"mfaConfirmation": true,
					"invalid_otp":     true,
				}

				if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
					RenderInternalServerError(c, errors.Wrapf(err, "failed to merge template data"))
					return
				}

				Render(c, provider.GetTemplateName(), templateData)
				return
			}

			redirect := fmt.Sprintf("?%s", c.Request.URL.Query().Encode())
			logrus.Debugf("otp is valid, redirect: %s", redirect)
			c.Redirect(http.StatusMovedPermanently, redirect)
			return

		case "verify_okta":
			var (
				verifyURL string
				oktaID    string
				pollURL   string
				status    string = "WAITING"
			)

			if s.Config.OktaUseUser {
				subject = s.Config.OktaUser
			}

			if oktaID, err = s.OktaHandler.GetOktaID(s.Config.OktaAPIToken, subject); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to get okta id"))
				return
			}

			if verifyURL, err = s.OktaHandler.GetVerifyURL(s.Config.OktaAPIToken, oktaID); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to get verify url"))
				return
			}

			if pollURL, err = s.OktaHandler.SendVerify(s.Config.OktaAPIToken, oktaID, verifyURL); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to send verify and get poll url"))
				return
			}

			for status == "WAITING" {
				if status, err = s.OktaHandler.GetVerificationStatus(s.Config.OktaAPIToken, pollURL); err != nil {
					RenderInternalServerError(c, errors.Wrapf(err, "failed to get verification status"))
					return
				}
				time.Sleep(time.Second * 2)
			}

			switch status {
			case "SUCCESS", "success":
				s.OktaHandler.SetStorage(r, true)
				redirect := fmt.Sprintf("?%s", c.Request.URL.Query().Encode())
				logrus.Debugf("okta is valid, redirect: %s", redirect)
				c.Redirect(http.StatusMovedPermanently, redirect)
				return
			case "REJECTED", "rejected":
				err = errors.New("user rejected consent on okta")
				RenderError(c, 401, err.Error(), err)
			default:
				RenderInternalServerError(c, fmt.Errorf("received a status of %s", status))
			}

		default:
			templateData := map[string]interface{}{
				"mobile":     MaskMobile(mobile),
				"mfaRequest": true,
			}

			if s.Config.EnableMFAOkta {
				var (
					oktaID  string
					hasPush bool
				)

				templateData["showOkta"] = true

				if s.Config.OktaUseUser {
					subject = s.Config.OktaUser
				}

				if oktaID, err = s.OktaHandler.GetOktaID(s.Config.OktaAPIToken, subject); err != nil {
					logrus.Debugf("unable to retrieve okta id (err: %s). Not rendering okta button...")
					templateData["showOkta"] = false
				}

				if hasPush = s.OktaHandler.HasFactorType(s.Config.OktaAPIToken, oktaID, "push"); !hasPush {
					logrus.Debugf("no factor type push configured for okta id %s. Not rendering okta button...")
					templateData["showOkta"] = false
				}

			}

			if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
				RenderInternalServerError(c, errors.Wrapf(err, "failed to validate otp"))
				return
			}

			Render(c, provider.GetTemplateName(), templateData)
		}
	}
}
