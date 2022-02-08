package main

import (
	obukModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
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
			WithWid(o.Config.OpenbankingUKWorkspaceID),
		nil,
	); err != nil {
		return cac, err
	}

	for _, oc := range cs.Payload.Clients {
		if consents, err = o.Client.Openbanking.Openbankinguk.ListOBConsents(
			obukModels.NewListOBConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingUKWorkspaceID).
				WithConsentsRequest(&obModels.ConsentsRequest{
					ClientID: oc.ClientID,
				}),
			nil,
		); err != nil {
			return cac, err
		}

		if !oc.System {
			clientConsent := ClientConsents{Client: Client{
				ID:   oc.ClientID,
				Name: oc.ClientName,
			}}
			clientConsent.Consents = o.getConsents(consents)
			clientConsents = append(clientConsents, clientConsent)

		}
	}

	return clientConsents, nil
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

func (o *OBUKConsentFetcher) Revoke(c *gin.Context, revocationType RevocationType, id string) (err error) {
	switch revocationType {
	case ClientRevocation:
		if _, err = o.Client.Openbanking.Openbankinguk.RevokeOpenbankingConsents(
			obukModels.NewRevokeOpenbankingConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingUKWorkspaceID).
				WithConsentTypes(ConsentTypes).
				WithClientID(&id),
			nil,
		); err != nil {
			return err
		}
	case ConsentRevocation:
		if _, err = o.Client.Openbanking.Openbankinguk.RevokeOpenbankingConsent(
			obukModels.NewRevokeOpenbankingConsentParamsWithContext(c).
				WithWid(o.Config.OpenbankingUKWorkspaceID).
				WithConsentID(id),
			nil,
		); err != nil {
			return err
		}
	}

	return nil
}
