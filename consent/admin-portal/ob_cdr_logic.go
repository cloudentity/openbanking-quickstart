package main

import (
	"errors"

	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
	"github.com/gin-gonic/gin"
)

type OBCDRConsentFetcher struct {
	*Server
}

func NewOBCDRConsentFetcher(server *Server) *OBCDRConsentFetcher {
	return &OBCDRConsentFetcher{server}
}

func (o *OBCDRConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
	var (
		consents       *cdr.ListCDRArrangementsOK
		clientConsents []ClientConsents
		cs             *system.ListClientsSystemOK
		err            error
		cac            []ClientConsents
	)

	if cs, err = o.Client.System.Clients.ListClientsSystem(
		system.NewListClientsSystemParamsWithContext(c).
			WithWid(o.Config.CDRWorkspaceID),
		nil,
	); err != nil {
		return cac, err
	}

	for _, oc := range cs.Payload.Clients {
		if consents, err = o.Client.Openbanking.Cdr.ListCDRArrangements(
			cdr.NewListCDRArrangementsParamsWithContext(c).
				WithWid(o.Config.CDRWorkspaceID).
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

func (o *OBCDRConsentFetcher) getConsents(response *cdr.ListCDRArrangementsOK) []Consent {
	var consents []Consent

	for _, arrangement := range response.Payload.Arrangements {
		if arrangement.Status == "Rejected" {
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

func (o *OBCDRConsentFetcher) Revoke(c *gin.Context, revocationType RevocationType, id string) (err error) {
	switch revocationType {
	case ClientRevocation:
		return errors.New("not yet implemented for client id")
	case ConsentRevocation:
		if _, err = o.Client.Openbanking.Cdr.RevokeCDRArrangementByID(
			cdr.NewRevokeCDRArrangementByIDParamsWithContext(c).
				WithWid(o.Config.CDRWorkspaceID).
				WithArrangementID(id),
			nil,
		); err != nil {
			return err
		}
	}

	return nil
}
