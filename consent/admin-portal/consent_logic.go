package main

import "fmt"

func ConsentFetcherFactory(spec Spec, server *Server) (ConsentFetchRevoker, error) {
	switch spec {
	case CDR:
		return NewOBCDRConsentFetcher(server), nil
	case OBUK:
		return NewOBUKConsentFetcher(server), nil
	case OBBR:
		return NewOBBRConsentFetcher(server), nil
	default:
		return nil, fmt.Errorf("unsupported spec %s", spec)
	}
}
