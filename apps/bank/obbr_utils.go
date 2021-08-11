package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
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
