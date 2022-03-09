package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	obbrModels "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
)

type OBBRConsentFetcher struct {
	*Server
}

func NewOBBRConsentFetcher(server *Server) *OBBRConsentFetcher {
	return &OBBRConsentFetcher{server}
}

func (o *OBBRConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
	var (
		consents       *obbrModels.ListOBBRConsentsOK
		cs             *system.ListClientsSystemOK
		clientConsents []ClientConsents
		err            error
		cac            []ClientConsents
	)

	if cs, err = o.Client.System.Clients.ListClientsSystem(
		system.NewListClientsSystemParamsWithContext(c).
			WithWid(o.Config.OpenbankingBrasilWorkspaceID),
		nil,
	); err != nil {
		return cac, err
	}

	for _, oc := range cs.Payload.Clients {
		if consents, err = o.Client.Openbanking.Openbankingbr.ListOBBRConsents(
			obbrModels.NewListOBBRConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingBrasilWorkspaceID).
				WithConsentsRequest(&obModels.OBBRConsentsRequest{
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
				ProviderType: string(OBBR),
			}}
			clientConsent.Consents = o.getConsents(consents)
			clientConsents = append(clientConsents, clientConsent)
		}
	}

	return clientConsents, nil
}

func (o *OBBRConsentFetcher) Revoke(c *gin.Context, revocationType RevocationType, id string) (err error) {
	switch revocationType {
	case ClientRevocation:
		if _, err = o.Client.Openbanking.Openbankingbr.RevokeOBBRConsents(
			obbrModels.NewRevokeOBBRConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingBrasilWorkspaceID).
				WithConsentTypes([]string{"consents"}).
				WithClientID(&id),
			nil,
		); err != nil {
			return err
		}
	case ConsentRevocation:
		if _, err = o.Client.Openbanking.Openbankingbr.RevokeOBBRConsent(
			obbrModels.NewRevokeOBBRConsentParamsWithContext(c).
				WithWid(o.Config.OpenbankingBrasilWorkspaceID).
				WithConsentID(id),
			nil,
		); err != nil {
			return err
		}
	}

	return nil
}

func (o *OBBRConsentFetcher) getConsents(resp *obbrModels.ListOBBRConsentsOK) []Consent {
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
		}

		consents = append(consents, c)
	}

	logrus.Infof("consents %+v", consents)

	return consents
}

func obbrPermissionsToStringSlice(permissions []obModels.OpenbankingBrasilConsentPermission1) []string {
	ret := make([]string, len(permissions))
	for idx, perm := range permissions {
		ret[idx] = string(perm)
	}
	return ret
}
