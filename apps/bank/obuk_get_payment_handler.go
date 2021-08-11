package main

import (
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

func NewOBUKGetPaymentHandler(server *Server) GetEndpointLogic {
	return &OBUKGetPaymentHandler{Server: server}
}

func (h *OBUKGetPaymentHandler) SetIntrospectionResponse(c *gin.Context) error {
	var err error
	h.introspectionResponse, err = h.IntrospectPaymentsToken(c)
	return err
}

func (h *OBUKGetPaymentHandler) Validate(c *gin.Context) error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		return NewErrForbidden("token has no payments scope granted")
	}
	return nil
}

func (h *OBUKGetPaymentHandler) MapError(c *gin.Context, err error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetPaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetPaymentHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	if len(data.Payments.OBUK) == 1 {
		return data.Payments.OBUK[0]
	}

	_, err := h.MapError(c, ErrNotFound{"payment with consent id " + *data.Payments.OBUK[0].Data.ConsentID})
	return err
}

func (h *OBUKGetPaymentHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var filteredPayments Payments

	for _, payment := range data.Payments.OBUK {
		if *payment.Data.ConsentID == c.Param("DomesticPaymentId") {
			filteredPayments.OBUK = append(filteredPayments.OBUK, payment)
			return BankUserData{Payments: filteredPayments}
		}
	}

	return BankUserData{}
}
