package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

// swagger:route GET /balances bank getBalancesRequest
//
// get balances
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadBalance1
//   403: OBErrorResponse1
//   404: OBErrorResponse1
type OBBRGetBalancesHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOBBRDataAccessConsentResponse
}

func NewOBBRGetBalancesHandler(server *Server) GetEndpointLogic {
	return &OBBRGetBalancesHandler{Server: server}
}

func (h *OBBRGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBBRIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBBRGetBalancesHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	accountID := c.Param("accountID")
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts/%s/balances", strconv.Itoa(h.Config.Port), accountID))

	var balance OBBRBalance

	for _, b := range data.OBBRBalances {
		if b.AccountID == accountID {
			balance = b
			break
		}
	}

	return NewOBBRBalancesResponse(balance, self)
}

func (h *OBBRGetBalancesHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return ErrForbidden.WithMessage("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions

	var perms []string

	for _, p := range grantedPermissions {
		perms = append(perms, string(p))
	}

	if !has(perms, string(models.OpenbankingBrasilPermissionACCOUNTSBALANCESREAD)) {
		return ErrForbidden.WithMessage("ReadBalances permission has not been granted")
	}
	return nil
}

func (h *OBBRGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBBRGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	filteredBalances := []models.OBReadBalance1DataBalanceItems0{}

	for _, balance := range data.OBBRBalances {
		if has(h.introspectionResponse.AccountIDs, string(*balance.AccountID)) {
			filteredBalances = append(filteredBalances, balance)
		}
	}

	return BankUserData{
		OBBRBalances: filteredBalances,
	}
}
