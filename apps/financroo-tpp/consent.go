package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	obbrModels "github.com/cloudentity/openbanking-quickstart/generated/obbr/consents/models"
	obModels "github.com/cloudentity/openbanking-quickstart/generated/obuk/payments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	obbrModels2 "github.com/cloudentity/acp-client-go/clients/obbr/models"

	acpclient "github.com/cloudentity/acp-client-go"

	obbrClientModels "github.com/cloudentity/acp-client-go/clients/obbr/client/o_b_b_r"
	obukModels "github.com/cloudentity/acp-client-go/clients/obuk/client/o_b_u_k"
	obukModels2 "github.com/cloudentity/acp-client-go/clients/obuk/models"
	"github.com/cloudentity/acp-client-go/clients/openbankingBR/payments/client/pagamentos"
)

func (o *OBUKConsentClient) CreateConsentExplicitly() bool {
	return true
}

func (o *OBUKConsentClient) UsePAR() bool {
	return false
}

func (o *OBUKConsentClient) DoPAR(c *gin.Context) (string, acpclient.CSRF, error) {
	return "", acpclient.CSRF{}, nil
}

func (o *OBUKConsentClient) CreateAccountConsent(c *gin.Context) (string, error) {
	var (
		registerResponse *obukModels.CreateAccountAccessConsentRequestCreated
		connectRequest   = ConnectBankRequest{}
		err              error
	)

	if err = c.BindJSON(&connectRequest); err != nil {
		return "", err
	}

	if registerResponse, err = o.Accounts.Obuk.Obuk.CreateAccountAccessConsentRequest(
		obukModels.NewCreateAccountAccessConsentRequestParamsWithContext(c).
			WithRequest(&obukModels2.AccountAccessConsentRequest{
				Data: &obukModels2.OBReadConsent1Data{
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
	Amount               string `binding:"required" json:"amount"`
	AccountID            string `binding:"required" json:"account_id"`
	PayeeAccountName     string `binding:"required" json:"payee_account_name"`
	PayeeAccountNumber   string `binding:"required" json:"payee_account_number"`
	PayeeAccountSortCode string `binding:"required" json:"payee_account_sort_code"`
	PaymentReference     string `binding:"required" json:"payment_reference"`
	BankID               BankID `binding:"required" json:"bank_id"`
}

func (o *OBUKConsentClient) CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error) {
	var (
		registerResponse *obukModels.CreateDomesticPaymentConsentCreated
		jwsSig           string
		payload          []byte
		err              error
	)

	authorisationType := "Single"
	identification := obukModels2.Identification0(req.PayeeAccountNumber)
	schemaName := obukModels2.OBExternalAccountIdentification4Code("UK.OBIE.SortCodeAccountNumber")
	account := obukModels2.OBWriteDomesticConsent4DataInitiationCreditorAccount{
		Identification: &identification,
		Name:           req.PayeeAccountName,
		SchemeName:     &schemaName,
	}

	debtorIdentification := obukModels2.Identification0(req.AccountID)
	debtorAccount := obukModels2.OBWriteDomesticConsent4DataInitiationDebtorAccount{
		Identification: &debtorIdentification,
		Name:           "myAccount", // todo
		SchemeName:     &schemaName,
	}
	id := uuid.New().String()[:10]
	currency := obukModels2.ActiveOrHistoricCurrencyCode("GBP")
	amount := obukModels2.OBActiveCurrencyAndAmountSimpleType(formatAmountAsCurrency(req.Amount))

	consentRequest := obukModels2.DomesticPaymentConsentRequest{
		Data: &obukModels2.OBWriteDomesticConsent4Data{
			Authorisation: &obukModels2.OBWriteDomesticConsent4DataAuthorisation{
				AuthorisationType:  authorisationType,
				CompletionDateTime: strfmt.DateTime(time.Now().Add(time.Hour)),
			},
			Initiation: &obukModels2.OBWriteDomesticConsent4DataInitiation{
				CreditorAccount:        &account,
				DebtorAccount:          &debtorAccount,
				EndToEndIdentification: id,
				InstructedAmount: &obukModels2.OBWriteDomesticConsent4DataInitiationInstructedAmount{
					Amount:   &amount,
					Currency: &currency,
				},
				InstructionIdentification: id,
				RemittanceInformation: &obukModels2.OBWriteDomesticConsent4DataInitiationRemittanceInformation{
					Reference:    req.PaymentReference,
					Unstructured: "Unstructured todo",
				},
			},
			ReadRefundAccount: "No",
		},
		Risk: &obukModels2.OBRisk1{},
	}

	if payload, err = json.Marshal(consentRequest); err != nil {
		return "", errors.Wrapf(err, "failed to register domestic payment consent unable to marshal payload")
	}

	if jwsSig, err = o.Sign(payload); err != nil {
		return "", errors.Wrapf(err, "failed to create jws signature for payment consent request")
	}

	if registerResponse, err = o.Payments.Obuk.Obuk.CreateDomesticPaymentConsent(
		obukModels.NewCreateDomesticPaymentConsentParamsWithContext(c).
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

	params := obukModels.NewGetDomesticPaymentConsentRequestParamsWithContext(c).
		WithConsentID(consentID)

	if consentResponse, err = o.Payments.Obuk.Obuk.GetDomesticPaymentConsentRequest(params, nil); err != nil {
		return consentResponse, err
	}

	return consentResponse, nil
}

type PermissionGroup string

const (
	CadastroDadosCadastraisPF           PermissionGroup = "Cadastros Dados Cadastrais PF"
	CadastroInformacoesComplementaresPF PermissionGroup = "Cadastro Informações complementares PF"
	CadastroDadosCadastraisPJ           PermissionGroup = "Cadastro Dados Cadastrais PJ "
	CadastroInformacoesComplementaresPJ PermissionGroup = "Cadastro Informações complementares PJ"
	ContasSaldos                        PermissionGroup = "Contas Saldos"
	ContasLimites                       PermissionGroup = "Contas Limites"
	ContasExtratos                      PermissionGroup = "Contas Extratos"
	CartaoDeCreditoLimites              PermissionGroup = "Cartão de Crédito Limites"              //nolint
	CartaoDeCreditoTransacoes           PermissionGroup = "Cartão de Crédito Transações"           //nolint
	CartaoDeCreditoFaturas              PermissionGroup = "Cartão de Crédito Faturas"              //nolint
	OperacoesDeCreditoDadosDoContrato   PermissionGroup = "Operações de Crédito Dados do Contrato" //nolint
)

type Permissions []obbrModels.OpenbankingBrasilConsentV2Permission

var PermissionGroupMap = map[PermissionGroup]Permissions{
	CadastroDadosCadastraisPF: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALIDENTIFICATIONSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	CadastroInformacoesComplementaresPF: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCUSTOMERSPERSONALADITTIONALINFOREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	CadastroDadosCadastraisPJ: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCUSTOMERSBUSINESSIDENTIFICATIONSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	CadastroInformacoesComplementaresPJ: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCUSTOMERSBUSINESSADITTIONALINFOREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	ContasSaldos: {
		obbrModels.OpenbankingBrasilConsentV2PermissionACCOUNTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionACCOUNTSBALANCESREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	ContasLimites: {
		obbrModels.OpenbankingBrasilConsentV2PermissionACCOUNTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionACCOUNTSOVERDRAFTLIMITSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	ContasExtratos: {
		obbrModels.OpenbankingBrasilConsentV2PermissionACCOUNTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionACCOUNTSTRANSACTIONSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	CartaoDeCreditoLimites: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSLIMITSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	CartaoDeCreditoTransacoes: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSTRANSACTIONSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	CartaoDeCreditoFaturas: {
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSBILLSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionCREDITCARDSACCOUNTSBILLSTRANSACTIONSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
	OperacoesDeCreditoDadosDoContrato: {
		obbrModels.OpenbankingBrasilConsentV2PermissionLOANSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionLOANSWARRANTIESREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionLOANSSCHEDULEDINSTALMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionLOANSPAYMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionFINANCINGSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionFINANCINGSWARRANTIESREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionFINANCINGSSCHEDULEDINSTALMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionFINANCINGSPAYMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTWARRANTIESREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTSCHEDULEDINSTALMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionUNARRANGEDACCOUNTSOVERDRAFTPAYMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSWARRANTIESREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSSCHEDULEDINSTALMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionINVOICEFINANCINGSPAYMENTSREAD,
		obbrModels.OpenbankingBrasilConsentV2PermissionRESOURCESREAD,
	},
}

func (o *OBBRConsentClient) UsePAR() bool {
	return false
}

func (o *OBBRConsentClient) DoPAR(c *gin.Context) (string, acpclient.CSRF, error) {
	return "", acpclient.CSRF{}, nil
}

func (o *OBBRConsentClient) CreateConsentExplicitly() bool {
	return true
}

func (o *OBBRConsentClient) CreateAccountConsent(c *gin.Context) (string, error) {
	var (
		registerResponse *obbrClientModels.CreateDataAccessConsentCreated
		connectRequest   = ConnectBankRequest{}
		permissions      []obbrModels2.OpenbankingBrasilConsentPermission
		uniquePerms      = map[obbrModels.OpenbankingBrasilConsentV2Permission]bool{}
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
		permissions = append(permissions, obbrModels2.OpenbankingBrasilConsentPermission(uniquePerm))
	}

	if registerResponse, err = o.Accounts.Obbr.Obbr.CreateDataAccessConsent(
		obbrClientModels.NewCreateDataAccessConsentParamsWithContext(c).
			WithRequest(&obbrModels2.BrazilCustomerDataAccessConsentRequestV1{
				Data: &obbrModels2.OpenbankingBrasilConsentData{
					ExpirationDateTime: strfmt.DateTime(time.Now().Add(time.Hour * 24)),
					LoggedUser: &obbrModels2.OpenbankingBrasilConsentLoggedUser{
						Document: &obbrModels2.OpenbankingBrasilConsentDocument{
							Identification: "11111111111",
							Rel:            "CPF",
						},
					},
					Permissions: permissions,
					BusinessEntity: &obbrModels2.OpenbankingBrasilConsentBusinessEntity{
						Document: &obbrModels2.OpenbankingBrasilConsentDocument1{
							Identification: "11111111111111",
							Rel:            "CNPJ",
						},
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
	var (
		response *pagamentos.PaymentsPostConsentsCreated
		bytes    []byte
		jwt      string
		err      error
	)

	personType := obbrModels2.OpenbankingBrasilPaymentEnumPaymentPersonType("PESSOA_NATURAL")
	creditor := obbrModels2.OpenbankingBrasilPaymentIdentification{
		PersonType: &personType,
		CpfCnpj:    "11111111111111",
		Name:       "Marco Antonio de Brito",
	}

	localInstrument := obbrModels2.OpenbankingBrasilPaymentEnumLocalInstrument("DICT")
	accountType := obbrModels2.OpenbankingBrasilPaymentEnumAccountPaymentsType("CACC")

	payment := obbrModels2.OpenbankingBrasilPaymentPaymentConsent{
		Type:         "PIX",
		Date:         strfmt.Date(time.Now().Add(time.Hour * 24)),
		Currency:     "BRL",
		Amount:       formatAmountAsCurrency(req.Amount),
		IbgeTownCode: "1234567",
		Details: &obbrModels2.OpenbankingBrasilPaymentDetails{
			CreditorAccount: &obbrModels2.OpenbankingBrasilPaymentCreditorAccount{
				AccountType: &accountType,
				Ispb:        "12345678",
				Number:      req.AccountID,
			},
			LocalInstrument: &localInstrument,
		},
	}

	businessEntity := obbrModels2.OpenbankingBrasilPaymentBusinessEntity{
		Document: &obbrModels2.OpenbankingBrasilPaymentDocument{
			Identification: "11111111111111",
			Rel:            "CNPJ",
		},
	}
	loggedUser := obbrModels2.OpenbankingBrasilPaymentLoggedUser{
		Document: &obbrModels2.OpenbankingBrasilPaymentDocument1{
			Identification: "11111111111",
			Rel:            "CPF",
		},
	}

	debtorAccount := obbrModels2.OpenbankingBrasilPaymentDebtorAccount{
		Number:      req.AccountID,
		Ispb:        "11111111",
		AccountType: &accountType,
	}

	obbrPaymentConsentRequest := obbrModels2.BrazilCustomerCreatePaymentConsentRequest{
		Aud: fmt.Sprintf("%s/open-banking/payments/v1/consents", o.Payments.Config.IssuerURL.String()),
		Iat: time.Now().Unix(),
		Jti: uuid.New().String(),
		Iss: "3333-3333-3333-3333",
		Data: &obbrModels2.OpenbankingBrasilPaymentData{
			BusinessEntity: &businessEntity,
			LoggedUser:     &loggedUser,
			Creditor:       &creditor,
			Payment:        &payment,
			DebtorAccount:  &debtorAccount,
		},
	}

	if bytes, err = json.Marshal(obbrPaymentConsentRequest); err != nil {
		return "", errors.Wrapf(err, "failed to marshal payment consent request")
	}

	if jwt, err = o.Signer.Sign(bytes); err != nil {
		return "", errors.Wrapf(err, "failed to sign payment consent request")
	}

	if response, err = o.Payments.OpenbankingBrasil.Payments.Pagamentos.PaymentsPostConsents(
		pagamentos.NewPaymentsPostConsentsParamsWithContext(c).
			WithBody(jwt),
		nil,
	); err != nil {
		return "", err
	}

	if bytes, err = json.Marshal(response.Payload); err != nil {
		return "", err
	}
	var consent obbrModels2.BrazilCustomerPaymentConsentResponse
	if err = json.Unmarshal(bytes, &consent); err != nil {
		return "", err
	}

	return consent.Data.ConsentID, nil
}

func (o *OBBRConsentClient) GetPaymentConsent(c *gin.Context, consentID string) (response interface{}, err error) {
	return o.Payments.Obbr.Obbr.GetPaymentConsent(
		obbrClientModels.NewGetPaymentConsentParamsWithContext(c).
			WithConsentID(consentID),
		nil,
	)
}

func formatAmountAsCurrency(amount string) string {
	var (
		f   float64
		err error
	)

	if f, err = strconv.ParseFloat(amount, 32); err != nil {
		return amount
	}

	return fmt.Sprintf("%.2f", f)
}

func getInitiation(consentResponse *obukModels.GetDomesticPaymentConsentRequestOK) (pi obModels.OBWriteDomestic2DataInitiation, err error) {
	var initiationPayload []byte

	if initiationPayload, err = json.Marshal(consentResponse.Payload.Data.Initiation); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(initiationPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}

func getRisk(consentResponse *obukModels.GetDomesticPaymentConsentRequestOK) (pi obModels.OBRisk1, err error) {
	var riskPayload []byte

	if riskPayload, err = json.Marshal(consentResponse.Payload.Risk); err != nil {
		return pi, err
	}

	if err = json.Unmarshal(riskPayload, &pi); err != nil {
		return pi, err
	}

	return pi, nil
}
