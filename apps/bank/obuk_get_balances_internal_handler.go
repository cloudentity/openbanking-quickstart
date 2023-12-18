package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /internal/balances bank uk getInternalBalancesRequest
//
// get all balances for user
//
// Security:
//
//	defaultcc: accounts
//
// Responses:
//
//	200: OBReadBalance1
//	404: OBErrorResponse1
type OBUKGetBalancesInternalHandler struct {
	*Server
}

func NewOBUKGetBalancesInternalHandler(server *Server) GetEndpointLogic {
	return &OBUKGetBalancesInternalHandler{Server: server}
}

func (h *OBUKGetBalancesInternalHandler) SetIntrospectionResponse(_ *gin.Context) *Error {
	return nil
}

func (h *OBUKGetBalancesInternalHandler) MapError(_ *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetBalancesInternalHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.OBUKBalances, self), nil
}

func (h *OBUKGetBalancesInternalHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *OBUKGetBalancesInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Query("id")
}

func (h *OBUKGetBalancesInternalHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}
