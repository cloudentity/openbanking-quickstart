package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
)

// swagger:route GET /balances bank uk getBalancesRequest
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
type OBUKGetBalancesHandler struct {
	*Server
	introspectionResponse *obukModels.OpenbankingAccountAccessConsentIntrospectOKBody
}

func NewOBUKGetBalancesHandler(server *Server) GetEndpointLogic {
	return &OBUKGetBalancesHandler{Server: server}
}

func (h *OBUKGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBUKIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBUKGetBalancesHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.OBUKBalances, self), nil
}

func (h *OBUKGetBalancesHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return ErrForbidden.WithMessage("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadBalances") {
		return ErrForbidden.WithMessage("ReadBalances permission has not been granted")
	}
	return nil
}

func (h *OBUKGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	filteredBalances := []models.OBReadBalance1DataBalanceItems0{}

	for _, balance := range data.OBUKBalances {
		if has(h.introspectionResponse.AccountIDs, string(*balance.AccountID)) {
			filteredBalances = append(filteredBalances, balance)
		}
	}

	return BankUserData{
		OBUKBalances: filteredBalances,
	}
}
