package main

import "github.com/gin-gonic/gin"

type CDRArrangementFetcher struct {
	*Server
}

func NewCDRArrangementFetcher(s *Server) ConsentFetcher {
	return &CDRArrangementFetcher{s}
}

func (o *CDRArrangementFetcher) Fetch(c *gin.Context) ([]ClientConsents, error) {
	var (
		cac []ClientConsents
	)

	return cac, nil
}

func (o *CDRArrangementFetcher) getClients() []Client {
	return []Client{}
}

func (o *CDRArrangementFetcher) getConsents() []Consent {
	return []Consent{}
}
