package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type DomesticPaymentStatus string

const (
	AcceptedSettlementInProcess DomesticPaymentStatus = "AcceptedSettlementInProcess"
	AcceptedSettlementCompleted DomesticPaymentStatus = "AcceptedSettlementCompleted"
)

// swagger:parameters createDomesticPaymentRequest
type CreateDomesticPaymentRequest struct {
	RequestHeaders

	// in:body
	Request *paymentModels.OBWriteDomestic2
}

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
func (s *Server) CreateDomesticPayment() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
			paymentRequest        *paymentModels.OBWriteDomestic2
			err                   error
			errAlreadyExists      ErrAlreadyExists
		)

		if err = json.NewDecoder(c.Request.Body).Decode(&paymentRequest); err != nil {
			msg := "unable to decode domestic payments request object"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		if introspectionResponse, err = s.IntrospectPaymentsToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		scopes := strings.Split(introspectionResponse.Scope, " ")
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
		}

		id := uuid.New().String()
		status := string(AcceptedSettlementInProcess)
		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/domestic-payments/%s", strconv.Itoa(s.Config.Port), id))
		response := paymentModels.OBWriteDomesticResponse5{
			Data: &paymentModels.OBWriteDomesticResponse5Data{
				DomesticPaymentID:    &id,
				ConsentID:            introspectionResponse.ConsentID,
				Status:               &status,
				Charges:              []*paymentModels.OBWriteDomesticResponse5DataChargesItems0{},
				CreationDateTime:     newDateTimePtr(time.Now()),
				StatusUpdateDateTime: newDateTimePtr(time.Now()),
				Initiation:           toDomesticResponse5DataInitiation(introspectionResponse.Initiation),
			},
			Links: &paymentModels.Links{
				Self: &self,
			},
		}

		// create resource
		if err = s.Storage.CreateDomesticPayment(introspectionResponse.Sub, response); err != nil {
			msg := err.Error()
			if errors.As(err, &errAlreadyExists) {
				c.JSON(http.StatusConflict, models.OBError1{
					Message: &msg,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, models.OBError1{
				Message: &msg,
			})
			return
		}

		// add to payment queue worker
		// s.PaymentQueue.queue <- response

		c.PureJSON(http.StatusCreated, response)
	}
}

// swagger:parameters getDomesticPaymentRequest
type GetDomesticPaymentRequest struct {
	RequestHeaders

	// in:path
	DomesticPaymentID string `json:"DomesticPaymentId"`
}

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
func (s *Server) GetDomesticPayment() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			payment               paymentModels.OBWriteDomesticResponse5
			err                   error
			domesticPaymentID     = c.Param("DomesticPaymentId")
			introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
			errNotFound           ErrNotFound
		)

		if introspectionResponse, err = s.IntrospectPaymentsToken(c); err != nil {
			return
		}

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "payments") {
			msg := "token has no payments scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if payment, err = s.Storage.GetDomesticPayment(introspectionResponse.Sub, domesticPaymentID); err != nil {
			if errors.As(err, &errNotFound) {
				msg := "domestic payment id not found"
				c.JSON(http.StatusNotFound, models.OBErrorResponse1{
					Message: &msg,
				})
				return
			}
			msg := "retrieving domestic payment id failed"
			c.JSON(http.StatusInternalServerError, models.OBError1{
				Message: &msg,
			})
			return
		}

		c.PureJSON(http.StatusOK, payment)
	}
}

func has(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func newDateTimePtr(t time.Time) *strfmt.DateTime {
	str := strfmt.DateTime(t)
	return &str
}

func initiationsAreEqual(initiation1, initiation2 interface{}) bool {
	var (
		initiation1Bytes []byte
		initiation2Bytes []byte
		err              error
	)
	if initiation1Bytes, err = json.Marshal(initiation1); err != nil {
		return false
	}
	if initiation2Bytes, err = json.Marshal(initiation2); err != nil {
		return false
	}
	return bytes.Equal(initiation1Bytes, initiation2Bytes)
}

func toDomesticResponse5DataInitiation(initiation *acpClient.OBWriteDomesticConsentResponse5DataInitiation) *paymentModels.OBWriteDomesticResponse5DataInitiation {
	var (
		initiationBytes []byte
		err             error
		ret             paymentModels.OBWriteDomesticResponse5DataInitiation
	)

	if initiationBytes, err = json.Marshal(*initiation); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(initiationBytes, &ret); err != nil {
		panic(err)
	}

	return &ret
}
