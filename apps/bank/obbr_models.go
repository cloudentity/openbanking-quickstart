package main

import (
	obbrAccountModels "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/models"
	"github.com/cloudentity/openbanking-quickstart/generated/obbr/payments/models"
)

// swagger:parameters createOBBRPaymentRequest
type CreateOBBRPaymentRequest struct {
	RequestHeaders

	// in:body
	Request *models.OpenbankingBrasilPaymentCreatePixPayment
}

// swagger:parameters getBalancesRequest
type GetOBBRBalanceRequest struct {
	RequestHeaders

	// in:path
	AccountID string `json:"accountID"`
}

type OBBRBalance struct {
	obbrAccountModels.AccountBalancesData
	AccountID string `json:"accountId"`
}
