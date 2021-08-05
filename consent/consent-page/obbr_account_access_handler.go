package main

import (
	"github.com/gin-gonic/gin"
)

// TODO: acp client will need to be updated for this and will also need obbr generated code

type OBBRAccountAccessConsentHandler struct {
	*Server
	ConsentTools
}

func (s *OBBRAccountAccessConsentHandler) GetConsent(c *gin.Context, loginRequest LoginRequest) {
}

func (s *OBBRAccountAccessConsentHandler) ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	return "", nil
}

func (s *OBBRAccountAccessConsentHandler) DenyConsent(c *gin.Context, loginRequest LoginRequest) (string, error) {
	return "", nil
}

var _ ConsentHandler = &OBBRAccountAccessConsentHandler{}
