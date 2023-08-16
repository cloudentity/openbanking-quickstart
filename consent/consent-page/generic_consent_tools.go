package main

import (
	"fmt"

	"github.com/cloudentity/acp-client-go/clients/system/models"
)

type GenericConsentTools struct {
	Trans  *Trans
	Config Config
}

func (c *GenericConsentTools) GetClientName(client *models.ClientInfo) string {
	if client != nil {
		return client.ClientName
	}

	return "TPP"
}

func (c *GenericConsentTools) GetAccessConsentTemplateData(
	loginRequest LoginRequest,
	consent *models.ScopeGrantSessionResponse,
	accounts InternalAccounts,
) map[string]interface{} {

	clientName := c.GetClientName(consent.ClientInfo)
	return map[string]interface{}{
		"trans": map[string]interface{}{
			"headTitle":      c.Trans.T("uk.account.headTitle"),
			"title":          c.Trans.T("uk.account.title"),
			"selectAccounts": c.Trans.T("uk.account.selectAccounts"),
			"reviewData":     c.Trans.T("uk.account.review_data"),
			"purpose":        c.Trans.T("uk.account.purpose"),
			"purposeDetail":  c.Trans.T("uk.account.purposeDetail"),
			"expiration": c.Trans.TD("uk.account.expiration", map[string]interface{}{
				"client_name": clientName,
				// "expiration_date": expirationDate,
			}),
			"cancel": c.Trans.T("uk.account.cancel"),
			"agree":  c.Trans.T("uk.account.agree"),
		},
		"login_request": loginRequest,
		"accounts":      accounts.Accounts,
		"client_name":   clientName,
		// "expiration_date": expirationDate,
		"ctx": consent.AuthenticationContext,
	}
}

func (c *GenericConsentTools) GetInternalBankDataIdentifier(sub string, authCtx models.AuthenticationContext) string {
	if c.Config.BankIDClaim == "sub" {
		return sub
	}

	return fmt.Sprintf("%v", authCtx[c.Config.BankIDClaim])
}
