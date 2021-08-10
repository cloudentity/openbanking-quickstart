package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
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
type OBUKGetBalancesHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
}

func (h *OBUKGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) error {
	var err error
	h.introspectionResponse, err = h.IntrospectAccountsToken(c)
	return err
}

func (h *OBUKGetBalancesHandler) MapError(c *gin.Context, err error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.Balances.OBUK, self)
}

func (h *OBUKGetBalancesHandler) Validate(c *gin.Context) error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return NewErrForbidden("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadBalances") {
		return NewErrForbidden("ReadBalances permission has not been granted")
	}
	return nil
}

func (h *OBUKGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	filteredBalances := []models.OBReadBalance1DataBalanceItems0{}

	for _, balance := range data.Balances.OBUK {
		if has(h.introspectionResponse.AccountIDs, string(*balance.AccountID)) {
			filteredBalances = append(filteredBalances, balance)
		}
	}

	return BankUserData{
		Balances: Balances{
			OBUK: filteredBalances,
		},
	}
}
