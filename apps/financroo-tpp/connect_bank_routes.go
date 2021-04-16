package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
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

func (s *Server) WebConfig() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authorizationServerURL": s.Config.ACPURL,
			"clientId":               s.Config.Login.ClientID,
			"authorizationServerId":  s.Config.Login.ServerID,
			"tenantId":               s.Config.Login.TenantID,
		})
	}
}

type ConnectBankRequest struct {
	Permissions []string `json:"permissions" binding:"required"`
}

func (s *Server) ConnectBank() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID           = BankID(c.Param("bankId"))
			clients          Clients
			ok               bool
			registerResponse *openbanking.CreateAccountAccessConsentRequestCreated
			connectRequest   = ConnectBankRequest{}
			user             User
			err              error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&connectRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if clients, ok = s.Clients[bankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", bankID))
		}

		if registerResponse, err = clients.AcpAccountsClient.Openbanking.CreateAccountAccessConsentRequest(
			openbanking.NewCreateAccountAccessConsentRequestParams().
				WithTid(clients.AcpAccountsClient.TenantID).
				WithAid(clients.AcpAccountsClient.ServerID).
				WithRequest(&models.AccountAccessConsentRequest{
					Data: &models.AccountAccessConsentRequestData{
						Permissions:        connectRequest.Permissions,
						ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24 * 30)),
					},
				}),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register account access consent: %+v", err))
			return
		}

		s.CreateConsentResponse(c, bankID, registerResponse.Payload.Data.ConsentID, user, clients.AcpAccountsClient)
	}
}

func (s *Server) ConnectBankCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app        string
			appStorage = AppStorage{}
			code       = c.Query("code")
			ok         bool
			clients    Clients
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

		if clients, ok = s.Clients[appStorage.BankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", appStorage.BankID))
		}

		if token, err = clients.AcpAccountsClient.Exchange(code, c.Query("state"), appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if err = s.ConnectBankForUser(appStorage, token); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.UIURL + "?connected=yes")
	}
}

func (s *Server) ConnectedBanks() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			user           User
			err            error
			clients        Clients
			tokenResponse  *models.TokenResponse
			ok             bool
			connectedBanks = []string{}
			expiredBanks   = []string{}
			tokens         = []BankToken{}
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		for i, b := range user.Banks {
			if clients, ok = s.Clients[BankID(b.BankID)]; !ok {
				c.String(http.StatusInternalServerError, fmt.Sprintf("client not configured for bank: %s", b.BankID))
				return
			}

			if tokenResponse, err = clients.RenewAccountsToken(b); err != nil {
				logrus.Warnf("failed to renew token for bank: %s", b.BankID)
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
