package main

import (
	"fmt"
	"net/http"
	"strings"

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

func (s *Server) GetClientWithToken(bank ConnectedBank, tokens BankTokens) (BankClient, string, error) {
	var (
		clients Clients
		client  BankClient
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

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			client       BankClient
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

			if accountsData, err = client.GetAccounts(c, accessToken, b); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
				return
			}
		}

		c.JSON(200, gin.H{
			"accounts": accountsData,
		})
	}
}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			client       BankClient
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

			if balancesData, err = client.GetBalances(c, accessToken, b); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get balances: %+v", err))
				return
			}
		}

		c.JSON(200, gin.H{
			"balances": balancesData,
		})
	}
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			client           BankClient
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

			if transactionsData, err = client.GetTransactions(c, accessToken, b); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get transactions: %+v", err))
				return
			}
		}

		c.JSON(200, gin.H{
			"transactions": transactionsData,
		})
	}
}
