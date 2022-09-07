package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	fdx "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
	fdxModels "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
	// fdxPayment "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/client/payments"
)

type FDXCreatePaymentHandler struct {
	*Server
	introspectionResponse *fdx.FdxConsentIntrospectOKBody
	request               fdxModels.PaymentForUpdateentity1
}

func NewFDXCreatePaymentHandler(server *Server) CreateEndpointLogic {
	return &FDXCreatePaymentHandler{Server: server}
}

func (h *FDXCreatePaymentHandler) SetRequest(c *gin.Context) *Error {
	if err := json.NewDecoder(c.Request.Body).Decode(&h.request); err != nil {
		return ErrInternalServer.WithMessage("failed to decode request")
	}
	return nil
}

func (h *FDXCreatePaymentHandler) CreateResource(c *gin.Context, sub string) (interface{}, *Error) {
	var (
		data    BankUserData
		id      = uuid.New().String()
		self    = strfmt.URI(fmt.Sprintf("http://localhost:%s/payments/%s", strconv.Itoa(h.Config.Port), id))
		payment = NewFDXPayment(h.introspectionResponse, self, id)
		err     error
	)

	amount := 32.99
	payment.Amount = &amount // TODO remove hardcode

	if data, err = h.Storage.Get(sub); err != nil {
		return "", ErrInternalServer.WithMessage("failed to retrieve resource")
	}

	for _, p := range data.FDXPayments {
		if p.PaymentID == payment.PaymentID {
			return payment, ErrAlreadyExists
		}
	}
	data.FDXPayments = append(data.FDXPayments, payment)

	if err = h.Storage.Put(sub, data); err != nil {
		return "", ErrInternalServer.WithMessage("failed to store resource")
	}

	return payment, nil
}

func (h *FDXCreatePaymentHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.FDXIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *FDXCreatePaymentHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	log.Printf("Checking scopes in bank")
	if !has(scopes, "ACCOUNT_PAYMENTS") {
		log.Printf("Checked scopes in bank and missing account_payments")
		return ErrForbidden.WithMessage("token has no ACCOUNT_PAYMENTS scope granted")
	}

	return nil
}

func (h *FDXCreatePaymentHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXCreatePaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}
