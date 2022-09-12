package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudentity/openbanking-quickstart/utils"
	"github.com/gin-gonic/gin"

	acpclient "github.com/cloudentity/acp-client-go"
)

func (s *Server) CreateDomesticPaymentConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			request CreatePaymentRequest
			user    User
			err     error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&request); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		s.CreateConsentResponse(c, request.BankID, user, s.Clients.AcpPaymentsClient, s.LoginURLBuilder, true, request)
	}
}

func (s *Server) DomesticPaymentCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app             string
			appStorage      = AppStorage{}
			responseClaims  utils.ResponseData
			consentResponse interface{}
			paymentCreated  PaymentCreated
			token           acpclient.Token
			err             error
		)

		if responseClaims, err = utils.HandleAuthResponseMode(c.Request, s.SignatureVerificationKey); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode response jwt token %v", err))
			return
		}

		if responseClaims.Error != "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %s: %s", responseClaims.Error, responseClaims.ErrorDescription))
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

		if token, err = s.Clients.AcpPaymentsClient.Exchange(responseClaims.Code, responseClaims.State, appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if consentResponse, err = s.Clients.ConsentClient.GetPaymentConsent(c, appStorage.IntentID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent: %+v", err))
			return
		}

		if paymentCreated, err = s.Clients.BankClient.CreatePayment(c, consentResponse, token.AccessToken); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create payment: %+v", err))
			return
		}

		amount := url.QueryEscape(paymentCreated.Amount)
		currency := url.QueryEscape(paymentCreated.Currency)

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.UIURL+fmt.Sprintf("/investments/contribute/%s/success?amount=%s&currency=%s",
			paymentCreated.PaymentID, amount, currency))
	}
}
