package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) CreateConsent(c *gin.Context) {
	var (
		consent Consent
		err     error
	)

	if err = c.ShouldBindJSON(&consent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if consent.ID == "" {
		consent.ID = uuid.New().String()
	}

	if err = s.Repo.Create(consent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, consent)
}

type ConsentsReponse struct {
	Consents []Consent `json:"consents"`
}

func (s *Server) ListConsents(c *gin.Context) {
	var (
		res ConsentsReponse
		err error
	)

	if res.Consents, err = s.Repo.List(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
