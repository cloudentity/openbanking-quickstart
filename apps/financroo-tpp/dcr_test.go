package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterClient(t *testing.T) {
	resp, err := RegisterClient(
		context.Background(),
		Config{
			ACPInternalURL: "https://localhost:8443",
			Tenant:         "default",
			ServerID:       "generic",
			UIURL:          "https://localhost:8091",
			CertFile:       "../../data/tpp_cert.pem",
			Spec:           GENERIC,
			RootCA:         "../../data/ca.pem",
			EnableDCR:      true,
		},
	)
	require.NoError(t, err)
	require.NotEmpty(t, resp.ClientID)
}
