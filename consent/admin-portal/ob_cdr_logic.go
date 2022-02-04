package main

import (
	obuk "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type OBCDRConsentFetcher struct {
	server *Server
}

func NewOBCDRConsentFetcher(server *Server) *OBCDRConsentFetcher {
	return &OBCDRConsentFetcher{server}
}

func (o *OBCDRConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
	var (
		cs                  *system.ListClientsSystemOK
		consents            *obuk.ListOBConsentsOK
		clientsWithConsents []ClientConsents
		err                 error
	)

	if cs, err = o.server.Client.System.Clients.ListClientsSystem(
		system.NewListClientsSystemParamsWithContext(c).
			WithWid(o.server.Config.SystemClientsServerID),
		nil,
	); err != nil {
		return clientsWithConsents, errors.Wrap(err, "failed to list clients from acp:")
	}

	for _, oc := range cs.Payload.Clients {
		if consents, err = o.server.Client.Openbanking.Openbankinguk.ListOBConsents(
			obuk.NewListOBConsentsParamsWithContext(c).
				WithWid(o.server.Config.SystemClientsServerID).
				WithConsentsRequest(&obModels.ConsentsRequest{
					ClientID: oc.ClientID,
				}),
			nil,
		); err != nil {
			return clientsWithConsents, errors.Wrap(err, "failed to list consents for client")
		}

		if !oc.System {
			clientCon := ClientConsents{Client: Client{
				ID:   oc.ClientID,
				Name: oc.ClientName,
			}}
			for _, ukConsent := range consents.Payload.Consents {
				// TODO - need to check which consent type
				con := Consent{
					AccountIDs:  ukConsent.AccountIds,
					ConsentID:   ukConsent.ConsentID,
					ClientID:    ukConsent.ClientID,
					TenantID:    ukConsent.TenantID,
					ServerID:    ukConsent.ServerID,
					Status:      ukConsent.Status,
					Type:        string(ukConsent.Type),
					CreatedAt:   ukConsent.CreatedAt,
					ExpiresAt:   ukConsent.AccountAccessConsent.ExpirationDateTime,
					UpdatedAt:   ukConsent.DomesticPaymentConsent.StatusUpdateDateTime,
					Permissions: ukConsent.AccountAccessConsent.Permissions,
					// TODO: Add this - CompletionDateTime: ukConsent.DomesticPaymentConsent.Authorisation.CompletionDateTime,
				}
				clientCon.Consents = append(clientCon.Consents, con)
			}

			clientsWithConsents = append(clientsWithConsents, clientCon)
		}
	}

	return clientsWithConsents, nil
}
