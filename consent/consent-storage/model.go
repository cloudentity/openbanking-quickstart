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
	AccountIDs    []string  `json:"account_ids"`
}

type ConsentsByCreatedDate []Consent

func (c ConsentsByCreatedDate) Len() int {
	return len(c)
}

func (c ConsentsByCreatedDate) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ConsentsByCreatedDate) Less(i, j int) bool {
	return c[i].CreatedDate.After(c[j].CreatedDate)
}
