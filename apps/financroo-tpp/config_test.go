package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	for k, v := range map[string]string{
		"ACP_URL":          "https://localhost:8443",
		"ACP_INTERNAL_URL": "https://acp:8443",
		"APP_HOST":         "localhost",
		"UI_URL":           "https://localhost:8091",
		"CERT_FILE":        "cert.pem",
		"KEY_FILE":         "key.pem",
	} {
		os.Setenv(k, v)
	}

	config, err := LoadConfig()
	require.NoError(t, err)

	require.Equal(t, 8091, config.Port)
	require.Equal(t, "./data/my.db", config.DBFile)
	require.Equal(t, "https://localhost:8443", config.ACPURL)
	require.Equal(t, "https://acp:8443", config.ACPInternalURL)
	require.Equal(t, "localhost", config.AppHost)
	require.Equal(t, "https://localhost:8091", config.UIURL)
	require.Equal(t, "cert.pem", config.CertFile)
	require.Equal(t, "key.pem", config.KeyFile)

	require.NotEmpty(t, config.Login.ClientID)
	require.NotEmpty(t, config.Login.ServerID)
	require.NotEmpty(t, config.Login.TenantID)
	require.NotEmpty(t, config.Login.RootCA)
	require.NotEmpty(t, config.Login.Timeout)

	require.NotEmpty(t, config.Banks[0].ID)
	require.NotEmpty(t, config.Banks[0].URL)
	require.NotEmpty(t, config.Banks[0].AcpClient.TenantID)
	require.NotEmpty(t, config.Banks[0].AcpClient.ServerID)
	require.NotEmpty(t, config.Banks[0].AcpClient.ClientID)
	require.NotEmpty(t, config.Banks[0].AcpClient.CertFile)
	require.NotEmpty(t, config.Banks[0].AcpClient.KeyFile)
	require.NotEmpty(t, config.Banks[0].AcpClient.RootCA)
	require.Equal(t, 5*time.Second, config.Banks[0].AcpClient.Timeout)
}
