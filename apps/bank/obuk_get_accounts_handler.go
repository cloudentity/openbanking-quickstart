package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

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
type OBUKGetAccountsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
}

func NewOBUKGetAccountsHandler(server *Server) GetEndpointLogic {
	return &OBUKGetAccountsHandler{Server: server}
}

func (h *OBUKGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBUKIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBUKGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewAccountsResponse(data.Accounts.OBUK, self)
}

func (h *OBUKGetAccountsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return ErrForbidden.WithMessage("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadAccountsBasic") {
		return ErrForbidden.WithMessage("ReadAccountsBasic permission has not been granted")
	}

	return nil
}

func (h *OBUKGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	grantedPermissions := h.introspectionResponse.Permissions
	filteredAccounts := []models.OBAccount6{}

	for _, account := range data.Accounts.OBUK {
		if has(h.introspectionResponse.AccountIDs, string(*account.AccountID)) {
			if !has(grantedPermissions, "ReadAccountsDetail") {
				account.Account = []*models.OBAccount6AccountItems0{}
			}

			filteredAccounts = append(filteredAccounts, account)
		}
	}
	return BankUserData{
		Accounts: Accounts{
			OBUK: filteredAccounts,
		},
	}
}
