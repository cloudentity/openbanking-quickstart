package main

import (
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

func (h *GenericGetAccountsHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewGenericAccountsResponse(data.GenericAccounts), nil
}

func (h *GenericGetAccountsHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *GenericGetAccountsHandler) GetUserIdentifier(_ *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *GenericGetAccountsHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}
