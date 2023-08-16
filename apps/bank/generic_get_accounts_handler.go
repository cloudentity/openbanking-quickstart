package main

import (

	// obbrAccountModels "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/models"
	"github.com/gin-gonic/gin"

	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type GenericGetAccountsHandler struct {
	*Server
	introspectionResponse *oauth2Models.IntrospectResponse
}

func NewGenericGetAccountsHandler(server *Server) GetEndpointLogic {
	return &GenericGetAccountsHandler{Server: server}
}

func (h *GenericGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.GenericIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token").WithError(err)
	}
	return nil
}

func (h *GenericGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = GenericMapError(c, err)
	return
}

func (h *GenericGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewGenericAccountsResponse(data.GenericAccounts), nil
}

func (h *GenericGetAccountsHandler) Validate(c *gin.Context) *Error {
	// scopes := strings.Split(h.introspectionResponse.Scope, " ")
	// if !has(scopes, "consents") {
	// 	return ErrForbidden.WithMessage("token has no consents scope granted")
	// }

	// grantedPermissions := h.introspectionResponse.Permissions
	// if !has(GenericPermsToStringSlice(grantedPermissions), "ACCOUNTS_READ") {
	// 	return ErrForbidden.WithMessage("ACCOUNTS_READ permission has not been granted")
	// }

	return nil
}

func (h *GenericGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *GenericGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var ret BankUserData
	for _, account := range data.GenericAccounts {
		// if has(h.introspectionResponse.AccountIDs, *account.AccountID) {
		ret.GenericAccounts = append(ret.GenericAccounts, account)
		// }
	}

	return ret
}
