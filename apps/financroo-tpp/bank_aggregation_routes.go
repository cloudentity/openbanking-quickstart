package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func fakeUserinfo(token string) (body map[string]interface{}, err error) {
	var raw []byte

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return body, errors.New("invalid jwt token")
	}

	if raw, err = base64.RawURLEncoding.DecodeString(parts[1]); err != nil {
		return body, err
	}

	if err = json.Unmarshal(raw, &body); err != nil {
		return body, err
	}

	return body, nil
}

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

	if claims, err = fakeUserinfo(token); err != nil {
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
		client BankClient
		token  string
		err    error
	)

	if token, err = tokens.GetAccessToken(bank.BankID); err != nil {
		return client, token, err
	}

	if client, err = s.Clients.GetBankClient(BankID(bank.BankID)); err != nil {
		return client, token, err
	}

	return client, token, nil
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
				continue
			}

			var data []Account
			if data, err = client.GetAccounts(c, accessToken, b); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
				return
			}
			accountsData = append(accountsData, data...)
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
				continue
			}

			if err = s.UserRepo.Set(user); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			}

			var balances []Balance
			if balances, err = client.GetBalances(c, accessToken, b); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get balances: %+v", err))
				return
			}
			balancesData = append(balancesData, balances...)
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
				continue
			}

			var transactions []Transaction
			if transactions, err = client.GetTransactions(c, accessToken, b); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get transactions: %+v", err))
				return
			}

			transactionsData = append(transactionsData, transactions...)
		}

		c.JSON(200, gin.H{
			"transactions": transactionsData,
		})
	}
}
