package main

import paymentModels "github.com/cloudentity/openbanking-quickstart/generated/obuk/payments/models"

type RequestHeaders struct {
	// in:header
	AuthDate string `json:"x-fapi-auth-date"`
	// in:header
	CustomerIPAddress string `json:"x-fapi-customer-ip-address"`
	// in:header
	Authorization string `json:"authorization"`
	// in:header
	InteractionID string `json:"x-fapi-interaction-id"`
	// in:header
	CustomerAgent string `json:"x-customer-user-agent"`
}

// swagger:parameters getAccountsRequest
type GetAccountsRequest struct {
	RequestHeaders
}

// swagger:parameters getInternalAccountsRequest
type GetInternalAccountsRequest struct {
	// in:query
	ID string `json:"id"`
}

// swagger:parameters getBalancesRequest
type GetBalancesRequest struct {
	RequestHeaders
}

// swagger:parameters getInternalBalancesRequest
type GetInternalBalancesRequest struct {
	// in:query
	ID string `json:"id"`
}

type DomesticPaymentStatus string

const (
	AcceptedSettlementInProcess DomesticPaymentStatus = "AcceptedSettlementInProcess"
	AcceptedSettlementCompleted DomesticPaymentStatus = "AcceptedSettlementCompleted"
)

// swagger:parameters createDomesticPaymentRequest
type CreateDomesticPaymentRequest struct {
	RequestHeaders

	// in:body
	Request *paymentModels.OBWriteDomestic2
}

// swagger:parameters getDomesticPaymentRequest
type GetDomesticPaymentRequest struct {
	RequestHeaders

	// in:path
	DomesticPaymentID string `json:"DomesticPaymentId"`
}

// swagger:parameters getTransactionsRequest
type GetTransactionsRequest struct {
	RequestHeaders
}
