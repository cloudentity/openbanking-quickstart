package main

func ConsentFetcherFactory(spec string, server *Server) ConsentFetcher {
	switch spec {
	case "obcdr":
		return NewOBCDRConsentFetcher(server)
	case "obuk":
		return NewOBCDRConsentFetcher(server)
	default:
		return nil
	}
}
