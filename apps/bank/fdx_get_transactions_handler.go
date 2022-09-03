package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	fdx "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
)

type FDXGetTransactionsHandler struct {
	*Server
	introspectionResponse *fdx.FdxConsentIntrospectOKBody
}

func NewFDXGetTransactionsHandler(server *Server) GetEndpointLogic {
	return &FDXGetTransactionsHandler{Server: server}
}

func (h *FDXGetTransactionsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.FDXIntrospectAccountsToken(c); err != nil {
		logrus.Errorf("introspect cdr token for transactions failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *FDXGetTransactionsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *FDXGetTransactionsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewFDXTransactionsResponse(data.FDXTransactions), nil
}

func (h *FDXGetTransactionsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "TRANSACTIONS") {
		return ErrForbidden.WithMessage("token has no TRANSACTIONS scope granted")
	}
	return nil
}

func (h *FDXGetTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return GetFDXUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *FDXGetTransactionsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	// TODO filter transactions but there is only 1 account so passing for now
	return data
}
