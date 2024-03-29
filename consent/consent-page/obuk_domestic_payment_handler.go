package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	obukModels "github.com/cloudentity/acp-client-go/clients/obuk/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obuk/models"
)

type OBUKDomesticPaymentConsentHandler struct {
	*Server
	OBUKConsentTools
}

func (s *OBUKDomesticPaymentConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		response *obukModels.GetDomesticPaymentConsentSystemOK
		balances BalanceResponse
		err      error
		id       string
	)

	if response, err = s.Client.Obuk.Consentpage.GetDomesticPaymentConsentSystem(
		obukModels.NewGetDomesticPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		s.RenderInternalServerError(c, errors.Wrapf(err, "failed to get domestic payment consent"))
		return
	}

	id = s.OBUKConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		s.RenderInternalServerError(c, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	if balances, err = s.BankClient.GetInternalBalances(c, response.Payload.Subject); err != nil {
		s.RenderInternalServerError(c, errors.Wrapf(err, "failed to load account balances"))
		return
	}

	s.Render(c, s.GetTemplateNameForSpec("payment-consent.tmpl"), s.GetDomesticPaymentTemplateData(loginRequest, response.Payload, accounts, balances.Data))
}

func (s *OBUKDomesticPaymentConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent  *obukModels.GetDomesticPaymentConsentSystemOK
		accept   *obukModels.AcceptDomesticPaymentConsentSystemOK
		err      error
		redirect string
	)

	if consent, err = s.Client.Obuk.Consentpage.GetDomesticPaymentConsentSystem(
		obukModels.NewGetDomesticPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Obuk.Consentpage.AcceptDomesticPaymentConsentSystem(
		obukModels.NewAcceptDomesticPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
				AccountIds:    []string{string(*consent.Payload.DomesticPaymentConsent.Initiation.DebtorAccount.Identification)},
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

func (s *OBUKDomesticPaymentConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		reject   *obukModels.RejectDomesticPaymentConsentSystemOK
		redirect string
		err      error
	)

	if reject, err = s.Client.Obuk.Consentpage.RejectDomesticPaymentConsentSystem(
		obukModels.NewRejectDomesticPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithRejectConsent(&obModels.RejectConsentRequest{
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

var _ ConsentHandler = &OBUKDomesticPaymentConsentHandler{}
