package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/client/domestic_payments"
	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/client/openbanking"
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
		return created, errors.Wrapf(err, "failed to map consent data initiation")
	}

	if risk, err = getRisk(resp); err != nil {
		return created, errors.Wrapf(err, "failed to map consent risk")
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
		return created, errors.Wrapf(err, "failed to create payment")
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
