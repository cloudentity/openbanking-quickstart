package main

import (
	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type OBUKConsentFetcher struct {
	*Server
}

func NewOBUKConsentFetcher(server *Server) OBUKConsentFetcher {
	return OBUKConsentFetcher{server}
}

func (o *OBUKConsentFetcher) FetchConsents(c *gin.Context, accountIDs []string) ([]ClientConsents, error) {
	var (
		response *obukModels.ListOBConsentsOK
		err      error
		types    []string
		cac      []ClientConsents
		ok       bool
	)

	if types, ok = c.GetQueryArray("types"); !ok {
		types = nil
	}

	if response, err = o.Client.Openbanking.Openbankinguk.ListOBConsents(
		obukModels.NewListOBConsentsParamsWithContext(c).
			WithWid(o.Config.OpenbankingUKWorkspaceID).
			WithConsentsRequest(&obModels.ConsentsRequest{
				Types:    types,
				Accounts: accountIDs,
			}),
		nil,
	); err != nil {
		return cac, err
	}

	return MapClientsToConsents(o.getClients(response), o.getConsents(response)), nil
}

func (o *OBUKConsentFetcher) getClients(resp *obukModels.ListOBConsentsOK) []Client {
	var clients Clients

	for _, consent := range resp.Payload.Consents {
		if consent.Client != nil {
			clients = append(clients, Client{
				ID:        consent.Client.ID,
				Name:      consent.Client.Name,
				LogoURI:   consent.Client.LogoURI,
				ClientURI: consent.Client.ClientURI,
			})
		}
	}

	return clients.Unique()
}

func (o *OBUKConsentFetcher) getConsents(resp *obukModels.ListOBConsentsOK) []Consent {
	var consents []Consent

	for _, consent := range resp.Payload.Consents {
		var (
			expiresAt   strfmt.DateTime
			updatedAt   strfmt.DateTime
			permissions []string
		)

		c := Consent{
			AccountIDs:  consent.AccountIds,
			ConsentID:   consent.ConsentID,
			TenantID:    consent.TenantID,
			ServerID:    consent.ServerID,
			ClientID:    consent.ClientID,
			Status:      consent.Status,
			Type:        string(consent.Type),
			CreatedAt:   consent.CreatedAt,
			ExpiresAt:   expiresAt,
			UpdatedAt:   updatedAt,
			Permissions: permissions,
		}

		switch consent.Type {
		case "account_access":
			c.ExpiresAt = consent.AccountAccessConsent.ExpirationDateTime
			c.UpdatedAt = strfmt.DateTime(*consent.AccountAccessConsent.StatusUpdateDateTime)
			c.Permissions = consent.AccountAccessConsent.Permissions
		case "domestic_payment":
			c.UpdatedAt = consent.DomesticPaymentConsent.StatusUpdateDateTime
			c.DebtorAccountIdentification = string(*consent.DomesticPaymentConsent.Initiation.DebtorAccount.Identification)
			c.DebtorAccountName = consent.DomesticPaymentConsent.Initiation.DebtorAccount.Name
			c.CreditorAccountIdentification = string(*consent.DomesticPaymentConsent.Initiation.CreditorAccount.Identification)
			c.CreditorAccountName = consent.DomesticPaymentConsent.Initiation.CreditorAccount.Name
			c.Currency = string(*consent.DomesticPaymentConsent.Initiation.InstructedAmount.Currency)
			c.Amount = string(*consent.DomesticPaymentConsent.Initiation.InstructedAmount.Amount)
			c.CompletionDateTime = consent.DomesticPaymentConsent.Authorisation.CompletionDateTime
		}

		consents = append(consents, c)
	}

	return consents
}

func (o *OBUKConsentFetcher) RevokeConsent(c *gin.Context, consentID string) (err error) {
	if _, err = o.Client.Openbanking.Openbankinguk.RevokeOpenbankingConsent(
		obukModels.NewRevokeOpenbankingConsentParamsWithContext(c).
			WithWid(o.Config.OpenbankingUKWorkspaceID).
			WithConsentID(consentID),
		nil,
	); err != nil {
		return err
	}
	return nil
}
