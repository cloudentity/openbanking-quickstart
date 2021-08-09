package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:parameters getInternalBalancesRequest
type GetInternalBalancesRequest struct {
	// in:path
	Sub string `json:"sub"`
}

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
type OBUKBalancesInternalHandler struct {
	*Server
}

func (h *OBUKBalancesInternalHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKBalancesInternalHandler) SetRequest(c *gin.Context) error {
	return nil
}

func (h *OBUKBalancesInternalHandler) MapError(err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKBalancesInternalHandler) BuildResponse(data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.Balances, self)
}

func (h *OBUKBalancesInternalHandler) Validate() error {
	return nil
}

func (h *OBUKBalancesInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBUKBalancesInternalHandler) Filter(data BankUserData) BankUserData {

	return data
}
