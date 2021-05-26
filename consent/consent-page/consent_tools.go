package main

import (
	"time"

	"github.com/cloudentity/acp-client-go/models"
)

type ConsentTools struct {
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

	return map[string]interface{}{
		"login_request": loginRequest,
		"accounts":      accounts.Accounts,
		"permissions":   c.GetPermissionsWithDescription(consent.AccountAccessConsent.Permissions),
		// "client_name":     c.GetClientName(consent.Client), // client is no longer returned by acp
		"client_name":     c.GetClientName(nil),
		"expiration_date": expirationDate,
	}
}

func (c *ConsentTools) GetDomesticPaymentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetDomesticPaymentConsentResponse,
	accounts InternalAccounts,
	balances BalanceData,
) map[string]interface{} {
	return map[string]interface{}{
		"login_request": loginRequest,
		"accounts":      c.GetAccountsWithBalance(accounts, balances, string(*consent.DomesticPaymentConsent.Initiation.DebtorAccount.Identification)),
		// "client_name":   c.GetClientName(consent.Client),
		"client_name": c.GetClientName(nil),
		"consent":     consent.DomesticPaymentConsent,
	}
}
