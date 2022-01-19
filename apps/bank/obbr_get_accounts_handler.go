package main

import (
	"strings"

	obbrAccountModels "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/models"
	"github.com/gin-gonic/gin"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

// swagger:route GET /accounts bank br getAccountsRequest
//
// get accounts
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: ResponseAccountList
//	 400: OpenbankingBrasilResponseError
//   403: OpenbankingBrasilResponseError
//   404: OpenbankingBrasilResponseError
//   500: OpenbankingBrasilResponseError
type OBBRGetAccountsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOBBRDataAccessConsentResponse
}

func NewOBBRGetAccountsHandler(server *Server) GetEndpointLogic {
	return &OBBRGetAccountsHandler{Server: server}
}

func (h *OBBRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBBRIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBBRGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewOBBRAccountsResponse(data.OBBRAccounts), nil
}

func (h *OBBRGetAccountsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "consents") {
		return ErrForbidden.WithMessage("token has no consents scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(OBBRPermsToStringSlice(grantedPermissions), "ACCOUNTS_READ") {
		return ErrForbidden.WithMessage("ACCOUNTS_READ permission has not been granted")
	}

	return nil
}

func (h *OBBRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBBRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var (
		filteredAccounts     []obbrAccountModels.AccountData
		requestedAccountType = c.Query("accountType")
	)

	for _, account := range data.OBBRAccounts {
		if !has(h.introspectionResponse.AccountIDs, *account.AccountID) {
			continue
		}
		if requestedAccountType != "" && string(*account.Type) != requestedAccountType {
			continue
		}
		filteredAccounts = append(filteredAccounts, account)
	}

	return BankUserData{
		OBBRAccounts: filteredAccounts,
	}
}
