package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/transactions"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	models.OBTransaction6
	BankID string `json:"BankId"`
}

func (o *OBUKClient) GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error) {
	var (
		resp             *transactions.GetTransactionsOK
		transactionsData = []Transaction{}
		err              error
	)

	if resp, err = o.Transactions.GetTransactions(transactions.NewGetTransactionsParamsWithContext(c).WithAuthorization(accessToken), nil); err != nil {
		return transactionsData, err
	}

	for _, a := range resp.Payload.Data.Transaction {
		transactionsData = append(transactionsData, Transaction{
			OBTransaction6: *a,
			BankID:         bank.BankID,
		})
	}

	return transactionsData, nil
}

// TODO: /accounts/v1/accounts/{accountId}/transactions needs to be implemented in bank application
func (o *OBBRClient) GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error) {
	return []Transaction{}, nil
}
