package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Consent struct {
	ID      string   `json:"id"`
	Subject string   `json:"subject"`
	Scopes  []string `json:"scopes"`
}

// dummy in-memory implementation
var consents []Consent

// createConsent handles the creation of a new consent.
func createConsent(c *gin.Context) {
	var consent Consent
	if err := c.ShouldBindJSON(&consent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a random UUID for the consent
	consent.ID = uuid.New().String()

	// Append the new consent to the consents list
	consents = append(consents, consent)

	c.JSON(http.StatusCreated, consent)
}

// listConsents lists all consents.
func listConsents(c *gin.Context) {
	c.JSON(http.StatusOK, consents)
}
