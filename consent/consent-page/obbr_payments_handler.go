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
	SystemConsentRetriever
}

func NewOBBRPaymentConsentHandler(server *Server, consentTools OBBRConsentTools, version Version) *OBBRPaymentConsentHandler {
	handler := &OBBRPaymentConsentHandler{
		Server:           server,
		OBBRConsentTools: consentTools,
	}
	switch version {
	case V1:
		handler.SystemConsentRetriever = GetOBBRPaymentsV1SystemConsent
	case V2:
		handler.SystemConsentRetriever = GetOBBRPaymentsV2SystemConsent
	case V3:
		handler.SystemConsentRetriever = GetOBBRPaymentsV3SystemConsent
	}

	return handler
}

func (s *OBBRPaymentConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		balances = BalanceResponse{}
		wrapper  OBBRConsentWrapper
		err      error
		id       string
	)

	if wrapper, err = s.SystemConsentRetriever(c, s.Client, loginRequest); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get payment consent"))
		return
	}

	id = s.OBBRConsentTools.GetInternalBankDataIdentifier(wrapper.GetSubject(), wrapper.GetAuthenticationContext())

	if accounts, err = s.BankClient.GetInternalAccounts(c, id); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get accounts from bank"))
		return
	}

	if balances, err = s.BankClient.GetInternalBalances(c, wrapper.GetSubject()); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to load account balances"))
		return
	}

	Render(c, s.GetTemplateNameForSpec("payment-consent.tmpl"), s.GetOBBRPaymentConsentTemplateData(loginRequest, wrapper, accounts, balances.Data))
}

func (s *OBBRPaymentConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		accept   *obbrModels.AcceptOBBRCustomerPaymentConsentSystemOK
		wrapper  OBBRConsentWrapper
		err      error
		redirect string
	)

	if wrapper, err = s.SystemConsentRetriever(c, s.Client, loginRequest); err != nil {
		return "", err
	}

	if accept, err = s.Client.Obbr.Consentpage.AcceptOBBRCustomerPaymentConsentSystem(
		obbrModels.NewAcceptOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
				AccountIds:    []string{wrapper.GetDebtorAccountNumber()},
				GrantedScopes: s.GrantScopes(wrapper.GetRequestedScopes()),
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

var _ ConsentHandler = &OBBRPaymentConsentHandler{}
