package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp-client-go/clients/fdx/client/m_a_n_a_g_e_m_e_n_t"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
	clientmodels "github.com/cloudentity/acp-client-go/clients/fdx/models"
)

type OBFDXConsentFetcher struct {
	*Server
}

func NewOBFDXConsentFetcher(server *Server) *OBFDXConsentFetcher {
	return &OBFDXConsentFetcher{server}
}

func (o *OBFDXConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
	var (
		consents       *m_a_n_a_g_e_m_e_n_t.ListFDXConsentsOK
		clientConsents []ClientConsents
		cs             *system.ListClientsSystemOK
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
		if consents, err = o.Client.Fdx.Management.ListFDXConsents(
			m_a_n_a_g_e_m_e_n_t.NewListFDXConsentsParamsWithContext(c).
				WithWid(o.Config.OpenbankingWorkspaceID).
				WithFDXConsentsRequest(&clientmodels.FDXConsentsRequest{
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
				ProviderType: string(FDX),
			}}
			clientConsent.Consents = o.getConsents(consents)
			clientConsents = append(clientConsents, clientConsent)
		}
	}

	return clientConsents, nil
}

func (o *OBFDXConsentFetcher) Revoke(c *gin.Context, revocationType RevocationType, id string) (err error) {
	switch revocationType {
	case ClientRevocation:
		revocation := clientmodels.ConsentRevocationByCLientID{
			ClientID: id,
			RevocationDetails: &clientmodels.FDXConsentRevocation{
				Initiator: "DATA_ACCESS_PLATFORM",
				Reason:    "BUSINESS_RULE",
			},
		}
		if _, err = o.Client.Fdx.Management.RevokeFDXConsents(
			m_a_n_a_g_e_m_e_n_t.NewRevokeFDXConsentsParamsWithContext(c).
				WithConsentRevocationByClientID(&revocation).
				WithWid(o.Config.OpenbankingWorkspaceID),
			nil,
		); err != nil {
			return err
		}

	case ConsentRevocation:
		revocation := clientmodels.FDXConsentRevocation{
			Initiator: "DATA_ACCESS_PLATFORM",
			Reason:    "BUSINESS_RULE",
		}

		if _, err = o.Client.Fdx.Management.RevokeFDXConsentByID(
			m_a_n_a_g_e_m_e_n_t.NewRevokeFDXConsentByIDParamsWithContext(c).
				WithRevocationDetails(&revocation).
				WithWid(o.Config.OpenbankingWorkspaceID).
				WithConsentID(id),
			nil,
		); err != nil {
			return err
		}
	}

	return nil
}

func (o *OBFDXConsentFetcher) getConsents(response *m_a_n_a_g_e_m_e_n_t.ListFDXConsentsOK) []Consent {
	var (
		consents   []Consent
		accountIDs []string
	)

	for _, consent := range response.Payload.Consents {
		if consent.Status == "Rejected" {
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
