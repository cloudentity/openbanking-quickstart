package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/client/pagamentos"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/payments/client/domestic_payments"

	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	obbr "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obuk "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
)

type PaymentCreated struct {
	PaymentID string
	Amount    string
	Currency  string
}

func (o *OBUKClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	var (
		resp            *obuk.GetDomesticPaymentConsentRequestOK
		initiation      obModels.OBWriteDomestic2DataInitiation
		risk            obModels.OBRisk1
		createdResponse *domestic_payments.CreateDomesticPaymentsCreated
		created         PaymentCreated
		ok              bool
		err             error
	)

	if resp, ok = data.(*obuk.GetDomesticPaymentConsentRequestOK); !ok {
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
	var (
		paymentCreatedResponse *pagamentos.PaymentsPostPixPaymentsCreated
		payment                models.OpenbankingBrasilPaymentResponsePixPayment
		consent                *obbr.GetPaymentConsentOK
		ok                     bool
		err                    error
	)

	if consent, ok = data.(*obbr.GetPaymentConsentOK); !ok {
		return PaymentCreated{}, nil
	}

	if paymentCreatedResponse, err = o.Payments.Pagamentos.PaymentsPostPixPayments(
		pagamentos.NewPaymentsPostPixPaymentsParamsWithContext(c).
			WithAuthorization(accessToken).
			WithBody(&models.OpenbankingBrasilPaymentCreatePixPayment{
				Data: &models.OpenbankingBrasilPaymentCreatePixPaymentData{
					CreditorAccount: &models.OpenbankingBrasilPaymentCreditorAccount{
						AccountType: (*models.OpenbankingBrasilPaymentEnumAccountPaymentsType)(consent.Payload.Data.Payment.Details.CreditorAccount.AccountType),
						Ispb:        consent.Payload.Data.Payment.Details.CreditorAccount.Ispb,
						Number:      consent.Payload.Data.Payment.Details.CreditorAccount.Number,
					},
					LocalInstrument: (*models.OpenbankingBrasilPaymentEnumLocalInstrument)(consent.Payload.Data.Payment.Details.LocalInstrument),
					Payment: &models.OpenbankingBrasilPaymentPaymentPix{
						Amount:   consent.Payload.Data.Payment.Amount,
						Currency: consent.Payload.Data.Payment.Currency,
					},
				},
			}),
		nil,
	); err != nil {
		return PaymentCreated{}, errors.Wrapf(err, "failed to call pix payments endpoint")
	}

	if payment, ok = paymentCreatedResponse.Payload.(models.OpenbankingBrasilPaymentResponsePixPayment); !ok {
		return PaymentCreated{}, errors.New("failed to decode pix payment response")
	}
	return PaymentCreated{
		PaymentID: payment.Data.PaymentID,
		Amount:    payment.Data.Payment.Amount,
		Currency:  payment.Data.Payment.Currency,
	}, nil
}

func (o *CDRClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	return PaymentCreated{}, nil
}

func (o *FDXBankClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	return PaymentCreated{}, nil
}
