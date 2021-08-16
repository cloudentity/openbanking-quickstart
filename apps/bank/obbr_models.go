package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
)

// swagger:parameters createOBBRPaymentRequest
type CreateOBBRPaymentRequest struct {
	RequestHeaders

	// in:body
	Request *models.OpenbankingBrasilCreatePixPayment
}
