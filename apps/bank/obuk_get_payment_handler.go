package main

import (
	"errors"
	"strings"

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

type OBUKGetPaymentHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
}

func (h *OBUKGetPaymentHandler) SetIntrospectionResponse(c *gin.Context) error {
	var err error
	h.introspectionResponse, err = h.IntrospectPaymentsToken(c)
	return err
}

func (h *OBUKGetPaymentHandler) Validate(c *gin.Context) error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		return errors.New("token has no payments scope granted")
	}
	return nil
}

func (h *OBUKGetPaymentHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKGetPaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetPaymentHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	if len(data.Payments) == 1 {
		return data.Payments[0]
	}
	return h.MapError(c, errors.New("no payments matching id were found"))
}

func (h *OBUKGetPaymentHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	filtered := BankUserData{}

	for _, payment := range data.Payments {
		if *payment.Data.ConsentID == c.Param("DomesticPaymentId") {
			filtered.Payments = append(filtered.Payments, payment)
			return filtered
		}
	}

	return BankUserData{}
}
