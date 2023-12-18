package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudentity/openbanking-quickstart/shared"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"

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
		c.HTML(http.StatusOK, "index.html", gin.H{
			"featureFlags": s.Config.FeatureFlags,
			"spec":         s.Config.Spec,
			"currency":     s.Config.Currency,
		})
	}
}

type ConnectBankRequest struct {
	Permissions []string `binding:"required" json:"permissions"`
}

func (s *Server) ConnectBank() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID         = BankID(c.Param("bankId"))
			user           User
			consentID      string
			consentClient  ConsentClient
			accountsClient acpclient.Client
			err            error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if consentClient, err = s.Clients.GetConsentClient(bankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent client: %+v", err))
			return
		}

		if consentClient.CreateConsentExplicitly() {
			if consentID, err = consentClient.CreateAccountConsent(c); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to register account access consent: %+v", err))
				return
			}
		}

		if accountsClient, err = s.Clients.GetAccountsClient(bankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get accounts client: %+v", err))
			return
		}

		s.CreateConsentResponse(c, bankID, user, consentClient, accountsClient, consentID)
	}
}

func (s *Server) ConnectBankCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app                      string
			appStorage               = AppStorage{}
			token                    acpclient.Token
			responseClaims           shared.ResponseData
			ok                       bool
			signatureVerificationKey jose.JSONWebKey
			accountsClient           acpclient.Client
			err                      error
		)

		if app, err = c.Cookie("app"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get app cookie: %+v", err))
			return
		}

		if err = s.SecureCookie.Decode("app", app, &appStorage); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode app storage: %+v", err))
			return
		}

		if signatureVerificationKey, ok = s.Clients.SignatureVerificationKey[appStorage.BankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get signature verification key for bank: %s", appStorage.BankID))
			return
		}

		if responseClaims, err = shared.HandleAuthResponseMode(c.Request, signatureVerificationKey); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode response jwt token %v", err))
			return
		}

		if responseClaims.Error != "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %v: %v", responseClaims.Error, responseClaims.ErrorDescription))
			return
		}

		if accountsClient, err = s.Clients.GetAccountsClient(appStorage.BankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get accounts client: %+v", err))
			return
		}

		if token, err = accountsClient.Exchange(responseClaims.Code, responseClaims.State, appStorage.CSRF); err != nil {
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
			tokenResponse  *oauth2.TokenResponse
			connectedBanks = []string{}
			expiredBanks   = []string{}
			tokens         = []BankToken{}
			err            error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		for i, b := range user.Banks {
			var client acpclient.Client

			if client, err = s.Clients.GetAccountsClient(BankID(b.BankID)); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get accounts client: %+v for bank: %s", err, b.BankID))
				return
			}

			if tokenResponse, err = RenewAccountsToken(c, b, client); err != nil {
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

		availableBanks := []AvailableBank{}

		for _, b := range s.Config.Banks {
			availableBanks = append(availableBanks, AvailableBank{
				ID:      string(b.ID),
				Name:    b.Name,
				IconURL: b.IconURL,
				LogoURL: b.LogoURL,
			})
		}

		c.JSON(200, gin.H{
			"available_banks": availableBanks,
			"connected_banks": connectedBanks,
			"expired_banks":   expiredBanks,
		})
	}
}

type AvailableBank struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
	LogoURL string `json:"logo_url"`
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
