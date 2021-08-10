package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
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
	if resp, err := h.IntrospectAccountsToken(c); err != nil {
		return err
	} else {
		h.introspectionResponse = resp
		return nil
	}
}

func (h *OBUKGetBalancesHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(h.Config.Port)))
	return NewBalancesResponse(data.Balances, self)
}

func (h *OBUKGetBalancesHandler) Validate(c *gin.Context) error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return errors.New("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadBalances") {
		return errors.New("ReadBalances permission has not been granted")
	}
	return nil
}

func (h *OBUKGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	filteredBalances := []models.OBReadBalance1DataBalanceItems0{}

	for _, balance := range data.Balances {
		if has(h.introspectionResponse.AccountIDs, string(*balance.AccountID)) {
			filteredBalances = append(filteredBalances, balance)
		}
	}

	return BankUserData{Balances: filteredBalances}
}
