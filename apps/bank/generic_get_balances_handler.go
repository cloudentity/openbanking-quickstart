package main

import (
	"github.com/gin-gonic/gin"

	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type GenericGetBalancesHandler struct {
	*Server
	introspectionResponse *oauth2Models.IntrospectResponse
}

func NewGenericGetBalancesHandler(server *Server) GetEndpointLogic {
	return &GenericGetBalancesHandler{Server: server}
}

func (h *GenericGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.GenericIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *GenericGetBalancesHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = GenericMapError(c, err)
	return
}

func (h *GenericGetBalancesHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewGenericBalancesResponse(data.GenericBalances), nil
}

func (h *GenericGetBalancesHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *GenericGetBalancesHandler) GetUserIdentifier(_ *gin.Context) string {
	return GetGenericUserIdentifierClaimFromIntrospectionResponse(h.Config, h.introspectionResponse)
}

func (h *GenericGetBalancesHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}
