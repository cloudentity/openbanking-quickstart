package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	obuk "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
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
	GetMFAData(*gin.Context, LoginRequest) (MFAData, error)
	GetSMSBody(MFAData, OTP) string
	GetTemplateName() string
	GetConsentMockData(LoginRequest) map[string]interface{}
}

func (s *Server) GetMFAConsentProvider(loginRequest LoginRequest) (MFAConsentProvider, bool) {
	switch loginRequest.ConsentType {
	case "domestic_payment", "payments":
		return s.PaymentMFAConsentProvider, true
	case "account_access", "consents":
		return s.AccountAccessMFAConsentProvider, true
	default:
		return nil, false
	}
}

type OBUKAccountAccessMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *OBUKAccountAccessMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *obuk.GetAccountAccessConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.Openbankinguk.GetAccountAccessConsentSystem(
		obuk.NewGetAccountAccessConsentSystemParamsWithContext(c).
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

func (s *OBUKAccountAccessMFAConsentProvider) GetSMSBody(data MFAData, otp OTP) string {
	return fmt.Sprintf(
		"%s is requesting access to your accounts, please pre-authorize the consent %s using following code: %s to proceed.",
		data.ClientName,
		data.ConsentID,
		otp.OTP,
	)
}

func (s *OBUKAccountAccessMFAConsentProvider) GetTemplateName() string {
	return s.GetTemplateNameForSpec("account-consent.tmpl")
}

func (s *OBUKAccountAccessMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
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

type DomesticPaymentMFAConsentProvider struct {
	*Server
	ConsentTools
}

func (s *DomesticPaymentMFAConsentProvider) GetMFAData(c *gin.Context, loginRequest LoginRequest) (MFAData, error) {
	var (
		response *obuk.GetDomesticPaymentConsentSystemOK
		data     = MFAData{}
		err      error
	)

	if response, err = s.Client.Openbanking.Openbankinguk.GetDomesticPaymentConsentSystem(
		obuk.NewGetDomesticPaymentConsentSystemParamsWithContext(c).
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
	return s.GetTemplateNameForSpec("payment-consent.tmpl")
}

func (s *DomesticPaymentMFAConsentProvider) GetConsentMockData(loginRequest LoginRequest) map[string]interface{} {
	amount := obModels.OBActiveCurrencyAndAmountSimpleType("100")
	currency := obModels.ActiveOrHistoricCurrencyCode("GBP")
	creditorAccountName := "ACME Inc"
	debtorAccount := obModels.Identification0("08080021325698")

	return s.GetDomesticPaymentTemplateData(
		loginRequest,
		&obModels.GetDomesticPaymentConsentResponse{
			DomesticPaymentConsent: &obModels.DomesticPaymentConsent{
				OBWriteDomesticConsentResponse5Data: obModels.OBWriteDomesticConsentResponse5Data{
					Initiation: &obModels.OBWriteDomesticConsentResponse5DataInitiation{
						CreditorAccount: &obModels.OBWriteDomesticConsentResponse5DataInitiationCreditorAccount{
							Name: creditorAccountName,
						},
						DebtorAccount: &obModels.OBWriteDomesticConsentResponse5DataInitiationDebtorAccount{
							Identification: &debtorAccount,
						},
						InstructedAmount: &obModels.OBWriteDomesticConsentResponse5DataInitiationInstructedAmount{
							Amount:   &amount,
							Currency: &currency,
						},
						RemittanceInformation: &obModels.OBWriteDomesticConsentResponse5DataInitiationRemittanceInformation{
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
			r         = NewLoginRequest(c)
			provider  MFAConsentProvider
			requestId string
			data      MFAData
			ok        bool
			valid     bool
			mobile    string
			err       error
		)

		if err = r.Validate(); err != nil {
			RenderInvalidRequestError(c, s.Trans, err)
			return
		}

		if provider, ok = s.GetMFAConsentProvider(r); !ok {
			RenderInvalidRequestError(c, s.Trans, fmt.Errorf("invalid consent type %s", r.ConsentType))
			return
		}

		if data, err = provider.GetMFAData(c, r); err != nil {
			RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to get authn context"))
			return
		}

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
					"title": s.Trans.T("mfa.postRequest.title"),
					"subTitle": s.Trans.TD("mfa.postRequest.subTitle", map[string]interface{}{
						"id": mobile,
					}),
					"caption1": s.Trans.TD("mfa.postRequest.caption1", map[string]interface{}{
						"resend": isResend,
					}),
					"resend":             s.Trans.T("mfa.postRequest.resend"),
					"authenticationCode": s.Trans.T("mfa.postRequest.authenticationCode"),
					"errorInfo":          s.Trans.T("mfa.postRequest.errorInfo"),
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
						"title": s.Trans.T("mfa.postRequest.title"),
						"subTitle": s.Trans.TD("mfa.postRequest.subTitle", map[string]interface{}{
							"id": mobile,
						}),
						"caption1": s.Trans.TD("mfa.postRequest.caption1", map[string]interface{}{
							"resend": false,
						}),
						"resend":             s.Trans.T("mfa.postRequest.resend"),
						"authenticationCode": s.Trans.T("mfa.postRequest.authenticationCode"),
						"errorInfo":          s.Trans.T("mfa.postRequest.errorInfo"),
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
		case "verify_hypr":
			var devices UserDevices
			if devices, err = s.HyprHandler.GetUserDevices(s.Config.HyprUser); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to get user devices"))
				return
			}

			if len(devices) < 1 {
				RenderError(c, 401, err.Error(), errors.New("User has no registered devices with Hypr"))
				return
			}

			if requestId, err = s.HyprHandler.StartAuthentication(s.Config.HyprUser); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to start authenticate with hypr"))
				return
			}

			var checkStatus *AuthStatusResponse
			if checkStatus, err = s.HyprHandler.PollHypr(requestId); err != nil {
				if errors.Is(err, ErrTimeoutWaitingForUser) {
					RenderError(c, 401, err.Error(), errors.Wrapf(err, "timeout waiting for user to approve or deny"))
					return
				}
				RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to check auth status"))
				return
			}

			switch checkStatus.State[len(checkStatus.State)-1].Value {
			case "COMPLETED":
				s.HyprHandler.SetStorage(r, true)
				redirect := fmt.Sprintf("?%s", c.Request.URL.Query().Encode())
				logrus.Debugf("hypr is valid, redirect: %s", redirect)
				c.Redirect(http.StatusMovedPermanently, redirect)
				return
			default:
				// TODO: remove this as it is for testing only
				// s.HyprHandler.SetStorage(r, true)
				// redirect := fmt.Sprintf("?%s", c.Request.URL.Query().Encode())
				// logrus.Debugf("hypr is valid, redirect: %s", redirect)
				// c.Redirect(http.StatusMovedPermanently, redirect)
				// return
				err = errors.New("user rejected consent on hypr")
				RenderError(c, 401, err.Error(), err)
				return
			}
		default:
			templateData := map[string]interface{}{
				"mobile":     MaskMobile(mobile),
				"mfaRequest": true,
				"mfaTrans": map[string]interface{}{
					"title":    s.Trans.T("mfa.init.title"),
					"subTitle": s.Trans.T("mfa.init.subTitle"),
					"caption1": s.Trans.T("mfa.init.caption1"),
					"caption2": s.Trans.T("mfa.init.caption2"),
					"sms":      s.Trans.T("mfa.init.sms"),
					"email":    s.Trans.T("mfa.init.email"),
				},
			}

			templateData["hyprUser"] = s.Config.HyprUser

			if err = mergo.Merge(&templateData, provider.GetConsentMockData(r)); err != nil {
				RenderInternalServerError(c, s.Trans, errors.Wrap(err, "failed to validate otp"))
				return
			}

			Render(c, provider.GetTemplateName(), templateData)
		}
	}

}
