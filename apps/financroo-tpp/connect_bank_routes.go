package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
	oauth2 "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type AppStorage struct {
	CSRF     acpclient.CSRF
	Sub      string
	IntentID string
	BankID   BankID
}

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"featureFlags": s.Config.FeatureFlags})
	}
}

type ConnectBankRequest struct {
	Permissions []string `json:"permissions" binding:"required"`
}

func (s *Server) ConnectBank() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID          = BankID(c.Param("bankId"))
			consentID       string
			user            User
			loginURLBuilder LoginURLBuilder
			err             error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if consentID, err = s.Clients.ConsentClient.CreateAccountConsent(c); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register account access consent: %+v", err))
			return
		}

		switch s.Config.Spec {
		case "obuk":
			loginURLBuilder, err = NewOBUKLoginURLBuilder()
		case "obbr":
			loginURLBuilder, err = NewOBBRLoginURLBuilder(c, s.Clients.AcpAccountsClient)
		}
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create login url builder: %+v", err))
			return
		}

		s.CreateConsentResponse(c, bankID, consentID, user, s.Clients.AcpAccountsClient, loginURLBuilder)
	}
}

func (s *Server) ConnectBankCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app        string
			appStorage = AppStorage{}
			code       = c.Query("code")
			token      acpclient.Token
			err        error
		)

		if c.Query("error") != "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %s: %s", c.Query("error"), c.Query("error_description")))
			return
		}

		if app, err = c.Cookie("app"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get app cookie: %+v", err))
			return
		}

		if err = s.SecureCookie.Decode("app", app, &appStorage); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode app storage: %+v", err))
			return
		}

		if token, err = s.Clients.AcpAccountsClient.Exchange(code, c.Query("state"), appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if err = s.ConnectBankForUser(appStorage, token); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.UIURL+"?connected=yes")
	}
}

func (s *Server) ConnectedBanks() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			user           User
			err            error
			tokenResponse  *oauth2.TokenResponse
			connectedBanks = []string{}
			expiredBanks   = []string{}
			tokens         = []BankToken{}
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}
		for i, b := range user.Banks {
			if tokenResponse, err = s.Clients.RenewAccountsToken(c, b); err != nil {
				logrus.WithError(err).Warnf("failed to renew token for bank: %s, err: %+v", b.BankID, err)
				expiredBanks = append(expiredBanks, b.BankID)
				continue
			}

			connectedBanks = append(connectedBanks, b.BankID)

			tokens = append(tokens, BankToken{
				BankID:      b.BankID,
				AccessToken: tokenResponse.AccessToken,
				ExpiresAt:   time.Now().Add(time.Second * time.Duration(tokenResponse.ExpiresIn)).Unix(),
			})

			user.Banks[i].RefreshToken = tokenResponse.RefreshToken
		}

		if err = s.UserRepo.Set(user); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			return
		}

		if err = s.UserSecureStorage.Store(c, tokens); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error while storing user data: %+v", err))
			return
		}

		c.JSON(200, gin.H{
			"connected_banks": connectedBanks,
			"expired_banks":   expiredBanks,
		})
	}
}

func (s *Server) DisconnectBank() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID = c.Param("bankId")
			user   User
			err    error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to get user: %+v", err))
			return
		}

		cb := []ConnectedBank{}
		for _, b := range user.Banks {
			if b.BankID != bankID {
				cb = append(cb, b)
			}
		}
		user.Banks = cb

		if err = s.UserRepo.Set(user); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) ConnectBankForUser(appStorage AppStorage, token acpclient.Token) error {
	var (
		user User
		err  error
		cb   = ConnectedBank{
			BankID:       string(appStorage.BankID),
			IntentID:     appStorage.IntentID,
			RefreshToken: token.RefreshToken,
		}
		found = false
	)

	if user, err = s.UserRepo.Get(appStorage.Sub); err != nil {
		return errors.Wrapf(err, "failed to get user")
	}

	for i, b := range user.Banks {
		if b.BankID == string(appStorage.BankID) {
			user.Banks[i] = cb
			found = true
			break
		}
	}

	if !found {
		user.Banks = append(user.Banks, cb)
	}

	if err = s.UserRepo.Set(user); err != nil {
		return errors.Wrapf(err, "failed to update user")
	}

	return nil
}
