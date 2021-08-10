package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

func OBUKMapError(err error) (int, models.OBError1) {
	msg := err.Error()
	resp := models.OBError1{
		Message: &msg,
	}

	if errors.As(err, &errNotFound) {
		return http.StatusNotFound, resp
	}
	if errors.As(err, &errInternalServer) {
		return http.StatusInternalServerError, resp
	}
	if errors.As(err, &errAlreadyExists) {
		return http.StatusConflict, resp
	}
	if errors.As(err, &errForbidden) {
		return http.StatusForbidden, resp
	}
	if errors.As(err, &errUnprocessableEntity) {
		return http.StatusUnprocessableEntity, resp
	}

	return http.StatusBadRequest, resp
}

func NewAccountsResponse(accounts []models.OBAccount6, self strfmt.URI) models.OBReadAccount6 {
	accountsPointers := make([]*models.OBAccount6, len(accounts))

	for i, a := range accounts {
		account := a
		accountsPointers[i] = &account
	}

	return models.OBReadAccount6{
		Data: &models.OBReadAccount6Data{
			Account: accountsPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(accounts)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}

func NewBalancesResponse(balances []models.OBReadBalance1DataBalanceItems0, self strfmt.URI) models.OBReadBalance1 {
	balancesPointers := make([]*models.OBReadBalance1DataBalanceItems0, len(balances))

	for i, b := range balances {
		balance := b
		balancesPointers[i] = &balance
	}

	return models.OBReadBalance1{
		Data: &models.OBReadBalance1Data{
			Balance: balancesPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(balances)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}

func NewTransactionsResponse(transactions []models.OBTransaction6, self strfmt.URI) models.OBReadTransaction6 {
	transactionPointers := []*models.OBTransaction6{}

	for _, transaction := range transactions {
		t := transaction
		transactionPointers = append(transactionPointers, &t)
	}

	return models.OBReadTransaction6{
		Data: &models.OBReadDataTransaction6{
			Transaction: transactionPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(transactions)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}

func NewOBUKPayment(introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse, self strfmt.URI, id string) paymentModels.OBWriteDomesticResponse5 {
	status := string(AcceptedSettlementInProcess)
	return paymentModels.OBWriteDomesticResponse5{
		Data: &paymentModels.OBWriteDomesticResponse5Data{
			DomesticPaymentID:    &id,
			ConsentID:            &introspectionResponse.ConsentID,
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
