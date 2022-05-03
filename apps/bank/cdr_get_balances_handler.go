package main

import (
	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CDRGetBalancesHandler struct {
	*Server
	introspectionResponse *cdr.CdrConsentIntrospectOKBody
}

func NewCDRGetBalancesHandler(server *Server) GetEndpointLogic {
	return &CDRGetBalancesHandler{Server: server}
}

func (h *CDRGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.CDRIntrospectAccountsToken(c); err != nil {
		logrus.Errorf("introspect cdr token for balances failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *CDRGetBalancesHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRBalancesResponse(data.CDRBalances), nil
}

func (h *CDRGetBalancesHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *CDRGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	logrus.Infof("introspection response sub is %s", h.introspectionResponse.Sub)
	return h.introspectionResponse.Sub
}

func (h *CDRGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
