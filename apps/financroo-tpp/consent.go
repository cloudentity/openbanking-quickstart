package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	obModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
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

const (
	// OpenbankingBrasilConsentPermissionACCOUNTSREAD captures enum value "ACCOUNTS_READ"
	OpenbankingBrasilConsentPermissionACCOUNTSREAD models.OpenbankingBrasilConsentPermission = "ACCOUNTS_READ"

	// OpenbankingBrasilConsentPermissionACCOUNTSBALANCESREAD captures enum value "ACCOUNTS_BALANCES_READ"
	OpenbankingBrasilConsentPermissionACCOUNTSBALANCESREAD models.OpenbankingBrasilConsentPermission = "ACCOUNTS_BALANCES_READ"

	// OpenbankingBrasilConsentPermissionACCOUNTSTRANSACTIONSREAD captures enum value "ACCOUNTS_TRANSACTIONS_READ"
	OpenbankingBrasilConsentPermissionACCOUNTSTRANSACTIONSREAD models.OpenbankingBrasilConsentPermission = "ACCOUNTS_TRANSACTIONS_READ"

	// OpenbankingBrasilConsentPermissionACCOUNTSOVERDRAFTLIMITSREAD captures enum value "ACCOUNTS_OVERDRAFT_LIMITS_READ"
	OpenbankingBrasilConsentPermissionACCOUNTSOVERDRAFTLIMITSREAD models.OpenbankingBrasilConsentPermission = "ACCOUNTS_OVERDRAFT_LIMITS_READ"

	// OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_READ"
	OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSREAD models.OpenbankingBrasilConsentPermission = "CREDIT_CARDS_ACCOUNTS_READ"

	// OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSBILLSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_BILLS_READ"
	OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSBILLSREAD models.OpenbankingBrasilConsentPermission = "CREDIT_CARDS_ACCOUNTS_BILLS_READ"

	// OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSBILLSTRANSACTIONSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_BILLS_TRANSACTIONS_READ"
	OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSBILLSTRANSACTIONSREAD models.OpenbankingBrasilConsentPermission = "CREDIT_CARDS_ACCOUNTS_BILLS_TRANSACTIONS_READ"

	// OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSLIMITSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_LIMITS_READ"
	OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSLIMITSREAD models.OpenbankingBrasilConsentPermission = "CREDIT_CARDS_ACCOUNTS_LIMITS_READ"

	// OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSTRANSACTIONSREAD captures enum value "CREDIT_CARDS_ACCOUNTS_TRANSACTIONS_READ"
	OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSTRANSACTIONSREAD models.OpenbankingBrasilConsentPermission = "CREDIT_CARDS_ACCOUNTS_TRANSACTIONS_READ"

	// OpenbankingBrasilConsentPermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD captures enum value "CUSTOMERS_PERSONAL_IDENTIFICATIONS_READ"
	OpenbankingBrasilConsentPermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD models.OpenbankingBrasilConsentPermission = "CUSTOMERS_PERSONAL_IDENTIFICATIONS_READ"

	// OpenbankingBrasilConsentPermissionCUSTOMERSPERSONALADITTIONALINFOREAD captures enum value "CUSTOMERS_PERSONAL_ADITTIONALINFO_READ"
	OpenbankingBrasilConsentPermissionCUSTOMERSPERSONALADITTIONALINFOREAD models.OpenbankingBrasilConsentPermission = "CUSTOMERS_PERSONAL_ADITTIONALINFO_READ"

	// OpenbankingBrasilConsentPermissionCUSTOMERSBUSINESSIDENTIFICATIONSREAD captures enum value "CUSTOMERS_BUSINESS_IDENTIFICATIONS_READ"
	OpenbankingBrasilConsentPermissionCUSTOMERSBUSINESSIDENTIFICATIONSREAD models.OpenbankingBrasilConsentPermission = "CUSTOMERS_BUSINESS_IDENTIFICATIONS_READ"

	// OpenbankingBrasilConsentPermissionCUSTOMERSBUSINESSADITTIONALINFOREAD captures enum value "CUSTOMERS_BUSINESS_ADITTIONALINFO_READ"
	OpenbankingBrasilConsentPermissionCUSTOMERSBUSINESSADITTIONALINFOREAD models.OpenbankingBrasilConsentPermission = "CUSTOMERS_BUSINESS_ADITTIONALINFO_READ"

	// OpenbankingBrasilConsentPermissionFINANCINGSREAD captures enum value "FINANCINGS_READ"
	OpenbankingBrasilConsentPermissionFINANCINGSREAD models.OpenbankingBrasilConsentPermission = "FINANCINGS_READ"

	// OpenbankingBrasilConsentPermissionFINANCINGSSCHEDULEDINSTALMENTSREAD captures enum value "FINANCINGS_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentPermissionFINANCINGSSCHEDULEDINSTALMENTSREAD models.OpenbankingBrasilConsentPermission = "FINANCINGS_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentPermissionFINANCINGSPAYMENTSREAD captures enum value "FINANCINGS_PAYMENTS_READ"
	OpenbankingBrasilConsentPermissionFINANCINGSPAYMENTSREAD models.OpenbankingBrasilConsentPermission = "FINANCINGS_PAYMENTS_READ"

	// OpenbankingBrasilConsentPermissionFINANCINGSWARRANTIESREAD captures enum value "FINANCINGS_WARRANTIES_READ"
	OpenbankingBrasilConsentPermissionFINANCINGSWARRANTIESREAD models.OpenbankingBrasilConsentPermission = "FINANCINGS_WARRANTIES_READ"

	// OpenbankingBrasilConsentPermissionINVOICEFINANCINGSREAD captures enum value "INVOICE_FINANCINGS_READ"
	OpenbankingBrasilConsentPermissionINVOICEFINANCINGSREAD models.OpenbankingBrasilConsentPermission = "INVOICE_FINANCINGS_READ"

	// OpenbankingBrasilConsentPermissionINVOICEFINANCINGSSCHEDULEDINSTALMENTSREAD captures enum value "INVOICE_FINANCINGS_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentPermissionINVOICEFINANCINGSSCHEDULEDINSTALMENTSREAD models.OpenbankingBrasilConsentPermission = "INVOICE_FINANCINGS_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentPermissionINVOICEFINANCINGSPAYMENTSREAD captures enum value "INVOICE_FINANCINGS_PAYMENTS_READ"
	OpenbankingBrasilConsentPermissionINVOICEFINANCINGSPAYMENTSREAD models.OpenbankingBrasilConsentPermission = "INVOICE_FINANCINGS_PAYMENTS_READ"

	// OpenbankingBrasilConsentPermissionINVOICEFINANCINGSWARRANTIESREAD captures enum value "INVOICE_FINANCINGS_WARRANTIES_READ"
	OpenbankingBrasilConsentPermissionINVOICEFINANCINGSWARRANTIESREAD models.OpenbankingBrasilConsentPermission = "INVOICE_FINANCINGS_WARRANTIES_READ"

	// OpenbankingBrasilConsentPermissionLOANSREAD captures enum value "LOANS_READ"
	OpenbankingBrasilConsentPermissionLOANSREAD models.OpenbankingBrasilConsentPermission = "LOANS_READ"

	// OpenbankingBrasilConsentPermissionLOANSSCHEDULEDINSTALMENTSREAD captures enum value "LOANS_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentPermissionLOANSSCHEDULEDINSTALMENTSREAD models.OpenbankingBrasilConsentPermission = "LOANS_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentPermissionLOANSPAYMENTSREAD captures enum value "LOANS_PAYMENTS_READ"
	OpenbankingBrasilConsentPermissionLOANSPAYMENTSREAD models.OpenbankingBrasilConsentPermission = "LOANS_PAYMENTS_READ"

	// OpenbankingBrasilConsentPermissionLOANSWARRANTIESREAD captures enum value "LOANS_WARRANTIES_READ"
	OpenbankingBrasilConsentPermissionLOANSWARRANTIESREAD models.OpenbankingBrasilConsentPermission = "LOANS_WARRANTIES_READ"

	// OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_READ"
	OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTREAD models.OpenbankingBrasilConsentPermission = "UNARRANGED_ACCOUNTS_OVERDRAFT_READ"

	// OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTSCHEDULEDINSTALMENTSREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_SCHEDULED_INSTALMENTS_READ"
	OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTSCHEDULEDINSTALMENTSREAD models.OpenbankingBrasilConsentPermission = "UNARRANGED_ACCOUNTS_OVERDRAFT_SCHEDULED_INSTALMENTS_READ"

	// OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTPAYMENTSREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_PAYMENTS_READ"
	OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTPAYMENTSREAD models.OpenbankingBrasilConsentPermission = "UNARRANGED_ACCOUNTS_OVERDRAFT_PAYMENTS_READ"

	// OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTWARRANTIESREAD captures enum value "UNARRANGED_ACCOUNTS_OVERDRAFT_WARRANTIES_READ"
	OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTWARRANTIESREAD models.OpenbankingBrasilConsentPermission = "UNARRANGED_ACCOUNTS_OVERDRAFT_WARRANTIES_READ"

	// OpenbankingBrasilConsentPermissionRESOURCESREAD captures enum value "RESOURCES_READ"
	OpenbankingBrasilConsentPermissionRESOURCESREAD models.OpenbankingBrasilConsentPermission = "RESOURCES_READ"
)

type PermissionGroup string

const (
	CadastroDadosCadastraisPF           PermissionGroup = "Cadastros Dados Cadastrais PF"
	CadastroInformacoesComplementaresPF PermissionGroup = "Cadastro Informações complementares PF"
	CadastroDadosCadastraisPJ           PermissionGroup = "Cadastro Dados Cadastrais PJ "
	CadastroInformacoesComplementaresPJ PermissionGroup = "Cadastro Informações complementares PJ"
	ContasSaldos                        PermissionGroup = "Contas Saldos"
	ContasLimites                       PermissionGroup = "Contas Limites"
	ContasExtratos                      PermissionGroup = "Contas Extratos"
	CartaoDeCreditoLimites              PermissionGroup = "Cartão de Crédito Limites"
	CartaoDeCreditoTransacoes           PermissionGroup = "Cartão de Crédito Transações"
	CartaoDeCreditoFaturas              PermissionGroup = "Cartão de Crédito Faturas"
	OperacoesDeCreditoDadosDoContrato   PermissionGroup = "Operações de Crédito Dados do Contrato"
)

type Permissions []models.OpenbankingBrasilConsentPermission

var PermissionGroupMap = map[PermissionGroup]Permissions{
	CadastroDadosCadastraisPF: {
		OpenbankingBrasilConsentPermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	CadastroInformacoesComplementaresPF: {
		OpenbankingBrasilConsentPermissionCUSTOMERSPERSONALADITTIONALINFOREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	CadastroDadosCadastraisPJ: {
		OpenbankingBrasilConsentPermissionCUSTOMERSBUSINESSIDENTIFICATIONSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	CadastroInformacoesComplementaresPJ: {
		OpenbankingBrasilConsentPermissionCUSTOMERSBUSINESSADITTIONALINFOREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	ContasSaldos: {
		OpenbankingBrasilConsentPermissionACCOUNTSREAD,
		OpenbankingBrasilConsentPermissionACCOUNTSBALANCESREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	ContasLimites: {
		OpenbankingBrasilConsentPermissionACCOUNTSREAD,
		OpenbankingBrasilConsentPermissionACCOUNTSOVERDRAFTLIMITSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	ContasExtratos: {
		OpenbankingBrasilConsentPermissionACCOUNTSREAD,
		OpenbankingBrasilConsentPermissionACCOUNTSTRANSACTIONSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	CartaoDeCreditoLimites: {
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSREAD,
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSLIMITSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	CartaoDeCreditoTransacoes: {
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSREAD,
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSTRANSACTIONSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	CartaoDeCreditoFaturas: {
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSREAD,
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSBILLSREAD,
		OpenbankingBrasilConsentPermissionCREDITCARDSACCOUNTSBILLSTRANSACTIONSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
	OperacoesDeCreditoDadosDoContrato: {
		OpenbankingBrasilConsentPermissionLOANSREAD,
		OpenbankingBrasilConsentPermissionLOANSWARRANTIESREAD,
		OpenbankingBrasilConsentPermissionLOANSSCHEDULEDINSTALMENTSREAD,
		OpenbankingBrasilConsentPermissionLOANSPAYMENTSREAD,
		OpenbankingBrasilConsentPermissionFINANCINGSREAD,
		OpenbankingBrasilConsentPermissionFINANCINGSWARRANTIESREAD,
		OpenbankingBrasilConsentPermissionFINANCINGSSCHEDULEDINSTALMENTSREAD,
		OpenbankingBrasilConsentPermissionFINANCINGSPAYMENTSREAD,
		OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTREAD,
		OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTWARRANTIESREAD,
		OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTSCHEDULEDINSTALMENTSREAD,
		OpenbankingBrasilConsentPermissionUNARRANGEDACCOUNTSOVERDRAFTPAYMENTSREAD,
		OpenbankingBrasilConsentPermissionINVOICEFINANCINGSREAD,
		OpenbankingBrasilConsentPermissionINVOICEFINANCINGSWARRANTIESREAD,
		OpenbankingBrasilConsentPermissionINVOICEFINANCINGSSCHEDULEDINSTALMENTSREAD,
		OpenbankingBrasilConsentPermissionINVOICEFINANCINGSPAYMENTSREAD,
		OpenbankingBrasilConsentPermissionRESOURCESREAD,
	},
}

func (o *OBBRConsentClient) CreateAccountConsent(c *gin.Context) (string, error) {
	var (
		registerResponse *openbanking.CreateDataAccessConsentCreated
		connectRequest   = ConnectBankRequest{}
		permissions      []models.OpenbankingBrasilConsentPermission
		uniquePerms      = map[models.OpenbankingBrasilConsentPermission]bool{}
		err              error
	)

	if err = c.BindJSON(&connectRequest); err != nil {
		return "", err
	}

	for _, p := range connectRequest.Permissions {
		if perms, ok := PermissionGroupMap[PermissionGroup(p)]; ok {
			for _, p1 := range perms {
				uniquePerms[p1] = true
			}
		}
	}

	for uniquePerm := range uniquePerms {
		permissions = append(permissions, uniquePerm)
	}

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
					Permissions: permissions,
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
