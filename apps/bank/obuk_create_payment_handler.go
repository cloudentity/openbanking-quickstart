package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

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

func (h *OBUKCreatePaymentHandler) BuildResource(c *gin.Context) interface{} {
	id := uuid.New().String()
	status := string(AcceptedSettlementInProcess)
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/domestic-payments/%s", strconv.Itoa(h.Config.Port), id))
	response := paymentModels.OBWriteDomesticResponse5{
		Data: &paymentModels.OBWriteDomesticResponse5Data{
			DomesticPaymentID:    &id,
			ConsentID:            &h.introspectionResponse.ConsentID,
			Status:               &status,
			Charges:              []*paymentModels.OBWriteDomesticResponse5DataChargesItems0{},
			CreationDateTime:     newDateTimePtr(time.Now()),
			StatusUpdateDateTime: newDateTimePtr(time.Now()),
			Initiation:           toDomesticResponse5DataInitiation(h.introspectionResponse.Initiation),
		},
		Links: &paymentModels.Links{
			Self: &self,
		},
	}

	return response
}

func (h *OBUKCreatePaymentHandler) StoreResource(c *gin.Context, sub string, resource interface{}) (interface{}, error) {
	var (
		data BankUserData
		res  paymentModels.OBWriteDomesticResponse5
		err  error
		ok   bool
	)

	if res, ok = resource.(paymentModels.OBWriteDomesticResponse5); !ok {
		return "", errors.New("wrong resource model")
	}

	if data, err = h.Storage.Get(sub); err != nil {
		return "", err
	}

	data.Payments = append(data.Payments, res)

	if err = h.Storage.Put(sub, data); err != nil {
		return "", err
	}

	return res, nil
}

func (h *OBUKCreatePaymentHandler) ResourceAlreadyExists(c *gin.Context, sub string, resource interface{}) bool {
	var (
		data BankUserData
		res  paymentModels.OBWriteDomesticResponse5
		err  error
		ok   bool
	)

	if res, ok = resource.(paymentModels.OBWriteDomesticResponse5); !ok {
		return false
	}

	if data, err = h.Storage.Get(sub); err != nil {
		return false
	}

	for _, payment := range data.Payments {
		if payment.Data.ConsentID == res.Data.ConsentID {
			return true
		}
	}

	return false
}

func (h *OBUKCreatePaymentHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKCreatePaymentHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBUKCreatePaymentHandler) MapError(c *gin.Context, err error) interface{} {
	return nil
}

func (h *OBUKCreatePaymentHandler) GetUserIdentifier(c *gin.Context) string {
	return ""
}
