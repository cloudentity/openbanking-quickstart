package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/models"
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
func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			userBalances          []models.OBReadBalance1DataBalanceItems0
			err                   error
		)

		if introspectionResponse, err = s.IntrospectAccountsToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

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

		if userBalances, err = s.Storage.GetBalances(introspectionResponse.Sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		balances := []*models.OBReadBalance1DataBalanceItems0{}

		for _, balance := range userBalances {
			b := balance
			if has(introspectionResponse.AccountIDs, string(*b.AccountID)) {
				balances = append(balances, &b)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(s.Config.Port)))
		response := models.OBReadBalance1{
			Data: &models.OBReadBalance1Data{
				Balance: balances,
			},
			Meta: &models.Meta{
				TotalPages: int32(len(balances)),
			},
			Links: &models.Links{
				Self: &self,
			},
		}

		c.PureJSON(http.StatusOK, response)
	}
}

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
func (s *Server) InternalGetBalances() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			balances []models.OBReadBalance1DataBalanceItems0
			sub      = c.Param("sub")
			err      error
		)

		if balances, err = s.Storage.GetBalances(sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		var balancesPointers []*models.OBReadBalance1DataBalanceItems0

		for _, b := range balances {
			balance := b
			balancesPointers = append(balancesPointers, &balance)
		}

		c.PureJSON(http.StatusOK, models.OBReadBalance1{
			Data: &models.OBReadBalance1Data{
				Balance: balancesPointers,
			},
		})
	}
}
