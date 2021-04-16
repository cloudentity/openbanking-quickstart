package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type DomesticPaymentConsentHandler struct {
	*Server
	ConsentTools
}

func (s *DomesticPaymentConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *openbanking.GetDomesticPaymentConsentSystemOK
		balances BalanceResponse
		err      error
	)

	if response, err = s.Client.Openbanking.GetDomesticPaymentConsentSystem(
		openbanking.NewGetDomesticPaymentConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, errors.Wrapf(err, "failed to get domestic payment consent"))
		return
	}

	if accounts, err = s.BankClient.GetInternalAccounts(response.Payload.Subject); err != nil {
		RenderInternalServerError(c, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	if balances, err = s.BankClient.GetInternalBalances(response.Payload.Subject); err != nil {
		RenderInternalServerError(c, errors.Wrapf(err, "failed to load account balances"))
		return
	}

	Render(c, "payment-consent.tmpl", s.GetDomesticPaymentTemplateData(loginRequest, response.Payload, accounts, balances.Data))
}

func (s *DomesticPaymentConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent  *openbanking.GetDomesticPaymentConsentSystemOK
		accept   *openbanking.AcceptDomesticPaymentConsentSystemOK
		err      error
		redirect string
	)

	if consent, err = s.Client.Openbanking.GetDomesticPaymentConsentSystem(
		openbanking.NewGetDomesticPaymentConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.AcceptDomesticPaymentConsentSystem(
		openbanking.NewAcceptDomesticPaymentConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID).
			WithAcceptDomesticPaymentConsent(&models.AcceptDomesticPaymentConsentRequest{
				AccountIDs:    []string{*consent.Payload.Initiation.DebtorAccount.Identification},
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

func (s *DomesticPaymentConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject   *openbanking.RejectDomesticPaymentConsentSystemOK
		redirect string
		err      error
	)

	if reject, err = s.Client.Openbanking.RejectDomesticPaymentConsentSystem(
		openbanking.NewRejectDomesticPaymentConsentSystemParams().
			WithTid(s.Client.TenantID).
			WithLoginID(loginRequest.ID).
			WithRejectDomesticPaymentConsent(&models.RejectDomesticPaymentConsentRequest{
				ID:         loginRequest.ID,
				LoginState: loginRequest.State,
				Error:      "rejected",
				StatusCode: 403,
			}),
		nil,
	); err != nil {
		return "", err
	}

	redirect = reject.Payload.RedirectTo

	logrus.Debugf("domestic payment consent denied, redirect to: %s", redirect)

	return redirect, nil
}

var _ SpecificConsentHandler = &DomesticPaymentConsentHandler{}
