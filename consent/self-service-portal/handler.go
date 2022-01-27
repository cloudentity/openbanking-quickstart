package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	o2Params "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type ConsentsResponse struct {
	ClientConsents []ClientConsents `json:"client_consents"`
	Accounts       InternalAccounts `json:"accounts"`
}

type ClientConsents struct {
	Client
	Consents []Consent `json:"consents"`
}

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LogoURI   string `json:"logo_uri"`
	ClientURI string `json:"client_uri"`
}

type Consent struct {
	AccountIDs  []string        `json:"AccountIDs"`
	ConsentID   string          `json:"ConsentID"`
	ClientID    string          `json:"client_id"`
	TenantID    string          `json:"tenant_id"`
	ServerID    string          `json:"server_id"`
	Status      string          `json:"Status"`
	Type        string          `json:"type"`
	CreatedAt   strfmt.DateTime `json:"CreationDateTime"`
	ExpiresAt   strfmt.DateTime `json:"ExpirationDateTime"`
	UpdatedAt   strfmt.DateTime `json:"StatusUpdateDateTime"`
	Permissions []string        `json:"Permissions"`
}

func MapClientsToConsents(clients []Client, consents []Consent) []ClientConsents {
	consentMap := make(map[string][]Consent)
	for _, consent := range consents {
		if _, ok := consentMap[consent.ClientID]; !ok {
			consentMap[consent.ClientID] = []Consent{}
		}
		consentMap[consent.ClientID] = append(consentMap[consent.ClientID], consent)
	}

	var clientAndConsents []ClientConsents
	for _, client := range clients {
		clientAndConsents = append(clientAndConsents, ClientConsents{
			Client:   client,
			Consents: consentMap[client.ID],
		})
	}

	return clientAndConsents
}

func (s *Server) ListConsents() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			sub                string
			accounts           InternalAccounts
			clientsAndConsents []ClientConsents
			err                error
		)

		if sub, err = s.GetSubject(c); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		if accounts, err = s.BankClient.GetInternalAccounts(sub); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		fetcher := NewOBUKConsentFetcher(s)

		if clientsAndConsents, err = fetcher.Fetch(c); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		c.JSON(http.StatusOK, &ConsentsResponse{
			ClientConsents: clientsAndConsents,
			Accounts:       accounts,
		})
	}
}

type ConsentsAndAccounts struct {
	Consents []*obModels.OBUKConsentWithClient `json:"consents"`
	Accounts InternalAccounts                  `json:"accounts"`
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		/*	var (
				id                 = c.Param("id")
				consentsByAccounts *ConsentsAndAccounts
				canBeRevoked       bool
				err                error
			)

			if consentsByAccounts, err = s.FetchConsents(c); err != nil {
				Error(c, ToAPIError(err))
				return
			}

			for _, c := range consentsByAccounts.Consents {
				if c.ConsentID == id {
					canBeRevoked = true
					break
				}
			}

			if !canBeRevoked {
				Error(c, APIError{
					Code:    http.StatusUnauthorized,
					Message: "user is not authorized to revoke this consent",
				})
				return
			}

			if _, err = s.Client.Openbanking.Openbankinguk.RevokeOpenbankingConsent(
				obukModels.NewRevokeOpenbankingConsentParamsWithContext(c).
					WithWid(s.Config.SystemClientsServerID).
					WithConsentID(id),
				nil,
			); err != nil {
				Error(c, ToAPIError(err))
				return
			}

			c.Status(http.StatusNoContent)*/
	}
}

var (
	ErrTokenNotActive = errors.New("token is not active")
	ErrTokenMissing   = errors.New("token is missing")
)

func (s *Server) GetSubject(c *gin.Context) (string, error) {
	var (
		introspectResp *o2Params.IntrospectOK
		err            error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if token == "" {
		return "", ErrTokenMissing
	}

	if introspectResp, err = s.IntrospectClient.Oauth2.Oauth2.Introspect(o2Params.NewIntrospectParamsWithContext(c).
		WithToken(&token), nil); err != nil {
		return "", err
	}

	if !introspectResp.Payload.Active {
		return "", ErrTokenNotActive
	}

	return introspectResp.Payload.Sub, nil
}

func ToAPIError(err error) APIError {
	if errors.Is(err, ErrTokenNotActive) {
		return APIError{
			http.StatusUnauthorized,
			"token is not valid",
			nil,
		}
	} else if errors.Is(err, ErrTokenMissing) {
		return APIError{
			http.StatusUnauthorized,
			"token is missing",
			nil,
		}
	}

	return APIError{
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		err,
	}
}
