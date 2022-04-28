package main

import (
	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	"github.com/gin-gonic/gin"
)

type CDRGetAccountsHandler struct {
	*Server
	introspectionResponse *cdr.CdrConsentIntrospectOKBody
}

func NewCDRGetAccountsHandler(server *Server) GetEndpointLogic {
	return &CDRGetAccountsHandler{Server: server}
}

func (h *CDRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	/*var err error
	if h.introspectionResponse, err = h.CDRIntrospectAccountsToken(c); err != nil {
		logrus.Errorf("introspect cdr token for accounts failed with %+v", err)
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	logrus.Infof("introspect succeeded")*/
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
	return nil
}

func (h *CDRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return "bfb689fb-7745-45b9-bbaa-b21e00072447"
}

func (h *CDRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
