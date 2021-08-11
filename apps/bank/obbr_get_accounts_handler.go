package main

import (
	"strings"

	"github.com/gin-gonic/gin"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

type OBBRGetAccountsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOBBRDataAccessConsentResponse
}

func NewOBBRGetAccountsHandler(server *Server) GetEndpointLogic {
	return &OBBRGetAccountsHandler{Server: server}
}

func (h *OBBRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBBRIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBBRGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	return NewOBBRAccountsResponse(data.Accounts.OBBR)
}

func (h *OBBRGetAccountsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "consents") {
		return ErrForbidden.WithMessage("token has no consents scope granted")
	}

	// TODO: look at how permissions for accounts in obbr work
	/*grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadAccountsBasic") {
		return ErrForbidden.WithMessage("ReadAccountsBasic permission has not been granted")
	}*/

	return nil
}

func (h *OBBRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

// TODO: will need to add filtering here depending on how permissions work in obbr
func (h *OBBRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}
