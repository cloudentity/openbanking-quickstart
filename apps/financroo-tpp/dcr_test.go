package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterClient(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			payload    interface{}
			statusCode int
		)

		switch r.URL.String() {
		case "/default/generic/.well-known/openid-configuration":
			statusCode = http.StatusOK
			payload = map[string]interface{}{
				"issuer": "http://localhost:8080",
			}
		case "/default/generic/oauth2/register":
			statusCode = http.StatusCreated
			payload = map[string]interface{}{
				"client_id": "123",
			}
		default:
			statusCode = http.StatusNotFound
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		if payload != nil {
			json.NewEncoder(w).Encode(payload)
		}
	})

	server := httptest.NewServer(mockHandler)
	defer server.Close()

	resp, err := RegisterClient(
		context.Background(),
		Config{
			ACPInternalURL: server.URL,
			// ACPInternalURL: "https://localhost:8443",
			Tenant:    "default",
			ServerID:  "generic",
			UIURL:     "https://localhost:8091",
			CertFile:  "../../data/tpp_cert.pem",
			Spec:      GENERIC,
			RootCA:    "../../data/ca.pem",
			EnableDCR: true,
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, resp.ClientID)
}
