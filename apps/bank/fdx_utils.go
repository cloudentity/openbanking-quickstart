package main

import (
	"fmt"
	"log"

	fdx "github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
	fdxAccounts "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/client/account_information"
	fdxModels "github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
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

func NewFDXAccountsResponse(accounts []fdxModels.Accountentity, self strfmt.URI) fdxAccounts.SearchForAccountsOK {
	accountPointers := &fdxModels.Accountsentity{}

	// TODO accounts is empty and returns the number request but with empty data
	log.Printf("NewFDXAccountsResponse called with accounts %+v", accounts)


	// acct := DepositAccount{
	// 	AccountID: "10001",
	// 	Nickname: "My Checking Acc XXXX3223",
	// 	Status: "OPEN",
	// 	BalanceAsOf: "2017-11-05T13:15:30.751Z",
	// 	CurrentBalance: 13300.35,
	// }



	for _, account := range accounts {
		accountPointers.Accounts = append(accountPointers.Accounts, account)
	}

	return fdxAccounts.SearchForAccountsOK{
		Payload: accountPointers,
	}
}

func NewFDXBalancesResponse(balances fdxModels.AccountWithDetailsentity) interface{} {
	// TODO hard code for now
	resp := fdxModels.AccountWithDetailsentity{
		DepositAccount: &fdxModels.DepositAccountentity2{CurrentBalance: 512.00},
	}

	return resp
}

func GetFDXUserIdentifierClaimFromIntrospectionResponse(config Config, introspectResponse *fdx.FdxConsentIntrospectOKBody) string {
	if claim, ok := introspectResponse.Ext[config.UserIdentifierClaim].(string); ok {
		return claim
	}

	logrus.Info("No user identifier claim configured. Falling back to sub")
	return introspectResponse.Sub
}
