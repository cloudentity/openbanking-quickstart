package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	clientmodels "github.com/cloudentity/acp-client-go/clients/obbr/models"
	"github.com/cloudentity/openbanking-quickstart/generated/obbr/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// swagger:route POST /payments/v1/pix/payments bank br createOBBRPaymentRequest
//
// create obbr  payment
//
// Security:
//
//	defaultcc: payments
//
// Responses:
//
//	201: ResponsePixPayment
//	400: ResponseError
//	401: ResponseError
//	403: ResponseError
//	404: ResponseError
//	405: ResponseError
//	406: ResponseError
//	415: ResponseError
//	422: ResponseError
//	429: ResponseError
//	500: ResponseError
type OBBRCreatePaymentHandler struct {
	*Server
	introspectionResponse *clientmodels.IntrospectOBBRPaymentConsentResponse
	request               models.OpenbankingBrasilPaymentCreatePixPayment
}

func NewOBBRCreatePaymentHandler(server *Server) CreateEndpointLogic {
	return &OBBRCreatePaymentHandler{Server: server}
}

func (h *OBBRCreatePaymentHandler) SetRequest(c *gin.Context) *Error {
	if err := json.NewDecoder(c.Request.Body).Decode(&h.request); err != nil {
		return ErrInternalServer.WithMessage("failed to decode request")
	}
	return nil
}

func (h *OBBRCreatePaymentHandler) CreateResource(c *gin.Context, sub string) (interface{}, *Error) {
	var (
		data    BankUserData
		id      = uuid.New().String()
		self    = strfmt.URI(fmt.Sprintf("http://localhost:%s/payments/v1/pix/payments/%s", strconv.Itoa(h.Config.Port), id))
		payment = NewOBBRPayment(h.introspectionResponse, self, id)
		err     error
	)

	if data, err = h.Storage.Get(sub); err != nil {
		return "", ErrInternalServer.WithMessage("failed to retrieve resource")
	}

	for _, p := range data.OBBRPayments {
		if p.Data.ConsentID == payment.Data.ConsentID {
			return payment, ErrAlreadyExists
		}
	}
	data.OBBRPayments = append(data.OBBRPayments, payment)

	if err = h.Storage.Put(sub, data); err != nil {
		return "", ErrInternalServer.WithMessage("failed to store resource")
	}

	return payment, nil
}

func (h *OBBRCreatePaymentHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBBRIntrospectPaymentsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBBRCreatePaymentHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		return ErrForbidden.WithMessage("token has no payments scope granted")
	}

	if !has(scopes, "openid") {
		return ErrForbidden.WithMessage("token has no openid scope granted")
	}

	consentDynamicScope := fmt.Sprintf("consent:%s", *h.introspectionResponse.ConsentID)
	if !has(scopes, consentDynamicScope) {
		return ErrForbidden.WithMessage("token has no " + consentDynamicScope + " scope granted")
	}

	if models.OpenbankingBrasilPaymentEnumAuthorisationStatusType(*h.introspectionResponse.Status) != models.OpenbankingBrasilPaymentEnumAuthorisationStatusTypeAUTHORISED {
		return ErrUnprocessableEntity.WithMessage("payment consent does not have status authorised")
	}

	// TODO: other request validation

	return nil
}

func (h *OBBRCreatePaymentHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRCreatePaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}
