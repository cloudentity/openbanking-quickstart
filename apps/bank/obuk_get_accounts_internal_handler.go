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

func NewOBUKGetAccountsInternalHandler(server *Server) GetEndpointLogic {
	return &OBUKGetAccountsInternalHandler{Server: server}
}

func (h *OBUKGetAccountsInternalHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	return nil
}

func (h *OBUKGetAccountsInternalHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetAccountsInternalHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.OBUKAccounts, self)
}

func (h *OBUKGetAccountsInternalHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *OBUKGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Query("id")
}

func (h *OBUKGetAccountsInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
