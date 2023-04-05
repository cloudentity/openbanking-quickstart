package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type OBBRPaymentConsentHandler struct {
	*Server
	OBBRConsentTools
	version Version
}

func (s *OBBRPaymentConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	if s.version == V2 {
		s.getConsentV2(c, loginRequest)
		return
	}
	s.getConsentV1(c, loginRequest)
}

func (s *OBBRPaymentConsentHandler) getConsentV1(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		balances = BalanceResponse{}
		response *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get payment consent"))
		return
	}

	id = s.OBBRConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	if balances, err = s.BankClient.GetInternalBalances(c, response.Payload.Subject); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to load account balances"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("payment-consent.tmpl"), s.GetOBBRPaymentConsentTemplateData(loginRequest, response.Payload, accounts, balances.Data))
}

func (s *OBBRPaymentConsentHandler) getConsentV2(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		balances = BalanceResponse{}
		response *obbrModels.GetOBBRCustomerPaymentConsentSystemV2OK
		err      error
		id       string
	)

	if response, err = s.Client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystemV2(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV2ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get payment consent"))
		return
	}

	id = s.OBBRConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	if balances, err = s.BankClient.GetInternalBalances(c, response.Payload.Subject); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to load account balances"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("payment-consent.tmpl"), s.GetOBBRPaymentConsentTemplateDataV2(loginRequest, response.Payload, accounts, balances.Data))
}

func (s *OBBRPaymentConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	if s.version == V2 {
		return s.confirmConsentV2(c, loginRequest)
	}
	return s.confirmConsentV1(c, loginRequest)
}

func (s *OBBRPaymentConsentHandler) confirmConsentV1(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent  *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		accept   *obbrModels.AcceptOBBRCustomerPaymentConsentSystemOK
		err      error
		redirect string
	)

	if consent, err = s.Client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	wrapper := OBBRConsentWrapper{v1: consent.Payload.CustomerPaymentConsent}

	if accept, err = s.Client.Obbr.Consentpage.AcceptOBBRCustomerPaymentConsentSystem(
		obbrModels.NewAcceptOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
				AccountIds:    []string{wrapper.GetDebtorAccountNumber()},
				GrantedScopes: s.GrantScopes(consent.Payload.RequestedScopes),
				LoginState:    loginRequest.State,
			}),
		nil,
	); err != nil {
		return "", err
	}

	redirect = accept.Payload.RedirectTo

	logrus.Debugf("domestic payment consent accepted, redirect to: %s", redirect)

	return redirect, nil
}

func (s *OBBRPaymentConsentHandler) confirmConsentV2(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent  *obbrModels.GetOBBRCustomerPaymentConsentSystemV2OK
		accept   *obbrModels.AcceptOBBRCustomerPaymentConsentSystemOK
		err      error
		redirect string
	)

	if consent, err = s.Client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystemV2(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV2ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	wrapper := OBBRConsentWrapper{v2: consent.Payload.CustomerPaymentConsentV2}

	if accept, err = s.Client.Obbr.Consentpage.AcceptOBBRCustomerPaymentConsentSystem(
		obbrModels.NewAcceptOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
				AccountIds:    []string{wrapper.GetDebtorAccountNumber()},
				GrantedScopes: s.GrantScopes(consent.Payload.RequestedScopes),
				LoginState:    loginRequest.State,
			}),
		nil,
	); err != nil {
		return "", err
	}

	redirect = accept.Payload.RedirectTo

	logrus.Debugf("domestic payment consent accepted, redirect to: %s", redirect)

	return redirect, nil
}

func (s *OBBRPaymentConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject *obbrModels.RejectOBBRCustomerPaymentConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Obbr.Consentpage.RejectOBBRCustomerPaymentConsentSystem(
		obbrModels.NewRejectOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&obModels.RejectConsentRequest{
				ID:               loginRequest.ID,
				LoginState:       loginRequest.State,
				Error:            "access_denied",
				ErrorDescription: "rejected",
				StatusCode:       403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return reject.Payload.RedirectTo, nil
}

type OBBRConsentWrapper struct {
	v1 *obModels.BrazilCustomerPaymentConsent
	v2 *obModels.BrazilCustomerPaymentConsentV2
}

func (w *OBBRConsentWrapper) GetDebtorAccountNumber() string {
	if w.v1 != nil && w.v1.DebtorAccount != nil {
		return w.v1.DebtorAccount.Number
	}

	if w.v2 != nil && w.v2.DebtorAccount != nil {
		return w.v2.DebtorAccount.Number
	}

	return "N/A"
}

var _ ConsentHandler = &OBBRPaymentConsentHandler{}
