package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	fdx "github.com/cloudentity/acp-client-go/clients/fdx/client/f_d_x"
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
		logrus.Errorf("introspect fdx token for transactions failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *FDXGetTransactionsHandler) MapError(_ *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetTransactionsHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewFDXTransactionsResponse(data.FDXTransactions), nil
}

func (h *FDXGetTransactionsHandler) Validate(_ *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "fdx:transactions:read") {
		return ErrForbidden.WithMessage("token has no TRANSACTIONS scope granted")
	}
	return nil
}

func (h *FDXGetTransactionsHandler) GetUserIdentifier(_ *gin.Context) string {
	return GetFDXUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *FDXGetTransactionsHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	// TODO filter transactions but there is only 1 account so passing for now
	return data
}
