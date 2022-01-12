package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
)

func (o *OBUKConsentClient) CreateAccountConsent(c *gin.Context) (string, error) {
	var (
		registerResponse *openbanking.CreateAccountAccessConsentRequestCreated
		connectRequest   = ConnectBankRequest{}
		err              error
	)

	if err = c.BindJSON(&connectRequest); err != nil {
		return "", err
	}

	if registerResponse, err = o.Accounts.Openbanking.CreateAccountAccessConsentRequest(
		openbanking.NewCreateAccountAccessConsentRequestParamsWithContext(c).
			WithTid(o.Accounts.TenantID).
			WithAid(o.Accounts.ServerID).
			WithRequest(&models.AccountAccessConsentRequest{
				Data: &models.OBReadConsent1Data{
					Permissions:        connectRequest.Permissions,
					ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24 * 30)),
				},
				Risk: map[string]interface{}{},
			}),
		nil,
	); err != nil {
		return "", err
	}

	return registerResponse.Payload.Data.ConsentID, nil
}

type CreatePaymentRequest struct {
	Amount               string `json:"amount" binding:"required"`
	AccountID            string `json:"account_id" binding:"required"`
	PayeeAccountName     string `json:"payee_account_name" binding:"required"`
	PayeeAccountNumber   string `json:"payee_account_number" binding:"required"`
	PayeeAccountSortCode string `json:"payee_account_sort_code" binding:"required"`
	PaymentReference     string `json:"payment_reference" binding:"required"`
	BankID               BankID `json:"bank_id" binding:"required"`
}

func (o *OBUKConsentClient) CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error) {
	var (
		registerResponse *openbanking.CreateDomesticPaymentConsentCreated
		jwsSig           string
		payload          []byte
		err              error
	)

	authorisationType := "Single"
	identification := models.Identification0(req.PayeeAccountNumber)
	schemaName := models.OBExternalAccountIdentification4Code("UK.OBIE.SortCodeAccountNumber")
	account := models.OBWriteDomesticConsent4DataInitiationCreditorAccount{
		Identification: &identification,
		Name:           req.PayeeAccountName,
		SchemeName:     &schemaName,
	}

	debtorIdentification := models.Identification0(req.AccountID)
	debtorAccount := models.OBWriteDomesticConsent4DataInitiationDebtorAccount{
		Identification: &debtorIdentification,
		Name:           "myAccount", // todo
		SchemeName:     &schemaName,
	}
	id := uuid.New().String()[:10]
	currency := models.ActiveOrHistoricCurrencyCode("GBP")
	amount := models.OBActiveCurrencyAndAmountSimpleType(req.Amount)

	consentRequest := models.DomesticPaymentConsentRequest{
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
					Reference:    req.PaymentReference,
					Unstructured: "Unstructured todo",
				},
			},
			ReadRefundAccount: "No",
		},
		Risk: &models.OBRisk1{},
	}

	if payload, err = json.Marshal(consentRequest); err != nil {
		return "", errors.Wrapf(err, "failed to register domestic payment consent unable to marshal payload")
	}

	if jwsSig, err = o.Sign(payload); err != nil {
		return "", errors.Wrapf(err, "failed to create jws signature for payment consent request")
	}

	if registerResponse, err = o.Payments.Openbanking.CreateDomesticPaymentConsent(
		openbanking.NewCreateDomesticPaymentConsentParamsWithContext(c).
			WithTid(o.Payments.TenantID).
			WithAid(o.Payments.ServerID).
			WithXJwsSignature(&jwsSig).
			WithRequest(&consentRequest),
		nil,
	); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent: %+v", err))
		return "", errors.Wrapf(err, "failed to register domestic payment consent")
	}

	return registerResponse.Payload.Data.ConsentID, nil
}

func (o *OBUKConsentClient) GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error) {
	var (
		consentResponse interface{}
		err             error
	)

	params := openbanking.NewGetDomesticPaymentConsentRequestParamsWithContext(c).
		WithTid(o.Payments.TenantID).
		WithAid(o.Payments.ServerID).
		WithConsentID(consentID)

	if consentResponse, err = o.Payments.Openbanking.GetDomesticPaymentConsentRequest(params, nil); err != nil {
		return consentResponse, err
	}

	return consentResponse, nil
}

func (o *OBBRConsentClient) CreateAccountConsent(c *gin.Context) (string, error) {
	var (
		registerResponse *openbanking.CreateDataAccessConsentCreated
		err              error
	)

	if registerResponse, err = o.Accounts.Openbanking.CreateDataAccessConsent(
		openbanking.NewCreateDataAccessConsentParamsWithContext(c).
			WithTid(o.Accounts.TenantID).
			WithAid(o.Accounts.ServerID).
			WithRequest(&models.OBBRCustomerDataAccessConsentRequest{
				Data: &models.OpenbankingBrasilConsentData{
					ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24)),
					LoggedUser: &models.OpenbankingBrasilConsentLoggedUser{
						Document: &models.OpenbankingBrasilConsentDocument{
							Identification: "11111111111",
							Rel:            "CPF",
						},
					},
					Permissions: []models.OpenbankingBrasilConsentPermission{
						"ACCOUNTS_READ",
						"RESOURCES_READ",
						"ACCOUNTS_OVERDRAFT_LIMITS_READ",
					},
				},
			}),
		nil,
	); err != nil {
		return "", err
	}

	return registerResponse.Payload.Data.ConsentID, nil
}

func (o *OBBRConsentClient) CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error) {
	return "", nil
}

func (o *OBBRConsentClient) GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error) {
	return nil, nil
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