package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/client/banking"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/transactions"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func (o *CDRClient) GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) (transactionsData []Transaction, err error) {
	var (
		resp *banking.GetTransactionsOK
	)

	// will need to loop
	if resp, err = o.Banking.Banking.GetTransactions(
		banking.NewGetTransactionsParams().
			WithDefaults().
			WithAccountID("1000001"),
	); err != nil {
		return transactionsData, err
	}

	for _, a := range resp.Payload.Data.Transactions {
		transactionsData = append(transactionsData, Transaction{
			OBTransaction6: models.OBTransaction6{
				AccountID:       (*models.AccountID)(a.AccountID),
				TransactionID:   models.TransactionID(a.TransactionID),
				BookingDateTime: &models.BookingDateTime{},
				Amount: &models.OBActiveOrHistoricCurrencyAndAmount9{
					Amount: (*models.OBActiveCurrencyAndAmountSimpleType)(a.Amount),
				},
				BankTransactionCode:    &models.OBBankTransactionCodeStructure1{},
				TransactionInformation: models.TransactionInformation(""),
			},
			BankID: bank.BankID,
		})
	}

	logrus.Infof("transactions data %v", transactionsData)

	return transactionsData, nil
}
