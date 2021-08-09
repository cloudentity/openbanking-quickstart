package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:parameters getInternalAccountsRequest
type GetInternalAccountsRequest struct {
	// in:path
	Sub string `json:"sub"`
}

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
type OBUKInternalAccountsHandler struct {
	*Server
}

func (h *OBUKInternalAccountsHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKInternalAccountsHandler) SetRequest(c *gin.Context) error {
	return nil
}

func (h *OBUKInternalAccountsHandler) MapError(err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKInternalAccountsHandler) BuildResponse(data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.Accounts, self)
}

func (h *OBUKInternalAccountsHandler) Validate() error {
	return nil
}

func (h *OBUKInternalAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBUKInternalAccountsHandler) Filter(data BankUserData) BankUserData {
	return data
}
