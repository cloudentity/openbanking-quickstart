package main

func ConsentFetcherFactory(spec Spec, server *Server) ConsentFetchRevoker {
	switch spec {
	case CDR:
		return NewOBCDRConsentFetcher(server)
	case OBUK:
		return NewOBUKConsentFetcher(server)
	case OBBR:
		return NewOBBRConsentFetcher(server)
	default:
		return nil
	}
}
