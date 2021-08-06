package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudentity/openbanking-quickstart/client/accounts"
	"github.com/cloudentity/openbanking-quickstart/client/balances"
	"github.com/cloudentity/openbanking-quickstart/client/transactions"
	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *Server) WithUser(c *gin.Context) (User, BankTokens, error) {
	var (
		user   User
		claims map[string]interface{}
		sub    string
		ok     bool
		tokens []BankToken
		err    error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if token == "" {
		return user, tokens, errors.New("no access token provided")
	}

	if claims, err = s.LoginClient.Userinfo(token); err != nil {
		return user, tokens, errors.Wrapf(err, "invalid token")
	}

	if sub, ok = claims["sub"].(string); !ok {
		return user, tokens, errors.New("sub claim is missing")
	}

	if user, err = s.UserRepo.Get(sub); err != nil {
		return user, tokens, errors.Wrapf(err, "failed to get user")
	}

	if tokens, err = s.UserSecureStorage.Read(c); err != nil {
		return user, tokens, errors.Wrapf(err, "failed to read user store")
	}

	return user, tokens, nil
}

func (s *Server) GetClientWithToken(bank ConnectedBank, tokens BankTokens) (OpenbankingClient, string, error) {
	var (
		clients Clients
		client  OpenbankingClient
		ok      bool
		token   string
		err     error
	)

	if clients, ok = s.Clients[BankID(bank.BankID)]; !ok {
		return client, token, fmt.Errorf("can't get client for a bank: %s", bank.BankID)
	}

	if token, err = tokens.GetAccessToken(bank.BankID); err != nil {
		return client, token, err
	}

	return clients.BankClient, token, nil
}

type Account struct {
	*models.OBAccount6
	BankID string `json:"BankId"`
}

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp         *accounts.GetAccountsOK
			client       OpenbankingClient
			accessToken  string
			accountsData = []Account{}
			user         User
			tokens       BankTokens
			err          error
		)

		if user, tokens, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		// todo parallel
		for _, b := range user.Banks {
			if client, accessToken, err = s.GetClientWithToken(b, tokens); err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				return
			}

			if resp, err = client.Accounts.GetAccounts(accounts.NewGetAccountsParams().WithAuthorization(accessToken), nil); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
				return
			}

			for _, a := range resp.Payload.Data.Account {
				accountsData = append(accountsData, Account{
					OBAccount6: a,
					BankID:     b.BankID,
				})
			}
		}

		c.JSON(200, gin.H{
			"accounts": accountsData,
		})
	}
}

type Balance struct {
	models.OBReadBalance1DataBalanceItems0
	BankID string `json:"BankId"`
}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp         *balances.GetBalancesOK
			client       OpenbankingClient
			accessToken  string
			balancesData = []Balance{}
			tokens       BankTokens
			user         User
			err          error
		)

		if user, tokens, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		// todo parallel
		for _, b := range user.Banks {
			if client, accessToken, err = s.GetClientWithToken(b, tokens); err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				return
			}
			if err = s.UserRepo.Set(user); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			}

			if resp, err = client.Balances.GetBalances(balances.NewGetBalancesParams().WithAuthorization(accessToken), nil); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get balances: %+v", err))
				return
			}

			for _, a := range resp.Payload.Data.Balance {
				balancesData = append(balancesData, Balance{
					OBReadBalance1DataBalanceItems0: *a,
					BankID:                          b.BankID,
				})
			}
		}

		c.JSON(200, gin.H{
			"balances": balancesData,
		})
	}
}

type Transaction struct {
	models.OBTransaction6
	BankID string `json:"BankId"`
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp             *transactions.GetTransactionsOK
			client           OpenbankingClient
			accessToken      string
			transactionsData = []Transaction{}
			tokens           BankTokens
			user             User
			err              error
		)

		if user, tokens, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		// todo parallel
		for _, b := range user.Banks {
			if client, accessToken, err = s.GetClientWithToken(b, tokens); err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				return
			}

			if resp, err = client.Transactions.GetTransactions(transactions.NewGetTransactionsParams().WithAuthorization(accessToken), nil); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get transactions: %+v", err))
				return
			}

			for _, a := range resp.Payload.Data.Transaction {
				transactionsData = append(transactionsData, Transaction{
					OBTransaction6: *a,
					BankID:         b.BankID,
				})
			}
		}

		c.JSON(200, gin.H{
			"transactions": transactionsData,
		})
	}
}
