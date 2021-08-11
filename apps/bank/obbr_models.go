package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
)

type ResponseAccountList struct {
	Data []AccountData `json:"data"`
}

// swagger:parameters createOBBRPaymentRequest
type CreateOBBRPaymentRequest struct {
	RequestHeaders

	// in:body
	Request *models.OpenbankingBrasilCreatePixPayment
}
