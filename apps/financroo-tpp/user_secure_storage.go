package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

type BankTokens []BankToken

func (b *BankTokens) GetAccessToken(bankID string) (string, error) {
	for _, x := range *b {
		if x.BankID == bankID {
			// check if expired
			return x.AccessToken, nil
		}
	}

	return "", fmt.Errorf("no access token found for bank: %s", bankID)

}

type BankToken struct {
	BankID      string `json:"bank_id"`
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type UserSecureStorage struct {
	sc *securecookie.SecureCookie
}

func NewUserSecureStorage(sc *securecookie.SecureCookie) UserSecureStorage {
	return UserSecureStorage{sc: sc}
}

func (s *UserSecureStorage) Store(c *gin.Context, tokens BankTokens) error {
	var (
		encodedData string
		err         error
	)
	if encodedData, err = s.sc.Encode("data", tokens); err != nil {
		return err
	}

	c.SetCookie("data", encodedData, 0, "/", "", false, true)

	return nil
}

func (s *UserSecureStorage) Read(c *gin.Context) (BankTokens, error) {
	var (
		encodedData string
		tokens      []BankToken
		err         error
	)

	if encodedData, err = c.Cookie("data"); err != nil {
		return BankTokens{}, nil
	}

	if err = s.sc.Decode("data", encodedData, &tokens); err != nil {
		return tokens, err
	}

	return tokens, nil
}
