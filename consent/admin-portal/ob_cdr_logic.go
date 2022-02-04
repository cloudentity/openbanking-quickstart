package main

import (
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
		arrangementsResponse *cdr.ListCDRArrangementsOK
		clientsResponse      *system.ListClientsSystemOK
		err                  error
		cac                  []ClientConsents
	)

	if arrangementsResponse, err = o.Client.Openbanking.Cdr.ListCDRArrangements(
		cdr.NewListCDRArrangementsParamsWithContext(c).
			WithWid(o.Config.CDRWorkspaceID).
			WithConsentsRequest(&obModels.ConsentsRequest{
				Accounts: accountIDs,
			}),
		nil,
	); err != nil {
		return cac, err
	}

	if clientsResponse, err = o.Client.System.Clients.ListClientsSystem(
		system.NewListClientsSystemParamsWithContext(c).
			WithWid(o.Config.CDRWorkspaceID),
		nil,
	); err != nil {
		return cac, err
	}

	return MapClientsToConsents(o.getClients(clientsResponse), o.getConsents(arrangementsResponse)), nil
}
