package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/transactions"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	models.OBTransaction6
	/*AccountID       string          `json:"AccountID"`
	Amount          string          `json:"Amount"`
	BookingDateTime strfmt.DateTime `json:"BookingDateTime"`
	TransactionCode string          `json:"TransactionCode"`*/
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
			/*AccountID:       string(*a.AccountID),
			Amount:          string(*a.Amount.Amount),
			BookingDateTime: strfmt.DateTime(*a.BookingDateTime),*/
			BankID: bank.BankID,
		})
	}

	return transactionsData, nil
}

func (o *OBBRClient) GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error) {
	return []Transaction{}, nil
}
