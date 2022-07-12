package main

import "github.com/go-openapi/strfmt"

type Consent struct {
	AccountIDs         []string        `json:"account_ids"`
	ConsentID          string          `json:"consent_id"`
	ClientID           string          `json:"client_id"`
	TenantID           string          `json:"tenant_id"`
	ServerID           string          `json:"server_id"`
	Status             string          `json:"status"`
	Type               string          `json:"consent_type"`
	CreatedAt          strfmt.DateTime `json:"created_at"`
	ExpiresAt          strfmt.DateTime `json:"expires_at"`
	UpdatedAt          strfmt.DateTime `json:"updated_at"`
	CompletionDateTime strfmt.DateTime `json:"completed_at"`

	Permissions []string `json:"permissions"`

	Currency string `json:"currency"`
	Amount   string `json:"Amount"`

	DebtorAccountIdentification string `json:"DebtorAccountIdentification"`
	DebtorAccountName           string `json:"DebtorAccountName"`

	CreditorAccountIdentification string `json:"CreditorAccountIdentification"`
	CreditorAccountName           string `json:"CreditorAccountName"`
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
	ID           string `json:"client_id"`
	Name         string `json:"client_name"`
	LogoURI      string `json:"logo_uri"`
	ClientURI    string `json:"client_uri"`
	ProviderType string `json:"provider_type"`
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
