package main

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v3"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/client/domestic_payments"
	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type CreateDomesticPaymentConsentRequest struct {
	Amount               string `json:"amount" binding:"required"`
	BankID               BankID `json:"bank_id" binding:"required"`
	AccountID            string `json:"account_id" binding:"required"`
	PayeeAccountName     string `json:"payee_account_name" binding:"required"`
	PayeeAccountNumber   string `json:"payee_account_number" binding:"required"`
	PayeeAccountSortCode string `json:"payee_account_sort_code" binding:"required"`
	PaymentReference     string `json:"payment_reference" binding:"required"`
}

func (s *Server) CreateDomesticPaymentConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			clients               Clients
			ok                    bool
			registerResponse      *openbanking.CreateDomesticPaymentConsentCreated
			paymentConsentRequest = CreateDomesticPaymentConsentRequest{}
			user                  User
			jwsSig                string
			payload               []byte
			err                   error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&paymentConsentRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if clients, ok = s.Clients[paymentConsentRequest.BankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", paymentConsentRequest.BankID))
		}

		authorisationType := "Single"
		identification := models.Identification0(paymentConsentRequest.PayeeAccountNumber)
		schemaName := models.OBExternalAccountIdentification4Code("UK.OBIE.SortCodeAccountNumber")
		account := models.OBWriteDomesticConsent4DataInitiationCreditorAccount{
			Identification: &identification,
			Name:           paymentConsentRequest.PayeeAccountName,
			SchemeName:     &schemaName,
		}

		debtorIdentification := models.Identification0(paymentConsentRequest.AccountID)
		debtorAccount := models.OBWriteDomesticConsent4DataInitiationDebtorAccount{
			Identification: &debtorIdentification,
			Name:           "myAccount", // todo
			SchemeName:     &schemaName,
		}
		id := uuid.New().String()[:10]
		currency := models.ActiveOrHistoricCurrencyCode("GBP")
		amount := models.OBActiveCurrencyAndAmountSimpleType(paymentConsentRequest.Amount)

		req := models.DomesticPaymentConsentRequest{
			Data: &models.OBWriteDomesticConsent4Data{
				Authorisation: &models.OBWriteDomesticConsent4DataAuthorisation{
					AuthorisationType:  authorisationType,
					CompletionDateTime: strfmt.DateTime(time.Now().Add(time.Hour)),
				},
				Initiation: &models.OBWriteDomesticConsent4DataInitiation{
					CreditorAccount:        &account,
					DebtorAccount:          &debtorAccount,
					EndToEndIdentification: id,
					InstructedAmount: &models.OBWriteDomesticConsent4DataInitiationInstructedAmount{
						Amount:   &amount,
						Currency: &currency,
					},
					InstructionIdentification: id,
					RemittanceInformation: &models.OBWriteDomesticConsent4DataInitiationRemittanceInformation{
						Reference:    paymentConsentRequest.PaymentReference,
						Unstructured: "Unstructured todo", // TODO invoice info?
					},
				},
				ReadRefundAccount: "No",
			},
			Risk: &models.OBRisk1{},
		}

		if payload, err = json.Marshal(req); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent unable to marshal paylaod: %+v", err))
			return
		}
		log.Println("vvvvv")
		if jwsSig, err = s.SigningKey(payload); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent unable to get cert: %+v", err))
			return
		} else {
			log.Println("sigis", jwsSig)
		}

		if registerResponse, err = clients.AcpPaymentsClient.Openbanking.CreateDomesticPaymentConsent(
			openbanking.NewCreateDomesticPaymentConsentParamsWithContext(c).
				WithTid(clients.AcpPaymentsClient.TenantID).
				WithAid(clients.AcpPaymentsClient.ServerID).
				//WithXJwsSignature(&jwsSig).
				WithRequest(&req),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent: %+v", err))
			return
		}

		s.CreateConsentResponse(c, paymentConsentRequest.BankID, registerResponse.Payload.Data.ConsentID, user, clients.AcpPaymentsClient)
	}
}

func (s *Server) SigningKey(payload []byte) (string, error) {
	var (
		cert   *x509.Certificate
		jwsSig string
		err    error
	)

	if cert, err = ParseCertificate([]byte(s.Config.CertFile)); err != nil {
		return "", errors.New(fmt.Sprintf("failed to register domestic payment consent unable to get cert: %+v", err))
	}

	key := jose.JSONWebKey{
		Key:   cert.PublicKey,
		KeyID: cert.SerialNumber.String(),
		Use:   "sig",
	}

	if jwsSig, err = Sign(payload, key, GetOpenbankingUKSignerOptions("application/json", key), true); err != nil {
		return "", errors.New(fmt.Sprintf("failed to register domestic payment consent: %+v", err))
	}

	return jwsSig, nil
}

func (s *Server) DomesticPaymentCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app             string
			appStorage      = AppStorage{}
			code            = c.Query("code")
			state           = c.Query("state")
			consentResponse *openbanking.GetDomesticPaymentConsentRequestOK
			initiation      obModels.OBWriteDomestic2DataInitiation
			risk            obModels.OBRisk1
			paymentCreated  *domestic_payments.CreateDomesticPaymentsCreated
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

		bank := s.Clients[appStorage.BankID]
		acpClient := bank.AcpPaymentsClient
		bankClient := bank.BankClient

		params := openbanking.NewGetDomesticPaymentConsentRequestParamsWithContext(c).
			WithTid(acpClient.TenantID).
			WithAid(acpClient.ServerID).
			WithConsentID(appStorage.IntentID)

		if token, err = acpClient.Exchange(code, state, appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if consentResponse, err = acpClient.Openbanking.GetDomesticPaymentConsentRequest(params, nil); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get consent: %+v", err))
			return
		}

		if initiation, err = getInitiation(consentResponse); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to map consent data initiation: %+v", err))
			return
		}

		if risk, err = getRisk(consentResponse); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to map consent risk: %+v", err))
			return
		}

		if paymentCreated, err = bankClient.DomesticPayments.CreateDomesticPayments(domestic_payments.NewCreateDomesticPaymentsParamsWithContext(c).
			WithAuthorization(token.AccessToken).
			WithOBWriteDomestic2Param(&obModels.OBWriteDomestic2{
				Data: &obModels.OBWriteDomestic2Data{
					ConsentID:  &appStorage.IntentID,
					Initiation: &initiation,
				},
				Risk: &risk,
			}), nil); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create payment: %+v", err))
			return
		}

		instructedAmount := consentResponse.Payload.Data.Initiation.InstructedAmount
		amount := url.QueryEscape(string(*instructedAmount.Amount))
		currency := url.QueryEscape(string(*instructedAmount.Currency))

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.UIURL+fmt.Sprintf("/investments/contribute/%s/success?amount=%s&currency=%s",
			*paymentCreated.Payload.Data.DomesticPaymentID, amount, currency))
	}
}

func getInitiation(consentResponse *openbanking.GetDomesticPaymentConsentRequestOK) (pi obModels.OBWriteDomestic2DataInitiation, err error) {
	var initiationPayload []byte

	if initiationPayload, err = json.Marshal(consentResponse.Payload.Data.Initiation); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(initiationPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}

func getRisk(consentResponse *openbanking.GetDomesticPaymentConsentRequestOK) (pi obModels.OBRisk1, err error) {
	var riskPayload []byte

	if riskPayload, err = json.Marshal(consentResponse.Payload.Risk); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(riskPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}
