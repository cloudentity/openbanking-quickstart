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

	// Removed Currency
	Amount   string `json:"Amount"`

	CompletionDateTime strfmt.DateTime `json:"CompletionDateTime"`
}

type ClientConsents struct {
	Client
	Consents []Consent `json:"consents"`
}

type Client struct {
	ID   string `json:"client_id"`
	Name string `json:"client_name,omitempty"`
}
