package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudentity/openbanking-quickstart/shared"
	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"

	acpclient "github.com/cloudentity/acp-client-go"
)

func (s *Server) CreateDomesticPaymentConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			request        CreatePaymentRequest
			user           User
			consentID      string
			consentClient  ConsentClient
			paymentsClient acpclient.Client
			err            error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&request); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if consentClient, err = s.Clients.GetConsentClient(request.BankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent client: %+v for bank: %s", err, request.BankID))
			return
		}

		if consentClient.CreateConsentExplicitly() {
			if consentID, err = consentClient.CreatePaymentConsent(c, request); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to register payment consent: %+v", err))
				return
			}
		}

		if paymentsClient, err = s.Clients.GetPaymentsClient(request.BankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get payment client: %+v for bank: %s", err, request.BankID))
			return
		}

		s.CreateConsentResponse(c, request.BankID, user, consentClient, paymentsClient, consentID)
	}
}

func (s *Server) DomesticPaymentCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app                      string
			appStorage               = AppStorage{}
			responseClaims           shared.ResponseData
			consentResponse          interface{}
			paymentCreated           PaymentCreated
			token                    acpclient.Token
			ok                       bool
			signatureVerificationKey jose.JSONWebKey
			bankClient               BankClient
			consentClient            ConsentClient
			paymentsClient           acpclient.Client
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
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %s: %s", responseClaims.Error, responseClaims.ErrorDescription))
			return
		}

		if paymentsClient, err = s.Clients.GetPaymentsClient(appStorage.BankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get payment client: %+v for bank: %s", err, appStorage.BankID))
			return
		}

		if token, err = paymentsClient.Exchange(responseClaims.Code, responseClaims.State, appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if consentClient, err = s.Clients.GetConsentClient(appStorage.BankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent client: %+v for bank: %s", err, appStorage.BankID))
			return
		}

		if consentResponse, err = consentClient.GetPaymentConsent(c, appStorage.IntentID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent: %+v", err))
			return
		}

		if bankClient, err = s.Clients.GetBankClient(appStorage.BankID); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get bank client: %+v for bank: %s", err, appStorage.BankID))
			return
		}

		if paymentCreated, err = bankClient.CreatePayment(c, consentResponse, token.AccessToken); err != nil {
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
