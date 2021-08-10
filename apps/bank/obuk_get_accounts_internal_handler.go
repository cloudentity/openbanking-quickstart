package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /internal/accounts/{sub} bank getInternalAccountsRequest
//
// get all accounts for user
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadAccount6
//   404: OBErrorResponse1
type OBUKGetAccountsInternalHandler struct {
	*Server
}

func (h *OBUKGetAccountsInternalHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKGetAccountsInternalHandler) MapError(c *gin.Context, err error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetAccountsInternalHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.Accounts.OBUK, self)
}

func (h *OBUKGetAccountsInternalHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBUKGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBUKGetAccountsInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
