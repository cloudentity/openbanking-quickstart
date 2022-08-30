package main

import (
	fdx "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.Errorf("introspect cdr token for balances failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *FDXGetBalancesHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewFDXBalancesResponse(data.FDXBalances), nil
}

func (h *FDXGetBalancesHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *FDXGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return GetFDXUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *FDXGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
