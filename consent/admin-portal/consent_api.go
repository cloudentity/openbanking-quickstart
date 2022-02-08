package main

import "github.com/gin-gonic/gin"

type RevocationType int

const (
	ClientRevocation RevocationType = iota
	ConsentRevocation
)

type ConsentFetcher interface {
	Fetch(c *gin.Context) ([]ClientConsents, error)
}

type ConsentRevoker interface {
	Revoke(c *gin.Context, revocationType RevocationType, id string) error
}

type ConsentFetchRevoker interface {
	ConsentFetcher
	ConsentRevoker
}
