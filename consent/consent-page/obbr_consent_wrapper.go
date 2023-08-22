package main

import (
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	obModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type OBBRConsentWrapper struct {
	version Version
	v1      *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
	v2      *obbrModels.GetOBBRCustomerPaymentConsentSystemV2OK
	// v3
}

func (w *OBBRConsentWrapper) GetDebtorAccountNumber() string {
	switch w.version {
	case V1:
		if w.v1.Payload.CustomerPaymentConsent.DebtorAccount != nil {
			return w.v1.Payload.CustomerPaymentConsent.DebtorAccount.Number
		}
	case V2:
		if w.v2.Payload.CustomerPaymentConsentV2.DebtorAccount != nil {
			return w.v2.Payload.CustomerPaymentConsentV2.DebtorAccount.Number
		}
	}

	return "N/A"
}

func (w *OBBRConsentWrapper) GetSubject() string {
	switch w.version {
	case V1:
		return w.v1.Payload.Subject
	case V2:
		return w.v2.Payload.Subject
	default:
		return "N/A"
	}
}

func (w *OBBRConsentWrapper) GetAuthenticationContext() obModels.AuthenticationContext {
	switch w.version {
	case V1:
		return w.v1.Payload.AuthenticationContext
	case V2:
		return w.v2.Payload.AuthenticationContext
	default:
		return make(map[string]interface{})
	}
}

func (w *OBBRConsentWrapper) GetRequestedScopes() []*obModels.RequestedScope {
	switch w.version {
	case V1:
		return w.v1.Payload.RequestedScopes
	case V2:
		return w.v2.Payload.RequestedScopes
	default:
		return make([]*obModels.RequestedScope, 1)
	}
}

func (w *OBBRConsentWrapper) GetConsentID() string {
	switch w.version {
	case V1:
		return w.v1.Payload.ConsentID
	case V2:
		return w.v2.Payload.ConsentID
	default:
		return "N/A"
	}
}

func (w *OBBRConsentWrapper) GetCreditorName() string {
	switch w.version {
	case V1:
		return w.v1.Payload.CustomerPaymentConsent.Creditor.Name
	case V2:
		return w.v2.Payload.CustomerPaymentConsentV2.Creditor.Name
	default:
		return "N/A"
	}
}

func (w *OBBRConsentWrapper) GetPaymentCurrency() string {
	switch w.version {
	case V1:
		return w.v1.Payload.CustomerPaymentConsent.Payment.Currency
	case V2:
		return w.v2.Payload.CustomerPaymentConsentV2.Payment.Currency
	default:
		return "N/A"
	}
}

func (w *OBBRConsentWrapper) GetPaymentAmount() string {
	switch w.version {
	case V1:
		return w.v1.Payload.CustomerPaymentConsent.Payment.Amount
	case V2:
		return w.v2.Payload.CustomerPaymentConsentV2.Payment.Amount
	default:
		return "N/A"
	}
}

func (w *OBBRConsentWrapper) GetClientInfo() *obModels.ClientInfo {
	switch w.version {
	case V1:
		return w.v1.Payload.ClientInfo
	case V2:
		return w.v2.Payload.ClientInfo
	default:
		return &obModels.ClientInfo{}
	}
}
