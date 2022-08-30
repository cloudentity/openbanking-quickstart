package main

import (
	"fmt"

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

func NewFDXAccountsResponse(accounts []fdxModels.Accountentity, self strfmt.URI) fdxAccounts.SearchForAccountsOK {
	accountPointers := &fdxModels.Accountsentity{}

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
