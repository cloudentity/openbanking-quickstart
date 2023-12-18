package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
			clientsWithConsents []ClientConsents
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if clientsWithConsents, err = s.ConsentClient.Fetch(c); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to fetch clients: %+v", err))
			return
		}

		resp := ListClientsResponse{Clients: clientsWithConsents}

		c.JSON(http.StatusOK, &resp)
	}
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                 = c.Param("id")
			canBeRevoked       bool
			clientsAndConsents []ClientConsents
			err                error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if clientsAndConsents, err = s.ConsentClient.Fetch(c); err != nil {
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

		if err = s.ConsentClient.Revoke(c, ConsentRevocation, id); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to revoke consent: %+v", err))
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

		if err = s.ConsentClient.Revoke(c, ClientRevocation, id); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to revoke consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) IntrospectToken(_ *gin.Context) error {
	return nil
}
