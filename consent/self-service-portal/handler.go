package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	o2Params "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
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

func (c *ClientConsents) HasConsentID(consentID string) bool {
	for _, consent := range c.Consents {
		if consent.ConsentID == consentID {
			return true
		}
	}
	return false
}

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LogoURI   string `json:"logo_uri"`
	ClientURI string `json:"client_uri"`
}

type Clients []Client

func (c Clients) Unique() []Client {
	var clients []Client
	m := make(map[string]bool)

	for _, client := range c {
		if _, exists := m[client.ID]; exists {
			continue
		}
		m[client.ID] = true
		clients = append(clients, client)
	}

	return clients
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

	DebtorAccountIdentification string `json:"DebtorAccountIdentification"`
	DebtorAccountName           string `json:"DebtorAccountName"`

	CreditorAccountIdentification string `json:"CreditorAccountIdentification"`
	CreditorAccountName           string `json:"CreditorAccountName"`

	Currency string `json:"Currency"`
	Amount   string `json:"Amount"`

	CompletionDateTime strfmt.DateTime `json:"CompletionDateTime"`
}

func MapClientsToConsents(clients []Client, consents []Consent) []ClientConsents {
	var (
		consentMap        = make(map[string][]Consent)
		clientAndConsents []ClientConsents
	)

	for _, consent := range consents {
		if _, ok := consentMap[consent.ClientID]; !ok {
			consentMap[consent.ClientID] = []Consent{}
		}
		consentMap[consent.ClientID] = append(consentMap[consent.ClientID], consent)
	}

	for _, client := range clients {
		consents := consentMap[client.ID]

		if len(consents) == 0 {
			continue
		}

		clientAndConsents = append(clientAndConsents, ClientConsents{
			Client:   client,
			Consents: consents,
		})
	}

	return clientAndConsents
}

func (s *Server) ListConsents() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			sub                    string
			accounts, acc          InternalAccounts
			clientsAndConsents, cc []ClientConsents
			err                    error
		)

		if sub, err = s.GetSubject(c); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		for _, spec := range []Spec{OBUK, CDR} {
			if acc, err = s.BankClients[spec].GetInternalAccounts(sub); err != nil {
				Error(c, ToAPIError(err))
				return
			}

			accounts.Accounts = append(accounts.Accounts, acc.Accounts...)

			if cc, err = s.ConsentClients[spec].FetchConsents(c, accounts.GetAccountIDs()); err != nil {
				Error(c, ToAPIError(err))
				return
			}

			clientsAndConsents = append(clientsAndConsents, cc...)
		}

		c.JSON(http.StatusOK, &ConsentsResponse{
			ClientConsents: clientsAndConsents,
			Accounts:       accounts,
		})
	}
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                 = c.Param("id")
			consentType        = c.Query("consent_type")
			clientsAndConsents []ClientConsents
			canBeRevoked       bool
			consentClient      ConsentClient = s.GetConsentClientByConsentType(consentType)
			err                error
		)

		if consentClient == nil {
			Error(c, APIError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("unable to retrieve consent client for consent type [%s]", consentType),
			})
			return
		}

		if clientsAndConsents, err = consentClient.FetchConsents(c, []string{}); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		for _, c := range clientsAndConsents {
			if c.HasConsentID(id) {
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

		if err = consentClient.RevokeConsent(c, id); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		c.Status(http.StatusNoContent)
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

func (s *Server) GetConsentClientByConsentType(consentType string) ConsentClient {
	switch consentType {
	case "account_access", "domestic_payment":
		return s.ConsentClients[OBUK]
	case "cdr_arrangement":
		return s.ConsentClients[CDR]
	}
	return nil
}

func (s *Server) GetBankClientByConsentType(consentType string) BankClient {
	switch consentType {
	case "account_access", "domestic_payment":
		return s.BankClients[OBUK]
	case "cdr_arrangement":
		return s.BankClients[CDR]
	}
	return nil
}
