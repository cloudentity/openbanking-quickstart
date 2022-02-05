package main

func ConsentFetcherFactory(spec Spec, server *Server) ConsentFetcher {
	switch spec {
	case CDR:
		return NewOBCDRConsentFetcher(server)
	case OBUK:
		return NewOBCDRConsentFetcher(server)
	default:
		return nil
	}
}
