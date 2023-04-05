package main

import (
	"github.com/cloudentity/openbanking-quickstart/generated/cdr/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	cdr "github.com/cloudentity/acp-client-go/clients/cdr/client/c_d_r"
)

func CDRMapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = 400, nil
	return
}

func NewCDRAccountsResponse(accounts []models.BankingAccount) interface{} {
	resp := models.ResponseBankingAccountList{
		Data: &models.Data2{
			Accounts: []*models.BankingAccount{},
		},
	}
	for _, account := range accounts {
		acc := account
		resp.Data.Accounts = append(resp.Data.Accounts, &acc)
	}
	return resp
}

func NewCDRTransactionsResponse(transactions []models.BankingTransaction) interface{} {
	resp := models.ResponseBankingTransactionList{
		Data: &models.Data3{
			Transactions: []*models.BankingTransaction{},
		},
	}
	for _, transaction := range transactions {
		trans := transaction
		resp.Data.Transactions = append(resp.Data.Transactions, &trans)
	}
	return resp
}

func NewCDRBalancesResponse(balances []models.BankingBalance) interface{} {
	resp := models.ResponseBankingAccountsBalanceList{
		Data: &models.Data4{
			Balances: []*models.BankingBalance{},
		},
	}
	for _, balance := range balances {
		bal := balance
		resp.Data.Balances = append(resp.Data.Balances, &bal)
	}
	return resp
}

func GetCDRUserIdentifierClaimFromIntrospectionResponse(config Config, introspectResponse *cdr.CdrConsentIntrospectOKBody) string {
	if claim, ok := introspectResponse.Ext[config.UserIdentifierClaim].(string); ok {
		return claim
	}

	logrus.Info("No user identifier claim configured. Falling back to sub")
	return introspectResponse.Sub
}
