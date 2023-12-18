package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	fdx "github.com/cloudentity/acp-client-go/clients/fdx/client/f_d_x"
)

type FDXGetBalancesHandler struct {
	*Server
	introspectionResponse *fdx.FdxConsentIntrospectOKBody
}

func NewFDXGetBalancesHandler(server *Server) GetEndpointLogic {
	return &FDXGetBalancesHandler{Server: server}
}

func (h *FDXGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.FDXIntrospectAccountsToken(c); err != nil {
		logrus.Errorf("introspect fdx token for balances failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *FDXGetBalancesHandler) MapError(_ *gin.Context, _ *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetBalancesHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewFDXBalancesResponse(data.FDXBalances), nil
}

func (h *FDXGetBalancesHandler) Validate(_ *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")

	if !has(scopes, "fdx:accountdetailed:read") {
		return ErrForbidden.WithMessage("token has no fdx:accountdetailed:read scope granted")
	}
	return nil
}

func (h *FDXGetBalancesHandler) GetUserIdentifier(_ *gin.Context) string {
	return GetFDXUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *FDXGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var (
		accountID = c.Param("accountId")
		res       BankUserData
	)

	for _, acct := range data.FDXBalances {
		if acct.DepositAccount.AccountID == accountID {
			res.FDXBalances = append(res.FDXBalances, acct)
		}
	}

	return res
}
