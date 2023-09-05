package main

type Consent struct {
	ID            string   `json:"id"`
	Subject       string   `json:"subject"`
	GrantedScopes []string `json:"granted_scopes"`
}
