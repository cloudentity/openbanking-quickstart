package main

import (
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

// swagger:route GET /domestic-payments/{DomesticPaymentId} bank getDomesticPaymentRequest
//
// get domestic payment
//
// Security:
//   defaultcc: payments
//
// Responses:
//   201: OBWriteDomesticResponse5
//   403: OBErrorResponse1
//   404: OBErrorResponse1
//   500: OBErrorResponse1

// swagger:route POST /domestic-payments bank createDomesticPaymentRequest
//
// create domestic payment
//
// Security:
//   defaultcc: payments
//
// Responses:
//   201: OBWriteDomesticResponse5
//   400: OBErrorResponse1
//   403: OBErrorResponse1
//   422: OBErrorResponse1
//   500: OBErrorResponse1

type OBUKGetPaymentHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
	request               *paymentModels.OBWriteDomestic2
}

func (h *OBUKGetPaymentHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKGetPaymentHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBUKGetPaymentHandler) MapError(c *gin.Context, err error) interface{} {
	return nil
}

func (h *OBUKGetPaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return ""
}

func (h *OBUKGetPaymentHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	return nil
}

func (h *OBUKGetPaymentHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
