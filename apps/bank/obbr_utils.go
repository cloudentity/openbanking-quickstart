package main

import (
	"time"

	obbrAccountModels "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

func OBBRMapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = err.Code, models.OpenbankingBrasilResponseError{
		Errors: []*models.OpenbankingBrasilError{
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

func OBBRPermsToStringSlice(perms []obModels.OpenbankingBrasilConsentPermission1) []string {
	var slice []string
	for _, perm := range perms {
		slice = append(slice, string(perm))
	}
	return slice
}

func NewOBBRPayment(introspectionResponse *obModels.IntrospectOBBRPaymentConsentResponse, self strfmt.URI, id string) paymentModels.OpenbankingBrasilResponsePixPayment {
	now := strfmt.DateTime(time.Now())
	status := paymentModels.OpenbankingBrasilStatus1PDNG
	localInstrument := paymentModels.OpenbankingBrasilEnumLocalInstrumentMANU
	return paymentModels.OpenbankingBrasilResponsePixPayment{
		Data: &paymentModels.OpenbankingBrasilResponsePixPaymentData{
			PaymentID:            id,
			ConsentID:            *introspectionResponse.ConsentID,
			CreationDateTime:     now,
			StatusUpdateDateTime: now,
			Status:               &status,
			LocalInstrument:      &localInstrument,
			Payment: &paymentModels.OpenbankingBrasilPaymentPix{
				Amount:   introspectionResponse.Payment.Amount,
				Currency: introspectionResponse.Payment.Currency,
			},
			CreditorAccount: &paymentModels.OpenbankingBrasilCreditorAccount{},
		},
		Links: &paymentModels.OpenbankingBrasilLinks{
			Self: string(self),
		},
		Meta: &paymentModels.OpenbankingBrasilMeta{},
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
	Data []obbrAccountModels.AccountBalancesData `json:"data"`

	// links
	// Required: true
	Links *obbrAccountModels.Links `json:"links"`

	// meta
	// Required: true
	Meta *obbrAccountModels.Meta `json:"meta"`
}

func NewOBBRBalancesResponse(data []OBBRBalance, self strfmt.URI) interface{} {
	var (
		balances []obbrAccountModels.AccountBalancesData
		selfLink       = self.String()
		now            = strfmt.DateTime(time.Now())
		pages    int32 = 1
		records        = int32(len(data))
	)

	for _, b := range data {
		balances = append(balances, b.AccountBalancesData)
	}

	return OBBRBalances{
		Data: balances,
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
