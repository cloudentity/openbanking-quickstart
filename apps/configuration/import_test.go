package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImport(t *testing.T) {
	testCases := []struct {
		name       string
		tenantURL  func(port string) string
		tenant     string
		client     *http.Client
		body       []byte
		mode       string
		mockServer *httptest.Server
		err        error
	}{
		{
			name: "import configuration for system tenant",
			tenantURL: func(port string) string {
				return fmt.Sprintf("http://localhost:%s/system", port)
			},
			tenant:     "system",
			client:     http.DefaultClient,
			body:       []byte(`{"tenants": []}`),
			mockServer: serverWithMockedImportEndpoint("/api/system/configuration"),
		},
		{
			name: "import configuration for a specific tenant",
			tenantURL: func(port string) string {
				return fmt.Sprintf("http://localhost:%s/custom", port)
			},
			tenant:     "custom",
			client:     http.DefaultClient,
			body:       []byte(`{"servers": [{"id": "new", "initialize": true}]}`),
			mockServer: serverWithMockedImportEndpoint("/api/system/custom/configuration"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockServer.Start()
			defer tc.mockServer.Close()

			serverURL, err := url.Parse(tc.mockServer.URL)
			require.NoError(tt, err)

			u, err := url.Parse(tc.tenantURL(serverURL.Port()))
			require.NoError(tt, err)

			err = ImportConfiguration(u, &tc.tenant, tc.client, tc.body, tc.mode)

			if tc.err != nil {
				require.Error(tt, err)
			} else {
				require.NoError(tt, err)
			}
		})
	}
}

func serverWithMockedImportEndpoint(path string) *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != "PUT" || req.URL.String() != path {
			rw.WriteHeader(404)
			return
		}

		rw.WriteHeader(204)
	}))
}
