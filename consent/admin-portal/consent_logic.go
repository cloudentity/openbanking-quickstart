package main

func ConsentFetcherFactory(spec Spec, server *Server) ConsentFetchRevoker {
	switch spec {
	case CDR:
		return NewOBCDRConsentFetcher(server)
	case OBUK:
		return NewOBUKConsentFetcher(server)
	default:
		return nil
	}
}