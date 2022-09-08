package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/client/pagamentos"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/client/domestic_payments"
	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
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
		consent                *obbr.GetPaymentConsentOK
		ok                     bool
		err                    error
	)

	if consent, ok = data.(*obbr.GetPaymentConsentOK); !ok {
		return PaymentCreated{}, nil
	}

	if paymentCreatedResponse, err = o.PaymentConsentsBrasil.Pagamentos.PaymentsPostPixPayments(
		pagamentos.NewPaymentsPostPixPaymentsParamsWithContext(c).
			WithAuthorization(accessToken).
			WithBody(&models.OpenbankingBrasilCreatePixPayment{
				Data: &models.OpenbankingBrasilCreatePixPaymentData{
					CreditorAccount: &models.OpenbankingBrasilCreditorAccount{
						AccountType: (*models.OpenbankingBrasilEnumAccountPaymentsType)(consent.Payload.Data.Payment.Details.CreditorAccount.AccountType),
						Ispb:        consent.Payload.Data.Payment.Details.CreditorAccount.Ispb,
						Number:      consent.Payload.Data.Payment.Details.CreditorAccount.Number,
					},
					LocalInstrument: (*models.OpenbankingBrasilEnumLocalInstrument)(consent.Payload.Data.Payment.Details.LocalInstrument),
					Payment: &models.OpenbankingBrasilPaymentPix{
						Amount:   consent.Payload.Data.Payment.Amount,
						Currency: consent.Payload.Data.Payment.Currency,
					},
				},
			}),
		nil,
	); err != nil {
		return PaymentCreated{}, errors.Wrapf(err, "failed to call pix payments endpoint")
	}

	return PaymentCreated{
		PaymentID: paymentCreatedResponse.Payload.Data.PaymentID,
		Amount:    paymentCreatedResponse.Payload.Data.Payment.Amount,
		Currency:  paymentCreatedResponse.Payload.Data.Payment.Currency,
	}, nil
}

func (o *CDRClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	return PaymentCreated{}, nil
}

func (o *FDXBankClient) CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error) {
	// TODO mocked until APIs for FDX added to bank
	return PaymentCreated{}, nil
}
