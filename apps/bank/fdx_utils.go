package main

import (
	"fmt"
	"log"
	"time"

	fdxModels "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"

	fdx "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"

	"github.com/sirupsen/logrus"

	"github.com/go-openapi/strfmt"
)

func FDXMapError(err *Error) (int, fdxModels.Error1) {
	return err.Code, fdxModels.Error1{
		Code:    fmt.Sprint(err.Code),
		Message: err.Message,
	}
}

type DepositAccount struct {
	AccountID      string  `json:"accountId"`
	Nickname       string  `json:"nickname"`
	Status         string  `json:"status"`
	BalanceAsOf    string  `json:"balanceAsOf"`
	CurrentBalance float64 `json:"currentBalance"`
}

func NewFDXAccountsResponse(accounts fdxModels.Accountsentity, self strfmt.URI) fdxModels.Accountsentity {
	return accounts
}

func NewFDXBalancesResponse(balance []fdxModels.AccountWithDetailsentity) *fdxModels.DepositAccountentity2 {
	if len(balance) < 1 {
		return &fdxModels.DepositAccountentity2{}
	}
	return balance[0].DepositAccount
}

func GetFDXUserIdentifierClaimFromIntrospectionResponse(config Config, introspectResponse *fdx.FdxConsentIntrospectOKBody) string {
	if claim, ok := introspectResponse.Ext[config.UserIdentifierClaim].(string); ok {
		return claim
	}

	logrus.Info("No user identifier claim configured. Falling back to sub")
	return introspectResponse.Sub
}

func NewFDXTransactionsResponse(transactions fdxModels.Transactionsentity) interface{} {
	return transactions
}

func NewFDXPayment(introspectionResponse *fdx.FdxConsentIntrospectOKBody, self strfmt.URI, id string) fdxModels.Paymententity {
	log.Println("NewFDXPayment called")
	t := strfmt.Date(time.Now())
	return fdxModels.Paymententity{
		PaymentID: &id,
		DueDate:   &t,
	}
}
