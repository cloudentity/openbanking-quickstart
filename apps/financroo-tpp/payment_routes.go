package main

import (
	"fmt"
	"net/http"
	"net/url"

	acpclient "github.com/cloudentity/acp-client-go"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateDomesticPaymentConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			clients   Clients
			ok        bool
			request   CreatePaymentRequest
			consentID string
			user      User
			err       error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&request); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if clients, ok = s.Clients[request.BankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", request.BankID))
		}

		if consentID, err = clients.ConsentClient.CreatePaymentConsent(c, request); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register payment consent: %+v", err))
			return
		}

		s.CreateConsentResponse(c, request.BankID, consentID, user, clients.AcpPaymentsClient)
	}
}

func (s *Server) DomesticPaymentCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app             string
			appStorage      = AppStorage{}
			code            = c.Query("code")
			state           = c.Query("state")
			consentResponse interface{}
			paymentCreated  PaymentCreated
			token           acpclient.Token
			err             error
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

		clients := s.Clients[appStorage.BankID]
		acpClient := clients.AcpPaymentsClient
		bankClient := clients.BankClient

		if token, err = acpClient.Exchange(code, state, appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if consentResponse, err = clients.ConsentClient.GetPaymentConsent(c, appStorage.IntentID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent: %+v", err))
			return
		}

		if paymentCreated, err = bankClient.CreatePayment(c, consentResponse, token.AccessToken); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create payment: %+v", err))
			return
		}

		amount := url.QueryEscape(string(paymentCreated.Amount))
		currency := url.QueryEscape(string(paymentCreated.Currency))

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.UIURL+fmt.Sprintf("/investments/contribute/%s/success?amount=%s&currency=%s",
			paymentCreated.PaymentID, amount, currency))
	}
}
