package main

import (
	"context"
	"encoding/json"
	"fmt"
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
			err := json.NewEncoder(w).Encode(payload)
			require.NoError(t, err)
		}
	})

	server := httptest.NewServer(mockHandler)
	defer server.Close()

	resp, err := RegisterClient(
		context.Background(),
		Config{
			UIURL:    "https://localhost:8091",
			CertFile: "../../data/tpp_cert.pem",
			Spec:     GENERIC,
			RootCA:   "../../data/ca.pem",
		},
		BankConfig{
			ACPInternalURL: server.URL,
			// ACPInternalURL: "https://localhost:8443",
			Tenant:    "default",
			Server:    "generic",
			EnableDCR: true,
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, resp.ClientID)
}

func TestClientIDStorage(t *testing.T) {
	dir := t.TempDir()

	db, err := InitDB(Config{DBFile: fmt.Sprintf("%s/test.db", dir)})
	require.NoError(t, err)
	defer db.Close()

	storage := DCRClientIDStorage{DB: db}

	bankID := BankID("gobank")

	id, exists, err := storage.Get(bankID)
	require.NoError(t, err)
	require.False(t, exists)
	require.Empty(t, id)

	err = storage.Set(bankID, "client-123")
	require.NoError(t, err)

	id, exists, err = storage.Get(bankID)
	require.NoError(t, err)
	require.True(t, exists)
	require.Equal(t, "client-123", id)
}
