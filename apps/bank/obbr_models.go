package main

import (
	obbrAccountModels "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
)

// swagger:parameters createOBBRPaymentRequest
type CreateOBBRPaymentRequest struct {
	RequestHeaders

	// in:body
	Request *models.OpenbankingBrasilCreatePixPayment
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
