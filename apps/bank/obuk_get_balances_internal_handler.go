package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /internal/balances/{sub} bank getInternalBalancesRequest
//
// get all balances for user
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadBalance1
//   404: OBErrorResponse1
type OBUKGetBalancesInternalHandler struct {
	*Server
}

func NewOBUKGetBalancesInternalHandler(server *Server) GetEndpointLogic {
	return &OBUKGetBalancesInternalHandler{Server: server}
}

func (h *OBUKGetBalancesInternalHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKGetBalancesInternalHandler) MapError(c *gin.Context, err error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetBalancesInternalHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.Balances.OBUK, self)
}

func (h *OBUKGetBalancesInternalHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBUKGetBalancesInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBUKGetBalancesInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
