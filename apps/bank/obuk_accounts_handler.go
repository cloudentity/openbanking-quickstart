package main

import (
	"fmt"
	"strconv"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
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
type OBUKAccountsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
}

func (h *OBUKAccountsHandler) SetIntrospectionResponse(c *gin.Context) error {
	var (
		resp *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
		err  error
	)

	if resp, err = h.IntrospectAccountsToken(c); err != nil {
		return err
	}

	h.introspectionResponse = resp
	return nil
}

func (h *OBUKAccountsHandler) MapError(err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKAccountsHandler) SetRequest(c *gin.Context) error {
	return nil
}

func (h *OBUKAccountsHandler) BuildResponse(data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.Accounts, self)
}

func (h *OBUKAccountsHandler) Validate() error {

	/*
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
	*/
	return nil
}

func (h *OBUKAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKAccountsHandler) Filter(data BankUserData) BankUserData {
	return data
}

func OBUKMapError(err error) models.OBError1 {
	msg := err.Error()
	return models.OBError1{
		Message: &msg,
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
