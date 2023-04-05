package main

import (
	"time"

	obbrAccountModels "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/models"
	"github.com/cloudentity/openbanking-quickstart/generated/obbr/consents/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/generated/obbr/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	clientmodels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

func OBBRMapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = err.Code, models.OpenbankingBrasilConsentV2ResponseError{
		Errors: []*models.OpenbankingBrasilConsentV2Error{
			{
				Detail: err.Message,
			},
		},
	}
	return
}

func NewOBBRAccountsResponse(accounts []obbrAccountModels.AccountData) obbrAccountModels.ResponseAccountList {
	accountPointers := []*obbrAccountModels.AccountData{}
	for _, account := range accounts {
		a := account
		accountPointers = append(accountPointers, &a)
	}

	return obbrAccountModels.ResponseAccountList{
		Data: accountPointers,
	}
}

func OBBRPermsToStringSlice(perms []clientmodels.OpenbankingBrasilConsentPermission1) []string {
	var slice []string
	for _, perm := range perms {
		slice = append(slice, string(perm))
	}
	return slice
}

func NewOBBRPayment(introspectionResponse *clientmodels.IntrospectOBBRPaymentConsentResponse, self strfmt.URI, id string) paymentModels.OpenbankingBrasilPaymentResponsePixPayment {
	now := strfmt.DateTime(time.Now())
	status := paymentModels.OpenbankingBrasilPaymentEnumPaymentStatusTypePDNG
	localInstrument := paymentModels.OpenbankingBrasilPaymentEnumLocalInstrumentMANU
	return paymentModels.OpenbankingBrasilPaymentResponsePixPayment{
		Data: &paymentModels.OpenbankingBrasilPaymentResponsePixPaymentData{
			PaymentID:            id,
			ConsentID:            *introspectionResponse.ConsentID,
			CreationDateTime:     now,
			StatusUpdateDateTime: now,
			Status:               &status,
			LocalInstrument:      &localInstrument,
			Payment: &paymentModels.OpenbankingBrasilPaymentPaymentPix{
				Amount:   introspectionResponse.Payment.Amount,
				Currency: introspectionResponse.Payment.Currency,
			},
			CreditorAccount: &paymentModels.OpenbankingBrasilPaymentCreditorAccount{},
		},
		Links: &paymentModels.OpenbankingBrasilPaymentLinks{
			Self: string(self),
		},
		Meta: &paymentModels.OpenbankingBrasilPaymentMeta{},
	}
}

func NewOBBRBalanceResponse(data OBBRBalance, self strfmt.URI) interface{} {
	var (
		selfLink       = self.String()
		now            = strfmt.DateTime(time.Now())
		pages    int32 = 1
		records  int32 = 1
	)

	return obbrAccountModels.ResponseAccountBalances{
		Data: &data.AccountBalancesData,
		Links: &obbrAccountModels.Links{
			Self: &selfLink,
		},
		Meta: &obbrAccountModels.Meta{
			RequestDateTime: &now,
			TotalPages:      &pages,
			TotalRecords:    &records,
		},
	}
}

// swagger:model OBBRBalances
type OBBRBalances struct {
	// data
	// Required: true
	Data []OBBRBalance `json:"data"`

	// links
	// Required: true
	Links *obbrAccountModels.Links `json:"links"`

	// meta
	// Required: true
	Meta *obbrAccountModels.Meta `json:"meta"`
}

func NewOBBRBalancesResponse(data []OBBRBalance, self strfmt.URI) interface{} {
	var (
		selfLink       = self.String()
		now            = strfmt.DateTime(time.Now())
		pages    int32 = 1
		records        = int32(len(data))
	)

	return OBBRBalances{
		Data: data,
		Links: &obbrAccountModels.Links{
			Self: &selfLink,
		},
		Meta: &obbrAccountModels.Meta{
			RequestDateTime: &now,
			TotalPages:      &pages,
			TotalRecords:    &records,
		},
	}
}
