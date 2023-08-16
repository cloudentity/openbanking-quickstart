package main

import (
	"github.com/gin-gonic/gin"

	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type GenericGetTransactionsHandler struct {
	*Server
	introspectionResponse *oauth2Models.IntrospectResponse
}

func NewGenericGetTransactionsHandler(server *Server) GetEndpointLogic {
	return &GenericGetTransactionsHandler{Server: server}
}

func (h *GenericGetTransactionsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.GenericIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *GenericGetTransactionsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = GenericMapError(c, err)
	return
}

func (h *GenericGetTransactionsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewGenericTransactionsResponse(data.GenericTransactions), nil
}

func (h *GenericGetTransactionsHandler) Validate(c *gin.Context) *Error {
	// scopes := strings.Split(h.introspectionResponse.Scope, " ")
	// if !has(scopes, "bank:transactions:read") {
	// 	return ErrForbidden.WithMessage("token has no bank:transactions:read scope granted")
	// }
	return nil
}

func (h *GenericGetTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return GetGenericUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *GenericGetTransactionsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var (
		ret       BankUserData
		accountID = c.Param("accountId")
	)

	for _, transaction := range data.GenericTransactions {
		if *transaction.AccountID == accountID {
			// if has(h.introspectionResponse.AccountIDs, *transaction.AccountID) && *transaction.AccountID == accountID {
			ret.GenericTransactions = append(ret.GenericTransactions, transaction)
		}
	}
	return ret
}
