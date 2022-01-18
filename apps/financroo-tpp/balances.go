package main

import (
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/client/balances"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
)

type Balance struct {
	models.OBReadBalance1DataBalanceItems0
	BankID string `json:"BankId"`
}

func (o *OBUKClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	var (
		resp         *balances.GetBalancesOK
		balancesData = []Balance{}
		err          error
	)

	if resp, err = o.Balances.GetBalances(balances.NewGetBalancesParamsWithContext(c).WithAuthorization(accessToken), nil); err != nil {
		return balancesData, err
	}

	for _, a := range resp.Payload.Data.Balance {
		balancesData = append(balancesData, Balance{
			OBReadBalance1DataBalanceItems0: *a,
			BankID:                          bank.BankID,
		})
	}

	return balancesData, nil
}

func (o *OBBRClient) GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error) {
	return []Balance{}, nil
}
