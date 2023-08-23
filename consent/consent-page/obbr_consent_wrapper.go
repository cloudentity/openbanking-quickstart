package main

import (
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type SystemConsent interface {
	GetDebtorAccountNumber() string
	GetSubject() string
	GetAuthenticationContext() obModels.AuthenticationContext
	GetRequestedScopes() []*obModels.RequestedScope
	GetConsentID() string
	GetCreditorName() string
	GetPaymentCurrency() string
	GetPaymentAmount() string
	GetClientInfo() *obModels.ClientInfo
}

type OBBRConsentWrapper struct {
	Version
	OBBRPaymentsV1SystemConsent
	OBBRPaymentsV2SystemConsent
	OBBRPaymentsV3SystemConsent
}

func (w *OBBRConsentWrapper) SelectSpecificConsent() SystemConsent {
	switch w.Version {
	case V1:
		return w.OBBRPaymentsV1SystemConsent
	case V2:
		return w.OBBRPaymentsV2SystemConsent
	case V3:
		return w.OBBRPaymentsV3SystemConsent
	}

	return nil
}

func (w *OBBRConsentWrapper) GetDebtorAccountNumber() string {
	return w.SelectSpecificConsent().GetDebtorAccountNumber()
}

func (w *OBBRConsentWrapper) GetSubject() string {
	return w.SelectSpecificConsent().GetSubject()
}

func (w *OBBRConsentWrapper) GetAuthenticationContext() obModels.AuthenticationContext {
	return w.SelectSpecificConsent().GetAuthenticationContext()
}

func (w *OBBRConsentWrapper) GetRequestedScopes() []*obModels.RequestedScope {
	return w.SelectSpecificConsent().GetRequestedScopes()
}

func (w *OBBRConsentWrapper) GetConsentID() string {
	return w.SelectSpecificConsent().GetConsentID()
}

func (w *OBBRConsentWrapper) GetCreditorName() string {
	return w.SelectSpecificConsent().GetCreditorName()
}

func (w *OBBRConsentWrapper) GetPaymentCurrency() string {
	return w.SelectSpecificConsent().GetPaymentCurrency()
}

func (w *OBBRConsentWrapper) GetPaymentAmount() string {
	return w.SelectSpecificConsent().GetPaymentAmount()
}

func (w *OBBRConsentWrapper) GetClientInfo() *obModels.ClientInfo {
	return w.SelectSpecificConsent().GetClientInfo()
}

type OBBRPaymentsV1SystemConsent struct {
	*obbrModels.GetOBBRCustomerPaymentConsentSystemOK
}

func (c OBBRPaymentsV1SystemConsent) GetDebtorAccountNumber() string {
	if c.Payload.CustomerPaymentConsent.DebtorAccount != nil {
		return c.Payload.CustomerPaymentConsent.DebtorAccount.Number
	}
	return ""
}

func (c OBBRPaymentsV1SystemConsent) GetSubject() string {
	return c.Payload.Subject
}

func (c OBBRPaymentsV1SystemConsent) GetAuthenticationContext() obModels.AuthenticationContext {
	return c.Payload.AuthenticationContext
}
func (c OBBRPaymentsV1SystemConsent) GetRequestedScopes() []*obModels.RequestedScope {
	return c.Payload.RequestedScopes
}
func (c OBBRPaymentsV1SystemConsent) GetConsentID() string {
	return c.Payload.ConsentID
}

func (c OBBRPaymentsV1SystemConsent) GetCreditorName() string {
	return c.Payload.CustomerPaymentConsent.Creditor.Name
}

func (c OBBRPaymentsV1SystemConsent) GetPaymentCurrency() string {
	return c.Payload.CustomerPaymentConsent.Payment.Currency
}

func (c OBBRPaymentsV1SystemConsent) GetPaymentAmount() string {
	return c.Payload.CustomerPaymentConsent.Payment.Amount
}

func (c OBBRPaymentsV1SystemConsent) GetClientInfo() *obModels.ClientInfo {
	return c.Payload.ClientInfo
}

type OBBRPaymentsV2SystemConsent struct {
	*obbrModels.GetOBBRCustomerPaymentConsentSystemV2OK
}

func (c OBBRPaymentsV2SystemConsent) GetDebtorAccountNumber() string {
	if c.Payload.CustomerPaymentConsentV2.DebtorAccount != nil {
		return c.Payload.CustomerPaymentConsentV2.DebtorAccount.Number
	}
	return ""
}

func (c OBBRPaymentsV2SystemConsent) GetSubject() string {
	return c.Payload.Subject
}

func (c OBBRPaymentsV2SystemConsent) GetAuthenticationContext() obModels.AuthenticationContext {
	return c.Payload.AuthenticationContext
}
func (c OBBRPaymentsV2SystemConsent) GetRequestedScopes() []*obModels.RequestedScope {
	return c.Payload.RequestedScopes
}
func (c OBBRPaymentsV2SystemConsent) GetConsentID() string {
	return c.Payload.ConsentID
}

func (c OBBRPaymentsV2SystemConsent) GetCreditorName() string {
	return c.Payload.CustomerPaymentConsentV2.Creditor.Name
}

func (c OBBRPaymentsV2SystemConsent) GetPaymentCurrency() string {
	return c.Payload.CustomerPaymentConsentV2.Payment.Currency
}

func (c OBBRPaymentsV2SystemConsent) GetPaymentAmount() string {
	return c.Payload.CustomerPaymentConsentV2.Payment.Amount
}

func (c OBBRPaymentsV2SystemConsent) GetClientInfo() *obModels.ClientInfo {
	return c.Payload.ClientInfo
}

type OBBRPaymentsV3SystemConsent struct {
	*obbrModels.GetOBBRCustomerPaymentConsentSystemV3OK
}

func (c OBBRPaymentsV3SystemConsent) GetDebtorAccountNumber() string {
	if c.Payload.CustomerPaymentConsentV3.DebtorAccount != nil {
		return c.Payload.CustomerPaymentConsentV3.DebtorAccount.Number
	}
	return ""
}

func (c OBBRPaymentsV3SystemConsent) GetSubject() string {
	return c.Payload.Subject
}

func (c OBBRPaymentsV3SystemConsent) GetAuthenticationContext() obModels.AuthenticationContext {
	return c.Payload.AuthenticationContext
}
func (c OBBRPaymentsV3SystemConsent) GetRequestedScopes() []*obModels.RequestedScope {
	return c.Payload.RequestedScopes
}
func (c OBBRPaymentsV3SystemConsent) GetConsentID() string {
	return c.Payload.ConsentID
}

func (c OBBRPaymentsV3SystemConsent) GetCreditorName() string {
	return c.Payload.CustomerPaymentConsentV3.Creditor.Name
}

func (c OBBRPaymentsV3SystemConsent) GetPaymentCurrency() string {
	return c.Payload.CustomerPaymentConsentV3.Payment.Currency
}

func (c OBBRPaymentsV3SystemConsent) GetPaymentAmount() string {
	return c.Payload.CustomerPaymentConsentV3.Payment.Amount
}

func (c OBBRPaymentsV3SystemConsent) GetClientInfo() *obModels.ClientInfo {
	return c.Payload.ClientInfo
}
