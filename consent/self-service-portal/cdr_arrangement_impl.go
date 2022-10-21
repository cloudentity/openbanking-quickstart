package main

import (
	"github.com/gin-gonic/gin"

	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
)

type CDRArrangementImpl struct {
	*Server
}

func NewCDRArrangementImpl(s *Server) ConsentClient {
	return &CDRArrangementImpl{s}
}

func (o *CDRArrangementImpl) FetchConsents(c *gin.Context, accountIDs []string) ([]ClientConsents, error) {
	var (
		arrangementsResponse *cdr.ListCDRArrangementsOK
		clientsResponse      *system.ListClientsSystemOK
		err                  error
		cac                  []ClientConsents
	)

	if arrangementsResponse, err = o.Client.Openbanking.Cdr.ListCDRArrangements(
		cdr.NewListCDRArrangementsParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithConsentsRequest(&obModels.ConsentsRequest{
				Accounts: accountIDs,
			}),
		nil,
	); err != nil {
		return cac, err
	}

	if clientsResponse, err = o.Client.System.Clients.ListClientsSystem(
		system.NewListClientsSystemParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID),
		nil,
	); err != nil {
		return cac, err
	}

	return MapClientsToConsents(o.getClients(clientsResponse), o.getConsents(arrangementsResponse)), nil
}

func (o *CDRArrangementImpl) getClients(response *system.ListClientsSystemOK) []Client {
	var clients Clients

	for _, c := range response.Payload.Clients {
		clients = append(clients, Client{
			ID:        c.ClientID,
			Name:      c.ClientName,
			LogoURI:   c.LogoURI,
			ClientURI: c.ClientURI,
		})
	}
	return clients
}

func (o *CDRArrangementImpl) getConsents(response *cdr.ListCDRArrangementsOK) []Consent {
	var consents []Consent

	for _, arrangement := range response.Payload.Arrangements {
		if arrangement.Status == "AwaitingAuthorisation" {
			continue
		}
		consents = append(consents, Consent{
			AccountIDs: arrangement.AccountIds,
			ConsentID:  string(arrangement.CdrArrangementID),
			TenantID:   arrangement.TenantID,
			ServerID:   arrangement.AuthorizationServerID,
			ClientID:   arrangement.ClientID,
			Status:     string(arrangement.Status),
			Type:       "cdr_arrangement",
			CreatedAt:  arrangement.CreatedAt,
			ExpiresAt:  arrangement.Expiry,
			UpdatedAt:  arrangement.UpdatedAt,
			// permission language is dependent on authorisation scope: https://consumerdatastandardsaustralia.github.io/standards/#banking-language
			// TODO: unmock this
			Permissions: []string{"CommonCustomerBasicRead"},
		})
	}
	return consents
}

func (o *CDRArrangementImpl) RevokeConsent(c *gin.Context, id string) (err error) {
	if _, err = o.Client.Openbanking.Cdr.RevokeCDRArrangementByID(
		cdr.NewRevokeCDRArrangementByIDParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithArrangementID(id),
		nil,
	); err != nil {
		return err
	}
	return nil
}
