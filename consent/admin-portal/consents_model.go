package main

import "github.com/go-openapi/strfmt"

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
