package main

import (
	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/payments/models"
	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

func OBBRMapError(c *gin.Context, err *Error) (int, interface{}) {
	return err.Code, models.OpenbankingBrasilResponseError{
		Errors: []*models.OpenbankingBrasilError{},
	}
}

/*
{
  "errors": [
    {
      "code": "string",
      "title": "string",
      "detail": "string"
    }
  ],
  "meta": {
    "totalRecords": 1,
    "totalPages": 1,
    "requestDateTime": "2021-05-21T08:30:00Z"
  }
}
*/
func NewOBBRAccountsResponse(accounts []AccountData) ResponseAccountList {
	return ResponseAccountList{
		Data: accounts,
	}
}

func OBBRPermsToStringSlice(perms []acpClient.OpenbankingBrasilPermission) []string {
	var slice []string
	for _, perm := range perms {
		slice = append(slice, string(perm))
	}
	return slice
}

func NewOBBRPayment(introspectionResponse *acpClient.IntrospectOBBRPaymentConsentResponse, self strfmt.URI, id string) paymentModels.OpenbankingBrasilResponsePixPayment {
	// TODO: create
	return paymentModels.OpenbankingBrasilResponsePixPayment{
		Links: &paymentModels.OpenbankingBrasilLinks{
			Self: string(self),
		},
	}
}
