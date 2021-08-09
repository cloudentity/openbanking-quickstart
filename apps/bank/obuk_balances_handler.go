package main

import (
	"fmt"
	"strconv"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:parameters getBalancesRequest
type GetBalancesRequest struct {
	RequestHeaders
}

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
type OBUKBalancesHandler struct {
	*Server
}

func (h *OBUKBalancesHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBUKBalancesHandler) SetRequest(c *gin.Context) error {
	return nil
}

func (h *OBUKBalancesHandler) MapError(err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKBalancesHandler) BuildResponse(data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.Balances, self)
}

func (h *OBUKBalancesHandler) Validate() error {

	/*
		grantedPermissions := introspectionResponse.Permissions

			scopes := strings.Split(introspectionResponse.Scope, " ")
			if !has(scopes, "accounts") {
				msg := "token has no accounts scope granted"
				c.JSON(http.StatusForbidden, models.OBErrorResponse1{
					Message: &msg,
				})
				return
			}

			if !has(grantedPermissions, "ReadBalances") {
				msg := "ReadBalances permission has not been granted"
				c.JSON(http.StatusForbidden, models.OBErrorResponse1{
					Message: &msg,
				})
				return
			}
	*/
	return nil
}

func (h *OBUKBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBUKBalancesHandler) Filter(data BankUserData) BankUserData {
	/*
		filteredBalances := []models.OBReadBalance1DataBalanceItems0{}

				for _, balance := range userBalances {
					if has(introspectionResponse.AccountIDs, string(*balance.AccountID)) {
						filteredBalances = append(filteredBalances, balance)
					}
				}
	*/
	return data
}

func NewBalancesResponse(balances []models.OBReadBalance1DataBalanceItems0, self strfmt.URI) models.OBReadBalance1 {
	balancesPointers := make([]*models.OBReadBalance1DataBalanceItems0, len(balances))

	for i, b := range balances {
		balance := b
		balancesPointers[i] = &balance
	}

	return models.OBReadBalance1{
		Data: &models.OBReadBalance1Data{
			Balance: balancesPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(balances)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}
