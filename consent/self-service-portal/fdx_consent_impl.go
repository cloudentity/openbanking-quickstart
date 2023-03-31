package main

import (
	"github.com/gin-gonic/gin"

	clientmodels "github.com/cloudentity/acp-client-go/clients/fdx/models"
	"github.com/cloudentity/acp-client-go/clients/fdx/client/m_a_n_a_g_e_m_e_n_t"
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
		consentsResponse *m_a_n_a_g_e_m_e_n_t.ListFDXConsentsOK
		clientsResponse  *system.ListClientsSystemOK
		err              error
		cac              []ClientConsents
		resource         *clientmodels.Resource
		consentRequest   clientmodels.FDXConsentsRequest
	)

	if len(accountIDs) > 0 {
		resource = &clientmodels.Resource{
			ResourceType: "ACCOUNT",
			Ids:          accountIDs,
		}
		consentRequest = clientmodels.FDXConsentsRequest{
			Resource: resource,
		}
	}

	if consentsResponse, err = o.Client.Fdx.Management.ListFDXConsents(
		m_a_n_a_g_e_m_e_n_t.NewListFDXConsentsParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithFDXConsentsRequest(&consentRequest),
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

func (o *FDXConsentImpl) getConsents(response *m_a_n_a_g_e_m_e_n_t.ListFDXConsentsOK) []Consent {
	var consents []Consent

	for _, consent := range response.Payload.Consents {
		var accountIDs []string
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
	revocation := clientmodels.FDXConsentRevocation{
		Initiator: "DATA_ACCESS_PLATFORM",
		Reason:    "BUSINESS_RULE",
	}

	if _, err = o.Client.Fdx.Management.RevokeFDXConsentByID(
		m_a_n_a_g_e_m_e_n_t.NewRevokeFDXConsentByIDParamsWithContext(c).
			WithWid(o.Config.OpenbankingWorkspaceID).
			WithConsentID(id).
			WithRevocationDetails(&revocation),
		nil,
	); err != nil {
		return err
	}
	return nil
}
