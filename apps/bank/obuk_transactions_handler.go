package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
)

// swagger:parameters getTransactionsRequest
type GetTransactionsRequest struct {
	RequestHeaders
}

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
type OBUKTransactionsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
}

func (h *OBUKTransactionsHandler) SetIntrospectionResponse(c *gin.Context) error {
	var (
		resp *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
		err  error
	)

	if resp, err = h.IntrospectAccountsToken(c); err != nil {
		return err
	}

	h.introspectionResponse = resp
	return nil
}

func (h *OBUKTransactionsHandler) MapError(err error) interface{} {
	return OBUKMapError(err)
}

func (h *OBUKTransactionsHandler) SetRequest(c *gin.Context) error {
	return nil
}

func (h *OBUKTransactionsHandler) BuildResponse(data BankUserData) interface{} {
	transactions := []*models.OBTransaction6{}

	for _, transaction := range data.Transactions {
		t := transaction
		transactions = append(transactions, &t)
	}
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/transactions", strconv.Itoa(h.Config.Port)))
	response := models.OBReadTransaction6{
		Data: &models.OBReadDataTransaction6{
			Transaction: transactions,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(transactions)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
	return response
}

func (h *OBUKTransactionsHandler) Validate() error {

	/*
		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadTransactionsBasic") {
			msg := "ReadTransactionsBasic permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}
	*/
	return nil
}

func (h *OBUKTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *OBUKTransactionsHandler) Filter(data BankUserData) BankUserData {
	/*
		transactions := []*models.OBTransaction6{}

			for _, transaction := range userTransactions {
				t := transaction
				if has(introspectionResponse.AccountIDs, string(*t.AccountID)) {
					if !has(grantedPermissions, "ReadTransactionsDetail") {
						t.TransactionInformation = ""
						t.Balance = &models.OBTransactionCashBalance{}
						t.MerchantDetails = &models.OBMerchantDetails1{}
						t.CreditorAgent = &models.OBBranchAndFinancialInstitutionIdentification61{}
						t.CreditorAccount = &models.OBCashAccount60{}
						t.DebtorAgent = &models.OBBranchAndFinancialInstitutionIdentification62{}
						t.DebtorAccount = &models.OBCashAccount61{}
					}

					transactions = append(transactions, &t)
				}
			}
	*/
	return data

}
