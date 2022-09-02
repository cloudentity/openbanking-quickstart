package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	fdxModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
)

// swagger:route GET /accounts bank fdx getAccountsRequest
//
// get accounts
//
// Security:
//
//	defaultcc: accounts
//
// TODO - add the correct responses
// Responses:
//
//	  200: OBReadAccount6
//		 400: OBErrorResponse1
//	  403: OBErrorResponse1
//	  404: OBErrorResponse1
//	  500: OBErrorResponse1
type FDXGetAccountsHandler struct {
	*Server
	introspectionResponse *fdxModels.FdxConsentIntrospectOKBody
}

func NewFDXGetAccountsHandler(server *Server) GetEndpointLogic {
	return &FDXGetAccountsHandler{Server: server}
}

func (h *FDXGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.FDXIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *FDXGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewFDXAccountsResponse(data.FDXAccounts, self), nil
}

func (h *FDXGetAccountsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "ACCOUNT_DETAILED") {
		return ErrForbidden.WithMessage("token has no ACCOUNT_DETAILED scope granted")
	}

	return nil
}

func (h *FDXGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *FDXGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {

	// TODO handle filter

	return data
}
