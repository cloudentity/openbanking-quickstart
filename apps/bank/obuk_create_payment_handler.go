package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	paymentModels "github.com/cloudentity/openbanking-quickstart/generated/obuk/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	obukModels "github.com/cloudentity/acp-client-go/clients/obuk/client/o_b_u_k"
)

// swagger:route POST /domestic-payments bank createDomesticPaymentRequest
//
// create domestic payment
//
// Security:
//
//	defaultcc: payments
//
// Responses:
//
//	201: OBWriteDomesticResponse5
//	400: OBErrorResponse1
//	403: OBErrorResponse1
//	422: OBErrorResponse1
//	500: OBErrorResponse1
type OBUKCreatePaymentHandler struct {
	*Server
	introspectionResponse *obukModels.OpenbankingDomesticPaymentConsentIntrospectOKBody
	request               *paymentModels.OBWriteDomestic2
}

func NewOBUKCreatePaymentHandler(server *Server) CreateEndpointLogic {
	return &OBUKCreatePaymentHandler{Server: server}
}

func (h *OBUKCreatePaymentHandler) SetRequest(c *gin.Context) *Error {
	if err := json.NewDecoder(c.Request.Body).Decode(&h.request); err != nil {
		return ErrInternalServer.WithMessage("failed to decode request")
	}
	return nil
}

func (h *OBUKCreatePaymentHandler) CreateResource(c *gin.Context, sub string) (interface{}, *Error) {
	var (
		data    BankUserData
		id      = uuid.New().String()
		self    = strfmt.URI(fmt.Sprintf("http://localhost:%s/domestic-payments/%s", strconv.Itoa(h.Config.Port), id))
		payment = NewOBUKPayment(h.introspectionResponse, self, id)
		err     error
	)

	if data, err = h.Storage.Get(sub); err != nil {
		return "", ErrInternalServer.WithMessage("failed to retrieve resource")
	}

	for _, p := range data.OBUKPayments {
		if p.Data.ConsentID == payment.Data.ConsentID {
			return payment, ErrAlreadyExists
		}
	}
	data.OBUKPayments = append(data.OBUKPayments, payment)

	if err = h.Storage.Put(sub, data); err != nil {
		return "", ErrInternalServer.WithMessage("failed to store resource")
	}

	return payment, nil
}

func (h *OBUKCreatePaymentHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBUKIntrospectPaymentsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBUKCreatePaymentHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		return ErrForbidden.WithMessage("token has no payments scope granted")
	}

	if h.introspectionResponse.Status != "Authorised" {
		return ErrUnprocessableEntity.WithMessage("domestic payment consent does not have status authorised")
	}

	if h.request.Data.Initiation == nil {
		return ErrBadRequest.WithMessage("initiation data not present in request")
	}

	if h.request.Risk == nil {
		return ErrBadRequest.WithMessage("no risk data in payment request")
	}

	if h.introspectionResponse.Initiation == nil {
		return ErrInternalServer.WithMessage("initiation data not present in introspection response")
	}

	if !initiationsAreEqual(h.request.Data.Initiation, h.introspectionResponse.Initiation) {
		return ErrBadRequest.WithMessage("request initiation does not match consent initiation")
	}

	consentRisk := &paymentModels.OBRisk1{}
	if err := copier.Copy(consentRisk, h.introspectionResponse); err != nil {
		return ErrInternalServer
	}

	paymentRisk := h.request.Risk
	if !reflect.DeepEqual(paymentRisk, consentRisk) {
		return ErrBadRequest.WithMessage("risk validation failed")
	}

	return nil
}

func (h *OBUKCreatePaymentHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKCreatePaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}
