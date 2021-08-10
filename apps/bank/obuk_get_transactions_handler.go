package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
)

// swagger:route GET /transactions bank getTransactionsRequest
//
// get transactions
//
// Security:
//   defaultcc: accounts
//
// Responses:
//   200: OBReadTransaction6
//   400: OBErrorResponse1
//   403: OBErrorResponse1
//   404: OBErrorResponse1
type OBUKGetTransactionsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
}

func (h *OBUKGetTransactionsHandler) SetIntrospectionResponse(c *gin.Context) error {
	var err error
	h.introspectionResponse, err = h.IntrospectAccountsToken(c)
	return err
}

func (h *OBUKGetTransactionsHandler) MapError(c *gin.Context, err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKGetTransactionsHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/transactions", strconv.Itoa(h.Config.Port)))
	return NewTransactionsResponse(data.Transactions, self)
}

func (h *OBUKGetTransactionsHandler) Validate(c *gin.Context) error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return errors.New("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadTransactionsBasic") {
		return errors.New("ReadTransactionsBasic permission has not been granted")
	}

	return nil
}

func (h *OBUKGetTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetTransactionsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	grantedPermissions := h.introspectionResponse.Permissions
	filteredTransactions := []models.OBTransaction6{}

	for _, transaction := range data.Transactions {
		if has(h.introspectionResponse.AccountIDs, string(*transaction.AccountID)) {
			if !has(grantedPermissions, "ReadTransactionsDetail") {
				transaction.TransactionInformation = ""
				transaction.Balance = &models.OBTransactionCashBalance{}
				transaction.MerchantDetails = &models.OBMerchantDetails1{}
				transaction.CreditorAgent = &models.OBBranchAndFinancialInstitutionIdentification61{}
				transaction.CreditorAccount = &models.OBCashAccount60{}
				transaction.DebtorAgent = &models.OBBranchAndFinancialInstitutionIdentification62{}
				transaction.DebtorAccount = &models.OBCashAccount61{}
			}

			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	return BankUserData{Transactions: filteredTransactions}
}
