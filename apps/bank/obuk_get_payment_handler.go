package main

import (
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

// swagger:route GET /domestic-payments/{DomesticPaymentId} bank uk getDomesticPaymentRequest
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

func (h *OBUKGetPaymentHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBUKIntrospectPaymentsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBUKGetPaymentHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		return ErrForbidden.WithMessage("token has no payments scope granted")
	}
	return nil
}

func (h *OBUKGetPaymentHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetPaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetPaymentHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	if len(data.OBUKPayments) == 1 {
		return data.OBUKPayments[0], nil
	}

	_, err := h.MapError(c, ErrNotFound.WithMessage("payment with consent id "+*data.OBUKPayments[0].Data.ConsentID))
	return err, nil
}

func (h *OBUKGetPaymentHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var filteredPayments []models.OBWriteDomesticResponse5

	for _, payment := range data.OBUKPayments {
		if *payment.Data.ConsentID == c.Param("DomesticPaymentId") {
			filteredPayments = append(filteredPayments, payment)
			return BankUserData{OBUKPayments: filteredPayments}
		}
	}

	return BankUserData{}
}
