package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	acpClient "github.com/cloudentity/acp-client-go/models"
)

// swagger:route GET /accounts/v1/accounts/{accountID}/balances bank br getBalancesRequest
//
// get balance
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: AccountBalancesData
//   403: OpenbankingBrasilResponseError
//   404: OpenbankingBrasilResponseError
type OBBRGetBalanceHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOBBRDataAccessConsentResponse
}

func NewOBBRGetBalanceHandler(server *Server) GetEndpointLogic {
	return &OBBRGetBalanceHandler{Server: server}
}

func (h *OBBRGetBalanceHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBBRIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBBRGetBalanceHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRGetBalanceHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	accountID := c.Param("accountID")
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts/%s/balances", strconv.Itoa(h.Config.Port), accountID))

	logrus.WithField("balances", data.OBBRBalances).Info("xxxxxxxxxxxxxx")

	for _, b := range data.OBBRBalances {
		if b.AccountID == accountID {
			return NewOBBRBalanceResponse(b, self), nil
		}
	}

	return nil, ErrNotFound.WithMessage(fmt.Sprintf("account %s not found", accountID))
}

func (h *OBBRGetBalanceHandler) Validate(c *gin.Context) *Error {
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
		return ErrForbidden.WithMessage("ACCOUNTS_BALANCES_READ permission has not been granted")
	}
	return nil
}

func (h *OBBRGetBalanceHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBBRGetBalanceHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var filteredBalances []OBBRBalance

	for _, balance := range data.OBBRBalances {
		if has(h.introspectionResponse.AccountIDs, balance.AccountID) {
			filteredBalances = append(filteredBalances, balance)
		}
	}

	return BankUserData{
		OBBRBalances: filteredBalances,
	}
}
