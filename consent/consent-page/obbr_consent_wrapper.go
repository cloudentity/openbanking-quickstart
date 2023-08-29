package main

import (
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type SystemConsent interface {
	GetDebtorAccountNumber() string
	GetCreditorName() string
	GetPaymentCurrency() string
	GetPaymentAmount() string
}

type OBBRConsentWrapper struct {
	*obbrModels.GetOBBRCustomerPaymentConsentSystemOK
	SystemConsent
}

func (w *OBBRConsentWrapper) GetClientInfo() *obModels.ClientInfo {
	return w.Payload.ClientInfo
}

func (w *OBBRConsentWrapper) GetSubject() string {
	return w.Payload.Subject
}

func (w *OBBRConsentWrapper) GetAuthenticationContext() obModels.AuthenticationContext {
	return w.Payload.AuthenticationContext
}

func (w *OBBRConsentWrapper) GetRequestedScopes() []*obModels.RequestedScope {
	return w.Payload.RequestedScopes
}

func (w *OBBRConsentWrapper) GetConsentID() string {
	return w.Payload.ConsentID
}

func (w *OBBRConsentWrapper) GetDebtorAccountNumber() string {
	return w.SystemConsent.GetDebtorAccountNumber()
}

func (w *OBBRConsentWrapper) GetCreditorName() string {
	return w.SystemConsent.GetCreditorName()
}

func (w *OBBRConsentWrapper) GetPaymentCurrency() string {
	return w.SystemConsent.GetPaymentCurrency()
}

func (w *OBBRConsentWrapper) GetPaymentAmount() string {
	return w.SystemConsent.GetPaymentAmount()
}

type OBBRPaymentsV1SystemConsent struct {
	*obModels.BrazilCustomerPaymentConsent
}

func (c OBBRPaymentsV1SystemConsent) GetDebtorAccountNumber() string {
	if c.DebtorAccount != nil {
		return c.DebtorAccount.Number
	}
	return ""
}

func (c OBBRPaymentsV1SystemConsent) GetCreditorName() string {
	return c.Creditor.Name
}

func (c OBBRPaymentsV1SystemConsent) GetPaymentCurrency() string {
	return c.Payment.Currency
}

func (c OBBRPaymentsV1SystemConsent) GetPaymentAmount() string {
	return c.Payment.Amount
}

type OBBRPaymentsV2SystemConsent struct {
	*obModels.BrazilCustomerPaymentConsentV2
}

func (c OBBRPaymentsV2SystemConsent) GetDebtorAccountNumber() string {
	if c.DebtorAccount != nil {
		return c.DebtorAccount.Number
	}
	return ""
}

func (c OBBRPaymentsV2SystemConsent) GetCreditorName() string {
	return c.Creditor.Name
}

func (c OBBRPaymentsV2SystemConsent) GetPaymentCurrency() string {
	return c.Payment.Currency
}

func (c OBBRPaymentsV2SystemConsent) GetPaymentAmount() string {
	return c.Payment.Amount
}

type OBBRPaymentsV3SystemConsent struct {
	*obModels.BrazilCustomerPaymentConsentV3
}

func (c OBBRPaymentsV3SystemConsent) GetDebtorAccountNumber() string {
	if c.DebtorAccount != nil {
		return c.DebtorAccount.Number
	}
	return ""
}

func (c OBBRPaymentsV3SystemConsent) GetCreditorName() string {
	return c.Creditor.Name
}

func (c OBBRPaymentsV3SystemConsent) GetPaymentCurrency() string {
	return c.Payment.Currency
}

func (c OBBRPaymentsV3SystemConsent) GetPaymentAmount() string {
	return c.Payment.Amount
}
