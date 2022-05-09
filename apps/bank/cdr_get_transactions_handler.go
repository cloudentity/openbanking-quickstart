package main

import (
	"strings"

	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CDRGetTransactionsHandler struct {
	*Server
	introspectionResponse *cdr.CdrConsentIntrospectOKBody
}

func NewCDRGetTransactionsHandler(server *Server) GetEndpointLogic {
	return &CDRGetTransactionsHandler{Server: server}
}

func (h *CDRGetTransactionsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.CDRIntrospectAccountsToken(c); err != nil {
		logrus.Errorf("introspect cdr token for transactions failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *CDRGetTransactionsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetTransactionsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRTransactionsResponse(data.CDRTransactions), nil
}

func (h *CDRGetTransactionsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "bank:transactions:read") {
		return ErrForbidden.WithMessage("token has no bank:transactions:read scope granted")
	}
	return nil
}

func (h *CDRGetTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return GetCDRUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *CDRGetTransactionsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var (
		ret       BankUserData
		accountID = c.Param("accountId")
	)

	for _, transaction := range data.CDRTransactions {
		if has(h.introspectionResponse.AccountIDs, *transaction.AccountID) && *transaction.AccountID == accountID {
			ret.CDRTransactions = append(ret.CDRTransactions, transaction)
		}
	}
	return ret
}
