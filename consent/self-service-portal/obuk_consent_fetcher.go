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

func NewOBUKConsentFetcher(s *Server) ConsentFetcher {
	return &OBUKConsentFetcher{s}
}

func (o *OBUKConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
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
			WithWid("openbanking").
			WithConsentsRequest(&obModels.ConsentsRequest{
				Types: types,
			}),
		nil,
	); err != nil {
		return cac, err
	}

	return MapClientsToConsents(o.getClients(response), o.getConsents(response)), nil
}

func (o *OBUKConsentFetcher) getClients(resp *obukModels.ListOBConsentsOK) []Client {
	var clients []Client

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

	return clients
}

func (o *OBUKConsentFetcher) getConsents(resp *obukModels.ListOBConsentsOK) []Consent {
	var consents []Consent

	for _, consent := range resp.Payload.Consents {
		var (
			expiresAt   strfmt.DateTime
			updatedAt   strfmt.DateTime
			permissions []string
		)

		switch consent.Type {
		case "account_access":
			expiresAt = consent.AccountAccessConsent.ExpirationDateTime
			updatedAt = strfmt.DateTime(*consent.AccountAccessConsent.StatusUpdateDateTime)
			permissions = consent.AccountAccessConsent.Permissions
		case "domestic_payment":
			updatedAt = consent.DomesticPaymentConsent.StatusUpdateDateTime
		}

		consents = append(consents, Consent{
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
		})
	}

	return consents
}
