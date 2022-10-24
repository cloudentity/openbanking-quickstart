package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accounts/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
)

// swagger:route GET /transactions bank uk getTransactionsRequest
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
	introspectionResponse *obukModels.OpenbankingAccountAccessConsentIntrospectOKBody
}

func NewOBUKGetTransactionsHandler(server *Server) GetEndpointLogic {
	return &OBUKGetTransactionsHandler{Server: server}
}

func (h *OBUKGetTransactionsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error
	if h.introspectionResponse, err = h.OBUKIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}
	return nil
}

func (h *OBUKGetTransactionsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBUKMapError(err)
	return
}

func (h *OBUKGetTransactionsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/transactions", strconv.Itoa(h.Config.Port)))
	return NewTransactionsResponse(data.OBUKTransactions, self), nil
}

func (h *OBUKGetTransactionsHandler) Validate(c *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "accounts") {
		return ErrForbidden.WithMessage("token has no accounts scope granted")
	}

	grantedPermissions := h.introspectionResponse.Permissions
	if !has(grantedPermissions, "ReadTransactionsBasic") {
		return ErrForbidden.WithMessage("ReadTransactionsBasic permission has not been granted")
	}

	return nil
}

func (h *OBUKGetTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKGetTransactionsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	var (
		filteredTransactions []models.OBTransaction6
		grantedPermissions   = h.introspectionResponse.Permissions
	)

	for _, transaction := range data.OBUKTransactions {
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

	return BankUserData{OBUKTransactions: filteredTransactions}
}
