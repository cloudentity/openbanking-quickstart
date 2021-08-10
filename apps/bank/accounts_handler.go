package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

type RequestHeaders struct {
	// in:header
	AuthDate string `json:"x-fapi-auth-date"`
	// in:header
	CustomerIPAddress string `json:"x-fapi-customer-ip-address"`
	// in:header
	Authorization string `json:"authorization"`
	// in:header
	InteractionID string `json:"x-fapi-interaction-id"`
	// in:header
	CustomerAgent string `json:"x-customer-user-agent"`
}

// swagger:parameters getAccountsRequest
type GetAccountsRequest struct {
	RequestHeaders
}

// swagger:route GET /accounts bank getAccountsRequest
//
// get accounts
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadAccount6
//	 400: OBErrorResponse1
//   403: OBErrorResponse1
//   404: OBErrorResponse1
//   500: OBErrorResponse1
func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			userAccounts          []models.OBAccount6
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

		if !has(grantedPermissions, "ReadAccountsBasic") {
			msg := "ReadAccountsBasic permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if userAccounts, err = s.Storage.GetAccounts(introspectionResponse.Sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		filteredAccounts := []models.OBAccount6{}

		for _, account := range userAccounts {
			if has(introspectionResponse.AccountIDs, string(*account.AccountID)) {
				if !has(grantedPermissions, "ReadAccountsDetail") {
					account.Account = []*models.OBAccount6AccountItems0{}
				}

				filteredAccounts = append(filteredAccounts, account)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port)))
		c.PureJSON(http.StatusOK, NewAccountsResponse(filteredAccounts, self))
	}
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID   models.AccountID `json:"id"`
	Name models.Nickname  `json:"name"`
}

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
func (s *Server) InternalGetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			accounts []models.OBAccount6
			sub      = c.Param("sub")
			err      error
		)

		if accounts, err = s.Storage.GetAccounts(sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port)))
		c.PureJSON(http.StatusOK, NewAccountsResponse(accounts, self))
	}
}

func NewAccountsResponse(accounts []models.OBAccount6, self strfmt.URI) models.OBReadAccount6 {
	accountsPointers := make([]*models.OBAccount6, len(accounts))

	for i, a := range accounts {
		account := a
		accountsPointers[i] = &account
	}

	return models.OBReadAccount6{
		Data: &models.OBReadAccount6Data{
			Account: accountsPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(accounts)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}
