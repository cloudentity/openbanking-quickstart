package main

import "github.com/gin-gonic/gin"

type ConsentFetcher interface {
	Fetch(c *gin.Context) ([]ClientConsents, error)
}

type ConsentRevoker interface {
	Revoke(c *gin.Context) ()
}
