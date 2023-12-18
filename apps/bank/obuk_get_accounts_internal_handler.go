package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /internal/accounts/{sub} bank uk getInternalAccountsRequest
//
// get all accounts for user
//
// Security:
//
//	defaultcc: accounts
//
// Responses:
//
//	200: OBReadAccount6
//	404: OBErrorResponse1
type OBUKGetAccountsInternalHandler struct {
	*Server
}

func NewOBUKGetAccountsInternalHandler(server *Server) GetEndpointLogic {
	return &OBUKGetAccountsInternalHandler{Server: server}
}

func (h *OBUKGetAccountsInternalHandler) SetIntrospectionResponse(_ *gin.Context) *Error {
	return nil
}

func (h *OBUKGetAccountsInternalHandler) MapError(_ *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetAccountsInternalHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/internal/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.OBUKAccounts, self), nil
}

func (h *OBUKGetAccountsInternalHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *OBUKGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Query("id")
}

func (h *OBUKGetAccountsInternalHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}
