package main

type ClientConsents struct {
	Client
	Consents []Consent `json:"consents"`
}
