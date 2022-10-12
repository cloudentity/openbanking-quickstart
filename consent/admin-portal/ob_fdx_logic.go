package main

import (
	"github.com/gin-gonic/gin"
)

type OBFDXConsentFetcher struct {
	*Server
}

func NewOBFDXConsentFetcher(server *Server) *OBFDXConsentFetcher {
	return &OBFDXConsentFetcher{server}
}

func (o *OBFDXConsentFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {

	return []ClientConsents{}, nil
}

func (o *OBFDXConsentFetcher) Revoke(c *gin.Context, revocationType RevocationType, id string) (err error) {
	return nil
}
