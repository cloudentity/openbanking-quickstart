package main

import (
	"fmt"
	"time"

	"github.com/cloudentity/acp-client-go/clients/fdx/models"
)

type FDXConsentTools struct {
	Trans  *Trans
	Config Config
}

func (c *FDXConsentTools) GetClientName(client *models.ClientInfo) string {
	if client != nil {
		return client.ClientName
	}

	return DefaultTPPName
}

func (c *FDXConsentTools) GetAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetFDXConsentResponse,
	accounts InternalAccounts,
) map[string]interface{} {
	var expirationDate string

	edt := time.Time(consent.FdxConsent.ExpirationTime)
	if !edt.IsZero() {
		expirationDate = edt.Format("02/01/2006")
	}

	clientName := c.GetClientName(nil)
	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":      c.Trans.T("fdx.account.headTitle"),
			"title":          c.Trans.T("fdx.account.title"),
			"selectAccounts": c.Trans.T("fdx.account.selectAccounts"),
			"reviewData":     c.Trans.T("fdx.account.review_data"),
			"permissions":    c.Trans.T("fdx.account.permissions"),
			"purpose":        c.Trans.T("fdx.account.purpose"),
			"purposeDetail":  c.Trans.T("fdx.account.purposeDetail"),
			"expiration": c.Trans.TD("fdx.account.expiration", map[string]interface{}{
				"client_name":     clientName,
				"expiration_date": expirationDate,
			}),
			"cancel": c.Trans.T("fdx.account.cancel"),
			"agree":  c.Trans.T("fdx.account.agree"),
		},
		"login_request": loginRequest,
		"accounts":      accounts.Accounts,
		// "permissions":     c.GetPermissionsWithDescription(consent.AccountAccessConsent.Permissions),
		"client_name":     clientName,
		"expiration_date": expirationDate,
		"ctx":             consent.AuthenticationContext,
	}
}

func (c *FDXConsentTools) GetPermissionsWithDescription(requestedPermissions []string) map[string][]Permission {
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

func (c *FDXConsentTools) GetInternalBankDataIdentifier(sub string, authCtx models.AuthenticationContext) string {
	if c.Config.BankIDClaim == SubClaim {
		return sub
	}

	return fmt.Sprintf("%v", authCtx[c.Config.BankIDClaim])
}

func (c *FDXConsentTools) GetFDXAccountAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *models.GetFDXConsentResponse,
	accounts InternalAccounts,
) map[string]interface{} {
	var (
		expirationDate string
		headTitle      = c.Trans.T("fdx.account.headTitle")
		title          = c.Trans.T("fdx.account.title")
	)

	//nolint
	for i, a := range accounts.Accounts {
		accounts.Accounts[i] = a
	}

	clientName := c.GetClientName(nil)

	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":      headTitle,
			"title":          title,
			"selectAccounts": c.Trans.T("fdx.account.selectAccounts"),
			"reviewData":     c.Trans.T("fdx.account.review_data"),
			"permissions":    c.Trans.T("fdx.account.permissions"),
			"purpose":        c.Trans.T("fdx.account.purpose"),
			"purposeDetail":  c.Trans.T("fdx.account.purposeDetail"),
			"expiration": c.Trans.TD("fdx.account.expiration", map[string]interface{}{
				"client_name":     clientName,
				"expiration_date": expirationDate,
			}),
			"cancel":   c.Trans.T("fdx.account.cancel"),
			"agree":    c.Trans.T("fdx.account.agree"),
			"continue": c.Trans.T("fdx.account.continue"),
		},
		"login_request":   loginRequest,
		"accounts":        accounts.Accounts,
		"resources":       consent.FdxConsent.Resources,
		"client_name":     clientName,
		"expiration_date": expirationDate,
		"ctx":             consent.AuthenticationContext,
	}
}

func (c *FDXConsentTools) GrantScopes(requestedScopes []*models.RequestedScope) []string {
	grantScopes := make([]string, len(requestedScopes))

	for i, r := range requestedScopes {
		grantScopes[i] = r.RequestedName
	}

	return grantScopes
}
