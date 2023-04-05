package main

import (
	"fmt"

	cdrModels "github.com/cloudentity/acp-client-go/clients/cdr/models"
)

type CDRConsentTools struct {
	Trans  *Trans
	Config Config
}

func (c *CDRConsentTools) GetInternalBankDataIdentifier(sub string, authCtx cdrModels.AuthenticationContext) string {
	if c.Config.BankIDClaim == "sub" {
		return sub
	}

	return fmt.Sprintf("%v", authCtx[c.Config.BankIDClaim])
}

func (c *CDRConsentTools) GetCDRAccountAccessConsentTemplateData(
	loginRequest LoginRequest,
	arrangement *cdrModels.GetCDRConsentResponse,
	accounts InternalAccounts,
) map[string]interface{} {
	var (
		expirationDate      string
		selectedAccountList []string
		selectedAccounts    = map[string]bool{}
		headTitle           = c.Trans.T("cdr.account.headTitle")
		title               = c.Trans.T("cdr.account.title")
	)

	if arrangement.PreviousCdrArrangement != nil {
		selectedAccountList = arrangement.PreviousCdrArrangement.AccountIds
		headTitle = c.Trans.T("cdr.account.amend.headTitle")
		title = c.Trans.T("cdr.account.amend.title")
	} else {
		for _, a := range accounts.Accounts {
			selectedAccountList = append(selectedAccountList, a.ID)
		}
	}

	for _, a := range selectedAccountList {
		selectedAccounts[a] = true
	}

	for i, a := range accounts.Accounts {
		if preselected, ok := selectedAccounts[a.ID]; ok {
			a.Preselected = preselected
		}

		accounts.Accounts[i] = a
	}

	clientName := c.GetClientName(nil)

	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":      headTitle,
			"title":          title,
			"selectAccounts": c.Trans.T("cdr.account.selectAccounts"),
			"reviewData":     c.Trans.T("cdr.account.review_data"),
			"permissions":    c.Trans.T("cdr.account.permissions"),
			"purpose":        c.Trans.T("cdr.account.purpose"),
			"purposeDetail":  c.Trans.T("cdr.account.purposeDetail"),
			"expiration": c.Trans.TD("cdr.account.expiration", map[string]interface{}{
				"client_name":     clientName,
				"expiration_date": expirationDate,
			}),
			"cancel": c.Trans.T("cdr.account.cancel"),
			"agree":  c.Trans.T("cdr.account.agree"),
		},
		"login_request":    loginRequest,
		"accounts":         accounts.Accounts,
		"selectedAccounts": selectedAccounts,
		// "permissions":     c.GetPermissionsWithDescription(consent.AccountAccessConsent.Permissions), // nolint
		"client_name":      clientName,
		"expiration_date":  expirationDate,
		"ctx":              arrangement.AuthenticationContext,
		"prev_arrangement": arrangement.PreviousCdrArrangement,
		"amend":            arrangement.PreviousCdrArrangement != nil,
	}
}

func (c *CDRConsentTools) GetClientName(client *cdrModels.ClientInfo) string {
	if client != nil {
		return client.ClientName
	}

	return "TPP"
}

func (c *CDRConsentTools) GrantScopes(requestedScopes []*cdrModels.RequestedScope) []string {
	grantScopes := make([]string, len(requestedScopes))

	for i, r := range requestedScopes {
		grantScopes[i] = r.RequestedName
	}

	return grantScopes
}
