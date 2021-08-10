package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	acpClient "github.com/cloudentity/acp-client-go/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
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

	// could just check if it already exists here...
	data.Payments = append(data.Payments, payment)

	/*if err = h.Storage.Put(sub, data); err != nil {
		return "", err
	}*/

	return payment, nil
}

func (h *OBUKCreatePaymentHandler) SetIntrospectionResponse(c *gin.Context) error {
	var err error
	h.introspectionResponse, err = h.IntrospectPaymentsToken(c)
	return err
}

func (h *OBUKCreatePaymentHandler) Validate(c *gin.Context) error {
	/*scopes := strings.Split(introspectionResponse.Scope, " ")
	if !has(scopes, "payments") {
		msg := "token has no payments scope granted"
		c.JSON(http.StatusForbidden, models.OBErrorResponse1{
			Message: &msg,
		})
		return
	}

	if *introspectionResponse.Status != "Authorised" {
		msg := "domestic payment consent does not have status authorised"
		c.JSON(http.StatusUnprocessableEntity, models.OBError1{
			Message: &msg,
		})
		return
	}

	if paymentRequest.Data.Initiation == nil {
		msg := "initiation data not present in request"
		c.JSON(http.StatusBadRequest, models.OBError1{
			Message: &msg,
		})
		return
	}

	if introspectionResponse.Initiation == nil {
		msg := "initiation data not present in introspection response"
		c.JSON(http.StatusInternalServerError, models.OBError1{
			Message: &msg,
		})
		return
	}

	if !initiationsAreEqual(*paymentRequest.Data.Initiation, *introspectionResponse.Initiation) {
		msg := "request initiation does not match consent initiation"
		c.JSON(http.StatusBadRequest, models.OBError1{
			Message: &msg,
		})
		return
	}

	if paymentRequest.Risk == nil {
		msg := "no risk data in payment request"
		c.JSON(http.StatusBadRequest, models.OBError1{
			Message: &msg,
		})
		return
	}

	paymentRisk := paymentRequest.Risk
	consentRisk := &paymentModels.OBRisk1{}

	if err = copier.Copy(consentRisk, introspectionResponse); err != nil {
		msg := "internal error"
		logrus.WithError(err).Error("field copying failed")
		c.JSON(http.StatusInternalServerError, models.OBError1{
			Message: &msg,
		})
		return
	}

	if !reflect.DeepEqual(paymentRisk, consentRisk) {
		msg := "risk validation failed"
		logrus.Errorf(msg)
		c.JSON(http.StatusBadRequest, models.OBError1{
			Message: &msg,
		})
		return
	}*/
	return nil
}

func (h *OBUKCreatePaymentHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKCreatePaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}
