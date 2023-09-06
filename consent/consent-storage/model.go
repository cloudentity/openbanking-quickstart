package main

import "time"

type ConsentID string

type Status string

var (
	AuthorizedStatus Status = "authorized"
	RevokedStatus    Status = "revoked"
)

type Consent struct {
	ID            ConsentID `json:"id"`
	CreatedDate   time.Time `json:"created_date"`
	Status        Status    `json:"status"`
	Subject       string    `json:"subject"`
	GrantedScopes []string  `json:"granted_scopes"`
}
