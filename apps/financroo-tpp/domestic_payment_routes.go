package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/client/domestic_payments"
	obukPaymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	acpclient "github.com/cloudentity/acp-client-go"

	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
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
			registerResponse      *obukModels.CreateDomesticPaymentConsentCreated
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
		identification := obModels.Identification0(paymentConsentRequest.PayeeAccountNumber)
		schemaName := obModels.OBExternalAccountIdentification4Code("UK.OBIE.SortCodeAccountNumber")
		account := obModels.OBWriteDomesticConsent4DataInitiationCreditorAccount{
			Identification: &identification,
			Name:           paymentConsentRequest.PayeeAccountName,
			SchemeName:     &schemaName,
		}

		debtorIdentification := obModels.Identification0(paymentConsentRequest.AccountID)
		debtorAccount := obModels.OBWriteDomesticConsent4DataInitiationDebtorAccount{
			Identification: &debtorIdentification,
			Name:           "myAccount", // todo
			SchemeName:     &schemaName,
		}
		id := uuid.New().String()[:10]
		currency := obModels.ActiveOrHistoricCurrencyCode("GBP")
		amount := obModels.OBActiveCurrencyAndAmountSimpleType(paymentConsentRequest.Amount)

		req := obModels.DomesticPaymentConsentRequest{
			Data: &obModels.OBWriteDomesticConsent4Data{
				Authorisation: &obModels.OBWriteDomesticConsent4DataAuthorisation{
					AuthorisationType:  authorisationType,
					CompletionDateTime: strfmt.DateTime(time.Now().Add(time.Hour)),
				},
				Initiation: &obModels.OBWriteDomesticConsent4DataInitiation{
					CreditorAccount:        &account,
					DebtorAccount:          &debtorAccount,
					EndToEndIdentification: id,
					InstructedAmount: &obModels.OBWriteDomesticConsent4DataInitiationInstructedAmount{
						Amount:   &amount,
						Currency: &currency,
					},
					InstructionIdentification: id,
					RemittanceInformation: &obModels.OBWriteDomesticConsent4DataInitiationRemittanceInformation{
						Reference:    paymentConsentRequest.PaymentReference,
						Unstructured: "Unstructured todo", // TODO invoice info?
					},
				},
				ReadRefundAccount: "No",
			},
			Risk: &obModels.OBRisk1{},
		}

		if payload, err = json.Marshal(req); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent unable to marshal paylaod: %+v", err))
			return
		}

		if jwsSig, err = s.JWSSignature(payload); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent unable to marshal paylaod: %+v", err))
			return
		}

		if registerResponse, err = clients.AcpPaymentsClient.Openbanking.Openbankinguk.CreateDomesticPaymentConsent(
			obukModels.NewCreateDomesticPaymentConsentParamsWithContext(c).
				WithXJwsSignature(&jwsSig).
				WithRequest(&req),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent: %+v", err))
			return
		}

		s.CreateConsentResponse(c, paymentConsentRequest.BankID, registerResponse.Payload.Data.ConsentID, user, clients.AcpPaymentsClient)
	}
}

func (s *Server) DomesticPaymentCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app             string
			appStorage      = AppStorage{}
			code            = c.Query("code")
			state           = c.Query("state")
			consentResponse *obukModels.GetDomesticPaymentConsentRequestOK
			initiation      obukPaymentModels.OBWriteDomestic2DataInitiation
			risk            obukPaymentModels.OBRisk1
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

		params := obukModels.NewGetDomesticPaymentConsentRequestParamsWithContext(c).
			WithConsentID(appStorage.IntentID)

		if token, err = acpClient.Exchange(code, state, appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if consentResponse, err = acpClient.Openbanking.Openbankinguk.GetDomesticPaymentConsentRequest(params, nil); err != nil {
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
			WithOBWriteDomestic2Param(&obukPaymentModels.OBWriteDomestic2{
				Data: &obukPaymentModels.OBWriteDomestic2Data{
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

func getInitiation(consentResponse *obukModels.GetDomesticPaymentConsentRequestOK) (pi obukPaymentModels.OBWriteDomestic2DataInitiation, err error) {
	var initiationPayload []byte

	if initiationPayload, err = json.Marshal(consentResponse.Payload.Data.Initiation); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(initiationPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}

func getRisk(consentResponse *obukModels.GetDomesticPaymentConsentRequestOK) (pi obukPaymentModels.OBRisk1, err error) {
	var riskPayload []byte

	if riskPayload, err = json.Marshal(consentResponse.Payload.Risk); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(riskPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}
