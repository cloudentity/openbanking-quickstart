package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp-client-go/client/clients"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type Client struct {
	ID       string                                              `json:"client_id"`
	Name     string                                              `json:"client_name,omitempty"`
	Consents []*models.OpenbankingAccountAccessConsentWithClient `json:"consents"`
}

type ListClientsResponse struct {
	Clients []Client `json:"clients"`
}

func (s *Server) ListClients() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			cs                  *clients.ListClientsSystemOK
			consents            *openbanking.ListAccountAccessConsentsOK
			clientsWithConsents []Client
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if cs, err = s.Client.Clients.ListClientsSystem(
			clients.NewListClientsSystemParams().
				WithTid(s.Client.TenantID).
				WithAid(s.Config.SystemClientsServerID),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to list clients from acp: %+v", err))
			return
		}

		for _, oc := range cs.Payload.Clients {
			if consents, err = s.Client.Openbanking.ListAccountAccessConsents(
				openbanking.NewListAccountAccessConsentsParams().
					WithTid(s.Client.TenantID).
					WithAid(s.Config.SystemClientsServerID).
					WithListAccountAccessConsentsRequest(&models.ListAccountAccessConsentsRequest{
						ClientID: oc.ID,
					}),
				nil,
			); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to list consents for client: %s, err: %+v", oc.ID, err))
				return
			}

			if !oc.System {
				clientsWithConsents = append(clientsWithConsents, Client{
					ID:       oc.ID,
					Name:     oc.Name,
					Consents: consents.Payload.Consents,
				})
			}
		}

		resp := ListClientsResponse{Clients: clientsWithConsents}

		c.JSON(http.StatusOK, &resp)
	}
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id  = c.Param("id")
			err error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if _, err = s.Client.Openbanking.RevokeOpenbankingConsent(
			openbanking.NewRevokeOpenbankingConsentParams().
				WithTid(s.Client.TenantID).
				WithAid(s.Config.SystemClientsServerID).
				WithConsentID(id),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

var accountAccessConsentType = "account_access"

func (s *Server) RevokeConsentsForClient() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id  = c.Param("id")
			err error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if _, err = s.Client.Openbanking.RevokeOpenbankingConsents(
			openbanking.NewRevokeOpenbankingConsentsParams().
				WithTid(s.Client.TenantID).
				WithAid(s.Config.SystemClientsServerID).
				WithClientID(&id).
				WithConsentType(&accountAccessConsentType),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consents for client: %s, err: %+v", id, err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) IntrospectToken(c *gin.Context) error {
	var err error

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if _, err = s.IntrospectClient.IntrospectToken(token); err != nil {
		return fmt.Errorf("failed to introspect client: %w", err)
	}

	return nil
}
