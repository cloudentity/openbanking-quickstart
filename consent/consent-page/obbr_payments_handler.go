package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type OBBRPaymentConsentHandler struct {
	*Server
	ConsentTools
}

func (s *OBBRPaymentConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		balances BalanceResponse
		response *openbanking.GetOBBRCustomerPaymentConsentSystemOK
		err      error
	)

	if response, err = s.Client.Openbanking.GetOBBRCustomerPaymentConsentSystem(
		openbanking.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get payment consent"))
		return
	}

	if accounts, err = s.BankClient.GetInternalAccounts(response.Payload.Subject); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	if balances, err = s.BankClient.GetInternalBalances(response.Payload.Subject); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to load account balances"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("payment-consent.tmpl"), s.GetOBBRPaymentConsentTemplateData(loginRequest, response.Payload, accounts, balances.Data))
}

func (s *OBBRPaymentConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent  *openbanking.GetOBBRCustomerPaymentConsentSystemOK
		accept   *openbanking.AcceptOBBRCustomerPaymentConsentSystemOK
		err      error
		redirect string
	)

	if consent, err = s.Client.Openbanking.GetOBBRCustomerPaymentConsentSystem(
		openbanking.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.AcceptOBBRCustomerPaymentConsentSystem(
		openbanking.NewAcceptOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&models.AcceptConsentRequest{
				AccountIds:    []string{consent.Payload.CustomerDataAccessConsent.DebtorAccount.Number},
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
		reject *openbanking.RejectOBBRCustomerPaymentConsentSystemOK
		err    error
	)

	if reject, err = s.Client.Openbanking.RejectOBBRCustomerPaymentConsentSystem(
		openbanking.NewRejectOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&models.RejectConsentRequest{
				ID:         loginRequest.ID,
				LoginState: loginRequest.State,
				Error:      "rejected",
				StatusCode: 403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	return reject.Payload.RedirectTo, nil
}

var _ ConsentHandler = &OBBRPaymentConsentHandler{}
