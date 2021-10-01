package main

import (
	"encoding/json"
	"fmt"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/client/domestic_payments"
	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"

	"github.com/gin-gonic/gin"
)

type PaymentCreated struct {
	PaymentID string
	Amount    string
	Currency  string
}

func (o *OBUKClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	var (
		resp            *openbanking.GetDomesticPaymentConsentRequestOK
		initiation      obModels.OBWriteDomestic2DataInitiation
		risk            obModels.OBRisk1
		createdResponse *domestic_payments.CreateDomesticPaymentsCreated
		created         PaymentCreated
		ok              bool
		err             error
	)

	if resp, ok = data.(*openbanking.GetDomesticPaymentConsentRequestOK); !ok {
		return PaymentCreated{}, nil
	}

	if initiation, err = getInitiation(resp); err != nil {
		return created, fmt.Errorf("failed to map consent data initiation: %+v", err)
	}

	if risk, err = getRisk(resp); err != nil {
		return created, fmt.Errorf("failed to map consent risk: %+v", err)

	}

	if createdResponse, err = o.DomesticPayments.CreateDomesticPayments(domestic_payments.NewCreateDomesticPaymentsParamsWithContext(c).
		WithAuthorization(accessToken).
		WithOBWriteDomestic2Param(&obModels.OBWriteDomestic2{
			Data: &obModels.OBWriteDomestic2Data{
				ConsentID:  &resp.Payload.Data.ConsentID,
				Initiation: &initiation,
			},
			Risk: &risk,
		}), nil); err != nil {
		return created, fmt.Errorf("failed to create payment: %+v", err)
	}

	return PaymentCreated{
		PaymentID: *createdResponse.Payload.Data.DomesticPaymentID,
		Amount:    string(*createdResponse.GetPayload().Data.Initiation.InstructedAmount.Amount),
		Currency:  string(*createdResponse.GetPayload().Data.Initiation.InstructedAmount.Currency),
	}, nil
}

func (o *OBBRClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	return PaymentCreated{}, nil
}

func getInitiation(consentResponse *openbanking.GetDomesticPaymentConsentRequestOK) (pi obModels.OBWriteDomestic2DataInitiation, err error) {
	var initiationPayload []byte

	if initiationPayload, err = json.Marshal(consentResponse.Payload.Data.Initiation); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(initiationPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}

func getRisk(consentResponse *openbanking.GetDomesticPaymentConsentRequestOK) (pi obModels.OBRisk1, err error) {
	var riskPayload []byte

	if riskPayload, err = json.Marshal(consentResponse.Payload.Risk); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(riskPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}
