package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	o2Params "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

const (
	ConsentsConsentType                      = "consents"
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
			clientsWithConsents []ClientConsents
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		for _, cc := range s.ConsentClients {
			var consents []ClientConsents
			if consents, err = cc.Fetch(c); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to fetch clients: %+v", err))
				return
			}
			clientsWithConsents = append(clientsWithConsents, consents...)
		}

		resp := ListClientsResponse{Clients: clientsWithConsents}

		c.JSON(http.StatusOK, &resp)
	}
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                  = c.Param("id")
			canBeRevoked        bool
			consentType         = c.Query("consent_type")
			clientsAndConsents  []ClientConsents
			consentFetchRevoker ConsentFetchRevoker = s.GetConsentClientByConsentType(consentType)
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if consentFetchRevoker == nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("unable to retrieve consent client for consent type [%s]", consentType))
			return
		}

		if clientsAndConsents, err = consentFetchRevoker.Fetch(c); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to fetch clients: %+v", err))
			return
		}

		for _, c := range clientsAndConsents {
			if c.HasConsentID(id) {
				canBeRevoked = true
				break
			}
		}

		if !canBeRevoked {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to fetch clients: %+v", err))
			return
		}

		if err = consentFetchRevoker.Revoke(c, ConsentRevocation, id); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to revoke consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) RevokeConsentsForClient() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                                      = c.Param("id")
			providerType                            = c.Query("provider_type")
			consentFetchRevoker ConsentFetchRevoker = s.GetConsentClientByProviderType(providerType)
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if consentFetchRevoker == nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("unable to retrieve consent client for consent type [%s]", providerType))
			return
		}

		if err = consentFetchRevoker.Revoke(c, ClientRevocation, id); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to revoke consent: %+v", err))
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

func (s *Server) GetConsentClientByConsentType(consentType string) ConsentFetchRevoker {
	switch consentType {
	case "account_access", "domestic_payment":
		for _, fetcherRevoker := range s.ConsentClients {
			if _, ok := fetcherRevoker.(*OBUKConsentFetcher); ok {
				return fetcherRevoker
			}
		}
	case "cdr_arrangement":
		for _, fetcherRevoker := range s.ConsentClients {
			if _, ok := fetcherRevoker.(*OBCDRConsentFetcher); ok {
				return fetcherRevoker
			}
		}

	case "consents":
		for _, fetcherRevoker := range s.ConsentClients {
			if _, ok := fetcherRevoker.(*OBBRConsentFetcher); ok {
				return fetcherRevoker
			}
		}
	}
	return nil
}

func (s *Server) GetConsentClientByProviderType(providerType string) ConsentFetchRevoker {
	switch providerType {
	case string(OBUK):
		for _, fetcherRevoker := range s.ConsentClients {
			if _, ok := fetcherRevoker.(*OBUKConsentFetcher); ok {
				return fetcherRevoker
			}
		}
	case string(CDR):
		for _, fetcherRevoker := range s.ConsentClients {
			if _, ok := fetcherRevoker.(*OBCDRConsentFetcher); ok {
				return fetcherRevoker
			}
		}
	case string(OBBR):
		for _, fetcherRevoker := range s.ConsentClients {
			if _, ok := fetcherRevoker.(*OBBRConsentFetcher); ok {
				return fetcherRevoker
			}
		}
	}
	return nil
}
