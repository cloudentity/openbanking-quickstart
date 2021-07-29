package main

import (
	"time"

	"github.com/cloudentity/acp-client-go/models"
)

type ConsentTools struct{
	Trans *Trans
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
			"headTitle": c.Trans.OrDefault("uk.account.headTitle", "Account Access Consent"),
			"title": c.Trans.OrDefault("uk.account.title", "Account selection"),
			"selectAccounts": c.Trans.OrDefault("uk.account.selectAccounts", "Select accounts to share with"),
			"reviewData": c.Trans.OrDefault("uk.account.review_data", "Review the data you will be sharing"),
			"permissions": c.Trans.OrDefault("uk.account.permissions", "Account permissions"),
			"purpose": c.Trans.OrDefault("uk.account.purpose", "Purpose for sharing data:"),
			"purposeDetail": c.Trans.OrDefault("uk.account.purposeDetail", "To uncover insights that can improve your financial well being."),
			"expiration": c.Trans.WithDataOrDefault("uk.account.expiration", map[string]interface{}{
				"client_name": clientName,
				"expiration_date": expirationDate,
			}, "<strong>{{ .client_name }}</strong> will have access to your account information until <strong>{{ .expiration_date }}</strong>", AsHTML),
			"cancel": c.Trans.OrDefault("uk.account.cancel", "Cancel"),
			"agree": c.Trans.OrDefault("uk.account.agree", "I Agree"),
		},
		"login_request":   loginRequest,
		"accounts":        accounts.Accounts,
		"permissions":     c.GetPermissionsWithDescription(consent.AccountAccessConsent.Permissions),
		"client_name":     clientName,
		"expiration_date": expirationDate,
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
			"headTitle": c.Trans.OrDefault("uk.payment.headTitle", "Domestic Payment Consent"),
			"title": c.Trans.OrDefault("uk.payment.title", "Account Confirm payment"),
			"paymentInfo": c.Trans.OrDefault("uk.payment.paymentInfo", "Payment Information"),
			"payeeAccountName": c.Trans.OrDefault("uk.payment.payeeAccountName", "Payee Account Name"),
			"sortCode": c.Trans.OrDefault("uk.payment.sortCode", "Sort code"),
			"accountNumber": c.Trans.OrDefault("uk.payment.accountNumber", "Account number"),
			"paymentReference": c.Trans.OrDefault("uk.payment.paymentReference", "Payment reference"),
			"amount": c.Trans.OrDefault("uk.payment.amount", "AMOUNT"),
			"accountInfo": c.Trans.OrDefault("uk.payment.accountInfo", "Account Information"),
			"clickToProceed": c.Trans.WithDataOrDefault("uk.payment.clickToProceed",map[string]interface{}{
				"client_name": clientName,
			}, "Click confirm to proceed with payment. Once payment is completed we’ll sign you out of <strong>gobank.com</strong> and return you to <strong>{{ .client_name }}</strong>", AsHTML),
			"cancel": c.Trans.OrDefault("uk.payment.cancel", "Cancel"),
			"confirm": c.Trans.OrDefault("uk.payment.confirm", "Confirm"),
		},
		"login_request": loginRequest,
		"accounts":      c.GetAccountsWithBalance(accounts, balances, string(*consent.DomesticPaymentConsent.Initiation.DebtorAccount.Identification)),
		"client_name":   clientName,
		"consent":       consent.DomesticPaymentConsent,
	}
}
