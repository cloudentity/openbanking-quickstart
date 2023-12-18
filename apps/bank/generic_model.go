package main

import (
	"github.com/cloudentity/openbanking-quickstart/generated/cdr/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

func NewGenericTransactionsResponse(transactions []models.BankingTransaction) interface{} {
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

func NewGenericBalancesResponse(balances []models.BankingBalance) interface{} {
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

func GetGenericUserIdentifierClaimFromIntrospectionResponse(config Config, introspectResponse *oauth2Models.IntrospectResponse) string {
	if claim, ok := introspectResponse.Ext[config.UserIdentifierClaim].(string); ok {
		return claim
	}

	logrus.Info("No user identifier claim configured. Falling back to sub")
	return introspectResponse.Sub
}

func NewGenericAccountsResponse(accounts []models.BankingAccount) interface{} {
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

func GenericMapError(_ *gin.Context, _ *Error) (code int, resp interface{}) {
	code, resp = 400, nil
	return
}
