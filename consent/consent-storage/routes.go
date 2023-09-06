package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/cloudentity/acp-client-go/clients/system/client/system"
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
		consent.ID = ConsentID(uuid.New().String())
	}
	consent.CreatedDate = time.Now()

	consent.Status = AuthorizedStatus

	if err = s.Repo.Create(consent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, consent)
}

func (s *Server) DeleteConsentHTML(c *gin.Context) {
	var (
		consent Consent
		id      = ConsentID(c.Param("id"))
		err     error
	)

	if consent, err = s.Repo.Get(id); err != nil {
		if errors.Is(err, ErrConsentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if consent.Status == RevokedStatus {
		c.JSON(http.StatusConflict, gin.H{"error": "consent already revoked"})
		return
	}

	consent.Status = RevokedStatus

	if err = s.Repo.Update(consent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	consentID := string(consent.ID)

	if _, err = s.Client.System.System.RevokeTokens(
		system.NewRevokeTokensParams().
			WithConsentID(&consentID).
			WithWid(s.Config.ServerID),
		nil,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

type ConsentData struct {
	Consent
	Pretty string
}

func (s *Server) ListConsentsHTML(c *gin.Context) {
	var (
		consents []Consent
		data     []ConsentData
		err      error
	)

	if consents, err = s.Repo.List(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, consent := range consents {
		var bs []byte

		if bs, err = json.MarshalIndent(consent, "", " "); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		data = append(data, ConsentData{
			Consent: consent,
			Pretty:  string(bs),
		})
	}

	c.HTML(http.StatusOK, "consents.tmpl", gin.H{"consents": data})
}
