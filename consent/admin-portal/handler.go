package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	o2Params "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	obuk "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
	system "github.com/cloudentity/acp-client-go/clients/system/client/clients"
)

const (
	AccountAccessConsentType                 = "account_access"
	DomesticPaymentConsentType               = "domestic_payment"
	DomesticScheduledPaymentConsentType      = "domestic_scheduled_payment"
	DomesticStandingOrderConsentType         = "domestic_standing_order"
	InternationalPaymentConsentType          = "international_payment"
	InternationalScheduledPaymentConsentType = "international_scheduled_payment"
	InternationalStandingOrderConsentType    = "international_standing_order"
	FilePaymentConsentType                   = "file_payment"
)

var ConsentTypes = []string{
	AccountAccessConsentType,
	DomesticPaymentConsentType,
	DomesticScheduledPaymentConsentType,
	DomesticStandingOrderConsentType,
	InternationalPaymentConsentType,
	InternationalStandingOrderConsentType,
	InternationalScheduledPaymentConsentType,
	FilePaymentConsentType,
}

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type ListClientsResponse struct {
	Clients []ClientConsents `json:"clients"`
}

func (s *Server) ListClients() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			cs                  *system.ListClientsSystemOK
			consents            *obuk.ListOBConsentsOK
			clientsWithConsents []ClientConsents
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Need to call for each passed in workspace id from env vars
		// TODO: for now just work with UK, later add others
		if cs, err = s.Client.System.Clients.ListClientsSystem(
			system.NewListClientsSystemParamsWithContext(c).
				WithWid(s.Config.SystemClientsServerID),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to list clients from acp: %+v", err))
			return
		}

		//s.Client.Openbanking.Cdr.ListCDRArrangements()
		for _, oc := range cs.Payload.Clients {
			if consents, err = s.Client.Openbanking.Openbankinguk.ListOBConsents(
				obuk.NewListOBConsentsParamsWithContext(c).
					WithWid(s.Config.SystemClientsServerID).
					WithConsentsRequest(&obModels.ConsentsRequest{
						ClientID: oc.ClientID,
					}),
				nil,
			); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to list consents for client: %s, err: %+v", oc.ClientID, err))
				return
			}

			if !oc.System {
				clientCon := ClientConsents{Client: Client{
					ID:   oc.ClientID,
					Name: oc.ClientName,
				}}
				for _, ukConsent := range consents.Payload.Consents {
					log.Println(ukConsent.DomesticPaymentConsent)
					log.Println(ukConsent.DomesticPaymentConsent.Authorisation)
					log.Println(ukConsent.DomesticPaymentConsent.Authorisation.CompletionDateTime)
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
						CompletionDateTime: ukConsent.DomesticPaymentConsent.Authorisation.CompletionDateTime,
					}
					clientCon.Consents = append(clientCon.Consents, con)
				}

				clientsWithConsents = append(clientsWithConsents, clientCon)
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

		if _, err = s.Client.Openbanking.Openbankinguk.RevokeOpenbankingConsent(
			obuk.NewRevokeOpenbankingConsentParamsWithContext(c).
				WithWid(s.Config.SystemClientsServerID).
				WithConsentID(id),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

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

		if _, err = s.Client.Openbanking.Openbankinguk.RevokeOpenbankingConsents(
			obuk.NewRevokeOpenbankingConsentsParamsWithContext(c).
				WithWid(s.Config.SystemClientsServerID).
				WithConsentTypes(ConsentTypes).
				WithClientID(&id),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke consents for client: %s, err: %+v", id, err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) IntrospectToken(c *gin.Context) error {
	var err error

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if _, err = s.IntrospectClient.Oauth2.Oauth2.Introspect(o2Params.NewIntrospectParamsWithContext(c).
		WithToken(&token), nil); err != nil {
		return fmt.Errorf("failed to introspect client: %w", err)
	}

	return nil
}
