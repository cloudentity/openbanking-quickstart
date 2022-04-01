package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	obbrModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

type OBBRConsentImpl struct {
	*Server
}

func NewOBBRConsentImpl(s *Server) ConsentClient {
	return &OBBRConsentImpl{s}
}

func (o *OBBRConsentImpl) FetchConsents(c *gin.Context, accountIDs []string) ([]ClientConsents, error) {
	var (
		response *obbrModels.ListOBBRConsentsOK
		err      error
		types    []string
		cac      []ClientConsents
		ok       bool
	)

	if types, ok = c.GetQueryArray("types"); !ok {
		types = nil
	}

	if response, err = o.Client.Openbanking.Openbankingbr.ListOBBRConsents(
		obbrModels.NewListOBBRConsentsParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithConsentsRequest(&obModels.OBBRConsentsRequest{
				Types:    types,
				Accounts: accountIDs,
			}),
		nil,
	); err != nil {
		return cac, err
	}

	return MapClientsToConsents(o.getClients(response), o.getConsents(response)), nil
}

func (o *OBBRConsentImpl) getClients(resp *obbrModels.ListOBBRConsentsOK) []Client {
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

func (o *OBBRConsentImpl) getConsents(resp *obbrModels.ListOBBRConsentsOK) []Consent {
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
		case "consents":
			c.ExpiresAt = consent.CustomerDataAccessConsent.ExpirationDateTime
			c.UpdatedAt = consent.CustomerDataAccessConsent.StatusUpdateDateTime
			c.Permissions = obbrPermissionsToStringSlice(consent.CustomerDataAccessConsent.Permissions)
		case "payments":
			// TODO:
		}

		consents = append(consents, c)
	}

	return consents
}

func obbrPermissionsToStringSlice(permissions []obModels.OpenbankingBrasilConsentPermission1) []string {
	ret := make([]string, len(permissions))
	for idx, perm := range permissions {
		ret[idx] = string(perm)
	}
	return ret
}

func (o *OBBRConsentImpl) RevokeConsent(c *gin.Context, consentID string) (err error) {
	if _, err = o.Client.Openbanking.Openbankingbr.RevokeOBBRConsent(
		obbrModels.NewRevokeOBBRConsentParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithConsentID(consentID),
		nil,
	); err != nil {
		return err
	}
	return nil
}
