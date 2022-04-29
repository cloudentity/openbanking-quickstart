package main

import (
	"fmt"
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/client/banking"
	cdrBankingModels "github.com/cloudentity/openbanking-quickstart/openbanking/cdr/banking/models"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/transactions"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
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
		resp     *banking.GetTransactionsOK
		accounts []Account
	)

	if accounts, err = o.GetAccounts(c, accessToken, bank); err != nil {
		return transactionsData, errors.Wrap(err, "failed to get account ids for transactions")
	}

	for _, account := range accounts {
		if resp, err = o.Banking.Banking.GetTransactions(
			banking.NewGetTransactionsParams().
				WithDefaults().
				WithAccountID(string(*account.AccountID)),
			runtime.ClientAuthInfoWriterFunc(func(request runtime.ClientRequest, registry strfmt.Registry) error {
				return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", accessToken))
			}),
		); err != nil {
			return transactionsData, err
		}

		for _, cdrTransaction := range resp.Payload.Data.Transactions {
			if transaction, err := cdrTransactionToInternalTransaction(cdrTransaction, bank); err != nil {
				logrus.Infof("failed to map cdr transaction to internal transaction: %+v", err)
			} else {
				transactionsData = append(transactionsData, transaction)
			}
		}
	}
	return transactionsData, nil
}

func cdrTransactionToInternalTransaction(transaction *cdrBankingModels.BankingTransaction, bank ConnectedBank) (Transaction, error) {
	var (
		parsedTime time.Time
		err        error
	)

	if parsedTime, err = time.Parse(time.RFC3339, transaction.ExecutionDateTime); err != nil {
		logrus.Infof("failed to parse time %v", err)
		return Transaction{}, err
	}

	bookingDateTime := models.BookingDateTime(parsedTime)

	return Transaction{
		OBTransaction6: models.OBTransaction6{
			AccountID:       (*models.AccountID)(transaction.AccountID),
			TransactionID:   models.TransactionID(transaction.TransactionID),
			BookingDateTime: &bookingDateTime,
			Amount: &models.OBActiveOrHistoricCurrencyAndAmount9{
				Amount: (*models.OBActiveCurrencyAndAmountSimpleType)(transaction.Amount),
			},
			BankTransactionCode: &models.OBBankTransactionCodeStructure1{
				Code: &transaction.MerchantCategoryCode,
			},
			TransactionInformation: models.TransactionInformation(*transaction.Description),
		},
		BankID: bank.BankID,
	}, nil
}
