package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	obbrModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type OBBRPaymentConsentHandler struct {
	*Server
	ConsentTools
	version Version
}

func (s *OBBRPaymentConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
	if s.version == V1 {
		s.getConsentV1(c, loginRequest)
		return
	}
	s.getConsentV2(c, loginRequest)
}

func (s *OBBRPaymentConsentHandler) getConsentV1(c *gin.Context, loginRequest LoginRequest) {
	var (
		accounts InternalAccounts
		balances = BalanceResponse{}
		response *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		err      error
		id       string
	)

	if response, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get payment consent"))
		return
	}

	id = s.ConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

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

	if response, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerPaymentConsentSystemV2(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV2ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		RenderInternalServerError(c, s.Server.Trans, errors.Wrapf(err, "failed to get payment consent"))
		return
	}

	id = s.ConsentTools.GetInternalBankDataIdentifier(response.Payload.Subject, response.Payload.AuthenticationContext)

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
	if s.version == V1 {
		return s.confirmConsentV1(c, loginRequest)
	}
	return s.confirmConsentV2(c, loginRequest)
}

func (s *OBBRPaymentConsentHandler) confirmConsentV1(c *gin.Context, loginRequest LoginRequest) (string, error) {
	var (
		consent  *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		accept   *obbrModels.AcceptOBBRCustomerPaymentConsentSystemOK
		err      error
		redirect string
	)

	if consent, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.Openbankingbr.AcceptOBBRCustomerPaymentConsentSystem(
		obbrModels.NewAcceptOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
				AccountIds:    []string{consent.Payload.CustomerPaymentConsent.DebtorAccount.Number},
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

	if consent, err = s.Client.Openbanking.Openbankingbr.GetOBBRCustomerPaymentConsentSystemV2(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV2ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return "", err
	}

	if accept, err = s.Client.Openbanking.Openbankingbr.AcceptOBBRCustomerPaymentConsentSystem(
		obbrModels.NewAcceptOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID).
			WithAcceptConsent(&obModels.AcceptConsentRequest{
				AccountIds:    []string{consent.Payload.CustomerPaymentConsentV2.DebtorAccount.Number},
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

	if reject, err = s.Client.Openbanking.Openbankingbr.RejectOBBRCustomerPaymentConsentSystem(
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
