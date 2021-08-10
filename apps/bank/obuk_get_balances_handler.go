package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:route GET /balances bank getBalancesRequest
//
// get balances
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadBalance1
//   403: OBErrorResponse1
//   404: OBErrorResponse1
type OBUKGetBalancesHandler struct {
	*Server
}

func (h *OBUKGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKGetBalancesHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.Balances, self)
}

func (h *OBUKGetBalancesHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBUKGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBUKGetBalancesHandler) CreateResource(c *gin.Context) (interface{}, error) {
	return "", nil
}

func (h *OBUKGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
