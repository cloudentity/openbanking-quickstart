package main

type FetcherBuilder func(spec Spec, server *Server) ConsentFetcher

func ConsentFetcherFactory(spec Spec, server *Server) ConsentFetcher {
	switch spec {
	case CDR:
		return NewOBCDRConsentFetcher(server)
	case OBUK:
		return NewOBUKConsentFetcher(server)
	default:
		return nil
	}
}
