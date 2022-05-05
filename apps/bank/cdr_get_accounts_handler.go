package main

import (
	"strings"

	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CDRGetAccountsHandler struct {
	*Server
	introspectionResponse *cdr.CdrConsentIntrospectOKBody
}

func NewCDRGetAccountsHandler(server *Server) GetEndpointLogic {
	return &CDRGetAccountsHandler{Server: server}
}

func (h *CDRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.CDRIntrospectAccountsToken(c); err != nil {
		logrus.Errorf("introspect cdr token for accounts failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *CDRGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRAccountsResponse(data.CDRAccounts), nil
}

func (h *CDRGetAccountsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "bank:accounts.basic:read") {
		return ErrForbidden.WithMessage("token has no bank:accounts.basic:read scope granted")
	}
	return nil
}

func (h *CDRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	logrus.Infof("introspection response sub is %s", h.introspectionResponse.Sub)
	return h.introspectionResponse.Sub
}

func (h *CDRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var ret BankUserData
	for _, account := range data.CDRAccounts {
		if has(h.introspectionResponse.AccountIDs, *account.AccountID) {
			ret.CDRAccounts = append(ret.CDRAccounts, account)
		}
	}
	return ret
}
