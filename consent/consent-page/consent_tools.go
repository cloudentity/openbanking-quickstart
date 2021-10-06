package main

import (
	"fmt"
	"time"

	"github.com/cloudentity/acp-client-go/models"
)

type ConsentTools struct {
	Trans  *Trans
	Config Config
}

func (c *ConsentTools) GetInternalBankDataIdentifier(sub string, authCtx models.AuthenticationContext) string {
	if c.Config.BankIDClaim == "sub" {
		return sub
	}

	return fmt.Sprintf("%v", authCtx[c.Config.BankIDClaim])
}

func (c *ConsentTools) GetPermissionsWithDescription(requestedPermissions []string) map[string][]Permission {
	permissions := map[string][]Permission{}

	for _, p := range requestedPermissions {
		for _, c := range DataClusterLanguage {
			if c.Permission == p {
				if permissions[c.Cluster] == nil {
					permissions[c.Cluster] = []Permission{}
				}
				permissions[c.Cluster] = append(permissions[c.Cluster], c)
			}
		}
	}

	return permissions
}

func (c *ConsentTools) GrantScopes(requestedScopes []*models.RequestedScope) []string {
	grantScopes := make([]string, len(requestedScopes))

	for i, r := range requestedScopes {
		grantScopes[i] = r.RequestedName
	}

	return grantScopes
}

func (c *ConsentTools) GetAccountsWithBalance(accounts InternalAccounts, balances BalanceData, accountID string) []InternalAccount {
	var filteredAccounts []InternalAccount

	for _, a := range accounts.Accounts {
		if a.ID == accountID {
			for _, b := range balances.Balance {
				if b.AccountID == a.ID {
					a.Balance = b
				}
			}
			filteredAccounts = append(filteredAccounts, a)
		}
	}

	return filteredAccounts
}

func (c *ConsentTools) GetClientName(client *models.ClientInfo) string {
	if client != nil {
		return client.ClientName
	}

	return "TPP"
}

func (c *ConsentTools) GetAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetAccountAccessConsentResponse,
	accounts InternalAccounts,
) map[string]interface{} {
	var expirationDate string

	edt := time.Time(consent.AccountAccessConsent.ExpirationDateTime)
	if !edt.IsZero() {
		expirationDate = edt.Format("02/01/2006")
	}

	clientName := c.GetClientName(nil)
	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":      c.Trans.T("uk.account.headTitle"),
			"title":          c.Trans.T("uk.account.title"),
			"selectAccounts": c.Trans.T("uk.account.selectAccounts"),
			"reviewData":     c.Trans.T("uk.account.review_data"),
			"permissions":    c.Trans.T("uk.account.permissions"),
			"purpose":        c.Trans.T("uk.account.purpose"),
			"purposeDetail":  c.Trans.T("uk.account.purposeDetail"),
			"expiration": c.Trans.TD("uk.account.expiration", map[string]interface{}{
				"client_name":     clientName,
				"expiration_date": expirationDate,
			}),
			"cancel": c.Trans.T("uk.account.cancel"),
			"agree":  c.Trans.T("uk.account.agree"),
		},
		"login_request":   loginRequest,
		"accounts":        accounts.Accounts,
		"permissions":     c.GetPermissionsWithDescription(consent.AccountAccessConsent.Permissions),
		"client_name":     clientName,
		"expiration_date": expirationDate,
		"ctx":             consent.AuthenticationContext,
	}
}

func (c *ConsentTools) GetDomesticPaymentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetDomesticPaymentConsentResponse,
	accounts InternalAccounts,
	balances BalanceData,
) map[string]interface{} {
	clientName := c.GetClientName(nil)
	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":        c.Trans.T("uk.payment.headTitle"),
			"title":            c.Trans.T("uk.payment.title"),
			"paymentInfo":      c.Trans.T("uk.payment.paymentInfo"),
			"payeeAccountName": c.Trans.T("uk.payment.payeeAccountName"),
			"sortCode":         c.Trans.T("uk.payment.sortCode"),
			"accountNumber":    c.Trans.T("uk.payment.accountNumber"),
			"paymentReference": c.Trans.T("uk.payment.paymentReference"),
			"amount":           c.Trans.T("uk.payment.amount"),
			"accountInfo":      c.Trans.T("uk.payment.accountInfo"),
			"clickToProceed": c.Trans.TD("uk.payment.clickToProceed", map[string]interface{}{
				"client_name": clientName,
			}),
			"cancel":  c.Trans.T("uk.payment.cancel"),
			"confirm": c.Trans.T("uk.payment.confirm"),
		},
		"login_request": loginRequest,
		"accounts":      c.GetAccountsWithBalance(accounts, balances, string(*consent.DomesticPaymentConsent.Initiation.DebtorAccount.Identification)),
		"client_name":   clientName,
		"consent":       OBUKPaymentConsentTemplateData(consent.DomesticPaymentConsent),
		"ctx":           consent.AuthenticationContext,
	}
}

func (c *ConsentTools) GetOBBRDataAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetOBBRCustomerDataAccessConsentResponse,
	accounts InternalAccounts,
) map[string]interface{} {
	var expirationDate string

	edt := time.Time(consent.CustomerDataAccessConsent.ExpirationDateTime)
	if !edt.IsZero() {
		expirationDate = edt.Format("02/01/2006")
	}

	clientName := c.GetClientName(nil)
	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":      c.Trans.T("br.account.headTitle"),
			"title":          c.Trans.T("br.account.title"),
			"selectAccounts": c.Trans.T("br.account.selectAccounts"),
			"reviewData":     c.Trans.T("br.account.review_data"),
			"permissions":    c.Trans.T("br.account.permissions"),
			"purpose":        c.Trans.T("br.account.purpose"),
			"purposeDetail":  c.Trans.T("br.account.purposeDetail"),
			"expiration": c.Trans.TD("br.account.expiration", map[string]interface{}{
				"client_name":     clientName,
				"expiration_date": expirationDate,
			}),
			"cancel": c.Trans.T("br.account.cancel"),
			"agree":  c.Trans.T("br.account.agree"),
		},
		"login_request": loginRequest,
		"accounts":      accounts.Accounts,
		// "permissions":     c.GetPermissionsWithDescription(consent.AccountAccessConsent.Permissions),
		"client_name":     clientName,
		"expiration_date": expirationDate,
		"ctx":             consent.AuthenticationContext,
	}
}

func (c *ConsentTools) GetOBBRPaymentConsentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetOBBRCustomerPaymentConsentResponse,
	accounts InternalAccounts,
	balances BalanceData,
) map[string]interface{} {
	clientName := c.GetClientName(nil)
	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":        c.Trans.T("br.payment.headTitle"),
			"title":            c.Trans.T("br.payment.title"),
			"paymentInfo":      c.Trans.T("br.payment.paymentInfo"),
			"payeeAccountName": c.Trans.T("br.payment.payeeAccountName"),
			"sortCode":         c.Trans.T("br.payment.sortCode"),
			"accountNumber":    c.Trans.T("br.payment.accountNumber"),
			"paymentReference": c.Trans.T("br.payment.paymentReference"),
			"amount":           c.Trans.T("br.payment.amount"),
			"accountInfo":      c.Trans.T("br.payment.accountInfo"),
			"clickToProceed": c.Trans.TD("br.payment.clickToProceed", map[string]interface{}{
				"client_name": clientName,
			}),
			"cancel":  c.Trans.T("br.payment.cancel"),
			"confirm": c.Trans.T("br.payment.confirm"),
		},
		"login_request": loginRequest,
		"accounts":      c.GetAccountsWithBalance(accounts, balances, consent.CustomerPaymentConsent.DebtorAccount.Number),
		"client_name":   clientName,
		"consent":       OBBRPaymentConsentTemplateData(consent.CustomerPaymentConsent),
		"ctx":           consent.AuthenticationContext,
	}
}

type PaymentConsentTemplateData struct {
	AccountName    string
	SortCode       string
	Identification string
	Reference      string
	Currency       string
	Amount         string
}

func OBUKPaymentConsentTemplateData(consent *models.DomesticPaymentConsent) PaymentConsentTemplateData {
	return PaymentConsentTemplateData{
		AccountName:    consent.Initiation.CreditorAccount.Name,
		Identification: string(*consent.Initiation.DebtorAccount.Identification),
		Reference:      consent.Initiation.RemittanceInformation.Reference,
		Currency:       string(*consent.Initiation.InstructedAmount.Currency),
		Amount:         string(*consent.Initiation.InstructedAmount.Amount),
	}
}

func OBBRPaymentConsentTemplateData(consent *models.OBBRCustomerPaymentConsent) PaymentConsentTemplateData {
	return PaymentConsentTemplateData{
		AccountName:    consent.Creditor.Name,
		Identification: consent.DebtorAccount.Number,
		Currency:       consent.Payment.Currency,
		Amount:         consent.Payment.Amount,
	}
}
