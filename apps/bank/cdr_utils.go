package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/models"

	"github.com/gin-gonic/gin"
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
