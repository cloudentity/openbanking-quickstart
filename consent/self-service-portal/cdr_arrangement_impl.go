package main

import (
	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	"github.com/gin-gonic/gin"
)

type CDRArrangementImpl struct {
	*Server
}

func NewCDRArrangementImpl(s *Server) ConsentInteractor {
	return &CDRArrangementImpl{s}
}

func (o *CDRArrangementImpl) FetchConsents(c *gin.Context) ([]ClientConsents, error) {
	var (
		response *cdr.ListCDRArrangementsOK
		err      error
		types    []string
		cac      []ClientConsents
		ok       bool
	)

	if types, ok = c.GetQueryArray("types"); !ok {
		types = nil
	}

	if response, err = o.Client.Openbanking.Cdr.ListCDRArrangements(
		cdr.NewListCDRArrangementsParamsWithContext(c).
			WithWid("cdr").
			WithConsentsRequest(&obModels.ConsentsRequest{
				Types: types,
			}),
		nil,
	); err != nil {
		return cac, err
	}

	return MapClientsToConsents(o.getClients(response), o.getConsents(response)), nil
}

func (o *CDRArrangementImpl) getClients(response *cdr.ListCDRArrangementsOK) (clients []Client) {
	for _, arrangement := range response.Payload.Arrangements {
		// TODO: cdr arrangement api does not return any additional client info
		clients = append(clients, Client{
			ID:   arrangement.ClientID,
			Name: "Babaloo",
		})
	}
	return clients
}

func (o *CDRArrangementImpl) getConsents(response *cdr.ListCDRArrangementsOK) (consents []Consent) {
	for _, arrangement := range response.Payload.Arrangements {
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
			// TODO: I think permissions are bound to the scopes in the access token?
			Permissions: []string{},
		})
	}
	return consents
}

func (o *CDRArrangementImpl) RevokeConsent(c *gin.Context, id string) (err error) {
	if _, err = o.Client.Openbanking.Cdr.RevokeCDRArrangement(
		cdr.NewRevokeCDRArrangementParamsWithContext(c).
			WithCdrArrangementID(&id),
	); err != nil {
		return err
	}
	return nil
}
