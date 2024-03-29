package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	obukModels "github.com/cloudentity/acp-client-go/clients/obuk/client/m_a_n_a_g_e_m_e_n_t"
	clientmodels "github.com/cloudentity/acp-client-go/clients/obuk/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
)

type OBUKConsentFetcher struct {
	*Server
}

func NewOBUKConsentFetcher(server *Server) *OBUKConsentFetcher {
	return &OBUKConsentFetcher{server}
}

func (o *OBUKConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
	var (
		consents       *obukModels.ListOBConsentsOK
		cs             *system.ListClientsSystemOK
		clientConsents []ClientConsents
		err            error
		cac            []ClientConsents
	)

	if cs, err = o.Client.System.Clients.ListClientsSystem(
		system.NewListClientsSystemParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID),
		nil,
	); err != nil {
		return cac, err
	}

	for _, oc := range cs.Payload.Clients {
		if consents, err = o.Client.Obuk.Management.ListOBConsents(
			obukModels.NewListOBConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingWorkspaceID).
				WithConsentsRequest(&clientmodels.ConsentsRequest{
					ClientID: oc.ClientID,
				}),
			nil,
		); err != nil {
			return cac, err
		}

		if !oc.System {
			clientConsent := ClientConsents{Client: Client{
				ID:           oc.ClientID,
				Name:         oc.ClientName,
				ProviderType: string(OBUK),
			}}
			clientConsent.Consents = o.getConsents(consents)
			clientConsents = append(clientConsents, clientConsent)
		}
	}

	return clientConsents, nil
}

func (o *OBUKConsentFetcher) Revoke(c *gin.Context, revocationType RevocationType, id string) (err error) {
	switch revocationType {
	case ClientRevocation:
		if _, err = o.Client.Obuk.Management.RevokeOBUKConsents(
			obukModels.NewRevokeOBUKConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingWorkspaceID).
				WithConsentTypes(ConsentTypes).
				WithClientID(&id),
			nil,
		); err != nil {
			return err
		}
	case ConsentRevocation:
		if _, err = o.Client.Obuk.Management.RevokeOBUKConsent(
			obukModels.NewRevokeOBUKConsentParamsWithContext(c).
				WithWid(o.Config.OpenbankingWorkspaceID).
				WithConsentID(id),
			nil,
		); err != nil {
			return err
		}
	}

	return nil
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
