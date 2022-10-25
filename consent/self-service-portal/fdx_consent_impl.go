package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp-client-go/clients/openbanking/client/f_d_x"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
)

type FDXConsentImpl struct {
	*Server
}

func NewFDXConsentImpl(s *Server) ConsentClient {
	return &FDXConsentImpl{s}
}

func (o *FDXConsentImpl) FetchConsents(c *gin.Context, accountIDs []string) ([]ClientConsents, error) {
	var (
		consentsResponse *f_d_x.ListFDXConsentsOK
		clientsResponse  *system.ListClientsSystemOK
		err              error
		cac              []ClientConsents
	)

	if consentsResponse, err = o.Client.Openbanking.Fdx.ListFDXConsents(
		f_d_x.NewListFDXConsentsParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithFDXConsentsRequest(&obModels.FDXConsentsRequest{
				Resource: &obModels.Resource{
					ResourceType: "ACCOUNT",
					Ids:          accountIDs,
				},
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

	return MapClientsToConsents(o.getClients(clientsResponse), o.getConsents(consentsResponse)), nil
}

func (o *FDXConsentImpl) getClients(response *system.ListClientsSystemOK) []Client {
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

func (o *FDXConsentImpl) getConsents(response *f_d_x.ListFDXConsentsOK) []Consent {
	var (
		consents   []Consent
		accountIDs []string
	)
	for _, consent := range response.Payload.Consents {
		if consent.Status == "Rejected" || consent.Status == "Revoked" {
			continue
		}

		for _, resource := range consent.GrantedResources {
			if resource.ResourceType == "ACCOUNT" {
				accountIDs = append(accountIDs, resource.ID)
			}
		}

		consents = append(consents, Consent{
			AccountIDs: accountIDs,
			ConsentID:  string(consent.ID),
			TenantID:   consent.TenantID,
			ServerID:   consent.AuthorizationServerID,
			ClientID:   consent.ClientID,
			Status:     string(consent.Status),
			CreatedAt:  consent.CreatedTime,
			ExpiresAt:  consent.ExpirationTime,
			UpdatedAt:  consent.UpdatedTime,
			Type:       "fdx_consent",
			// permission language is dependent on authorisation scope: https://consumerdatastandardsaustralia.github.io/standards/#banking-language
			// TODO: unmock this
			Permissions: []string{"CommonCustomerBasicRead"},
		})
	}
	return consents
}

func (o *FDXConsentImpl) RevokeConsent(c *gin.Context, id string) (err error) {
	revocation := obModels.FDXConsentRevocation{
		Initiator: "DATA_ACCESS_PLATFORM",
		Reason:    "BUSINESS_RULE",
	}
	if _, err = o.Client.Openbanking.Fdx.RevokeFDXConsent(
		f_d_x.NewRevokeFDXConsentParamsWithContext(c).
			WithConsentID(id).
			WithConsentRevocation(&revocation),
		nil,
	); err != nil {
		return err
	}
	return nil
}
