package main

import (
	"fmt"
	"time"

	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
)

type OBBRConsentTools struct {
	Trans  *Trans
	Config Config
}

func (c *OBBRConsentTools) GetClientName(client *obbrModels.ClientInfo) string {
	if client != nil {
		return client.ClientName
	}

	return DefaultTPPName
}

func (c *OBBRConsentTools) GetAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *obbrModels.GetOBBRCustomerDataAccessConsentResponse,
	accounts InternalAccounts,
) map[string]interface{} {
	var (
		expirationDate string
		permissions    []string
	)

	edt := time.Time(consent.CustomerDataAccessConsent.ExpirationDateTime)
	if !edt.IsZero() {
		expirationDate = edt.Format("02/01/2006")
	}

	for _, item := range consent.CustomerDataAccessConsent.Permissions {
		permissions = append(permissions, string(item))
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
		"permissions":     c.GetPermissionsWithDescription(permissions),
		"client_name":     clientName,
		"expiration_date": expirationDate,
		"ctx":             consent.AuthenticationContext,
	}
}

func (c *OBBRConsentTools) GetPermissionsWithDescription(requestedPermissions []string) map[string][]Permission {
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

func (c *OBBRConsentTools) GetInternalBankDataIdentifier(sub string, authCtx obbrModels.AuthenticationContext) string {
	if c.Config.BankIDClaim == SubClaim {
		return sub
	}

	return fmt.Sprintf("%v", authCtx[c.Config.BankIDClaim])
}

func (c *OBBRConsentTools) GetOBBRDataAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *obbrModels.GetOBBRCustomerDataAccessConsentResponse,
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

func (c *OBBRConsentTools) GrantScopes(requestedScopes []*obbrModels.RequestedScope) []string {
	grantScopes := make([]string, len(requestedScopes))

	for i, r := range requestedScopes {
		grantScopes[i] = r.RequestedName
	}

	return grantScopes
}

func (c *OBBRConsentTools) GetAccountsWithBalance(accounts InternalAccounts, balances BalanceData, accountID string) []InternalAccount {
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

func (c *OBBRConsentTools) GetOBBRPaymentConsentTemplateData(
	loginRequest LoginRequest,
	wrapper OBBRConsentWrapper,
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
		"accounts":      c.GetAccountsWithBalance(accounts, balances, wrapper.GetDebtorAccountNumber()),
		"client_name":   clientName,
		"consent":       OBBRPaymentConsentTemplateDataFromWrapper(wrapper, c.Config.Currency),
		"ctx":           wrapper.GetAuthenticationContext(),
	}
}

func OBBRPaymentConsentTemplateDataFromWrapper(wrapper OBBRConsentWrapper, customCurrency Currency) PaymentConsentTemplateData {
	data := PaymentConsentTemplateData{
		AccountName:    wrapper.GetCreditorName(),
		Identification: wrapper.GetDebtorAccountNumber(),
		Currency:       wrapper.GetPaymentCurrency(),
		Amount:         wrapper.GetPaymentAmount(),
	}

	if customCurrency != "" {
		data.Currency = customCurrency.ToString()
	}

	return data
}
