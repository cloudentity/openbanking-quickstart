package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /internal/balances bank br getInternalBalancesRequest
//
// get all balances for user
//
// Security:
//
//	defaultcc: accounts
//
// Responses:
//
//	200: AccountBalancesData
//	404: OpenbankingBrasilResponseError
type OBBRGetBalancesInternalHandler struct {
	*Server
}

func NewOBBRGetBalancesInternalHandler(server *Server) GetEndpointLogic {
	return &OBBRGetBalancesInternalHandler{Server: server}
}

func (h *OBBRGetBalancesInternalHandler) SetIntrospectionResponse(_ *gin.Context) *Error {
	return nil
}

func (h *OBBRGetBalancesInternalHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRGetBalancesInternalHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/internal/balances", strconv.Itoa(h.Config.Port)))

	return NewOBBRBalancesResponse(data.OBBRBalances, self), nil
}

func (h *OBBRGetBalancesInternalHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *OBBRGetBalancesInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Query("id")
}

func (h *OBBRGetBalancesInternalHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}
