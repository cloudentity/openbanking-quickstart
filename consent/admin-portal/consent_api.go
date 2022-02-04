package main

import "github.com/gin-gonic/gin"

type ConsentFetcher interface {
	Fetch(c *gin.Context, accountIDs []string) ([]ClientConsents, error)
}
