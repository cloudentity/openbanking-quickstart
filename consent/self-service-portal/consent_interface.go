package main

import "github.com/gin-gonic/gin"

type ConsentInteractor interface {
	ConsentFetcher
	ConsentRevoker
}

type ConsentFetcher interface {
	FetchConsents(c *gin.Context) ([]ClientConsents, error)
}

type ConsentRevoker interface {
	RevokeConsent(ctx *gin.Context, id string) error
}
