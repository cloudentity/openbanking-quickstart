package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	acpClient "github.com/cloudentity/acp-client-go/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

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
type OBUKCreatePaymentHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
	request               *paymentModels.OBWriteDomestic2
}

func (h *OBUKCreatePaymentHandler) SetRequest(c *gin.Context) error {
	return json.NewDecoder(c.Request.Body).Decode(&h.request)
}

func (h *OBUKCreatePaymentHandler) CreateResource(c *gin.Context, sub string) (interface{}, error) {
	var (
		data    BankUserData
		id      = uuid.New().String()
		self    = strfmt.URI(fmt.Sprintf("http://localhost:%s/domestic-payments/%s", strconv.Itoa(h.Config.Port), id))
		payment = NewOBUKPayment(h.introspectionResponse, self, id)
		err     error
	)

	if data, err = h.Storage.Get(sub); err != nil {
		return "", err
	}

	for _, p := range data.Payments.OBUK {
		if p.Data.ConsentID == payment.Data.ConsentID {
			return payment, ErrAlreadyExists{*payment.Data.ConsentID}
		}
	}
	data.Payments.OBUK = append(data.Payments.OBUK, payment)

	if err = h.Storage.Put(sub, data); err != nil {
		return "", err
	}

	return payment, nil
}

func (h *OBUKCreatePaymentHandler) SetIntrospectionResponse(c *gin.Context) error {
	var err error
	h.introspectionResponse, err = h.IntrospectPaymentsToken(c)
	return err
}

func (h *OBUKCreatePaymentHandler) Validate(c *gin.Context) error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		return errors.New("token has no payments scope granted")
	}

	if h.introspectionResponse.Status != "Authorised" {
		return errors.New("domestic payment consent does not have status authorised")
	}

	if h.request.Data.Initiation == nil {
		return errors.New("initiation data not present in request")
	}

	if h.request.Risk == nil {
		return errors.New("no risk data in payment request")
	}

	if h.introspectionResponse.Initiation == nil {
		return errors.New("initiation data not present in introspection response")

	}

	if !initiationsAreEqual(h.request.Data.Initiation, h.introspectionResponse.Initiation) {
		return errors.New("request initiation does not match consent initiation")
	}

	consentRisk := &paymentModels.OBRisk1{}
	if err := copier.Copy(consentRisk, h.introspectionResponse); err != nil {
		return errors.New("internal error")
	}

	paymentRisk := h.request.Risk
	if !reflect.DeepEqual(paymentRisk, consentRisk) {
		return errors.New("risk validation failed")
	}

	return nil
}

func (h *OBUKCreatePaymentHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKCreatePaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}
