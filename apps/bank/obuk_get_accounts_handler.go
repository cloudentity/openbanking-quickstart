package main

import (
	"fmt"
	"strconv"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /accounts bank getAccountsRequest
//
// get accounts
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadAccount6
//	 400: OBErrorResponse1
//   403: OBErrorResponse1
//   404: OBErrorResponse1
//   500: OBErrorResponse1
type OBUKGetAccountsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
}

func (h *OBUKGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) error {
	if resp, err := h.IntrospectAccountsToken(c); err != nil {
		return err
	} else {
		h.introspectionResponse = resp
		return nil
	}
}

func (h *OBUKGetAccountsHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.Accounts, self)
}

func (h *OBUKGetAccountsHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBUKGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
