package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	envVars := map[string]string{
		"CLIENT_ID":     "client id",
		"CLIENT_SECRET": "client secret",
		"ISSUER_URL":    "https://localhost:8443/default/system",
		"CERT_FILE":     "cert file",
		"KEY_FILE":      "key file",
		"ROOT_CA":       "root ca",
		"FAILURE_URL":   "failure url",
		"LOG_LEVEL":     "warning",
		"PORT":          "8443",
		"TIMEOUT":       "30s",

		"OIDC_AUTH_METHOD":   "client_secret_basic",
		"OIDC_CLIENT_ID":     "oidc_client_id",
		"OIDC_CLIENT_SECRET": "oidc_client_secret",
		"OIDC_ISSUER_URL":    "oidc_issuer_url",
		"OIDC_REDIRECT_URL":  "oidc_redirect_url",
		"OIDC_CA_PATH":       "oidc_ca_path",
	}
	for k, v := range envVars {
		t.Setenv(k, v)
	}

	config, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, envVars["CLIENT_ID"], config.ClientID)
	assert.Equal(t, envVars["CLIENT_SECRET"], config.ClientSecret)
	assert.Equal(t, envVars["ISSUER_URL"], config.IssuerURL)
	assert.Equal(t, envVars["ROOT_CA"], config.RootCA)
	assert.Equal(t, envVars["CERT_FILE"], config.CertFile)
	assert.Equal(t, envVars["KEY_FILE"], config.KeyFile)
	assert.Equal(t, envVars["FAILURE_URL"], config.FailureURL)
	assert.Equal(t, envVars["LOG_LEVEL"], config.LogLevel)
	assert.Equal(t, 8443, config.Port)
	assert.Equal(t, time.Second*30, config.Timeout)

	assert.Equal(t, envVars["OIDC_AUTH_METHOD"], config.OIDC.AuthMethod)
	assert.Equal(t, envVars["OIDC_CLIENT_ID"], config.OIDC.ClientID)
	assert.Equal(t, envVars["OIDC_CLIENT_SECRET"], config.OIDC.ClientSecret)
	assert.Equal(t, envVars["OIDC_ISSUER_URL"], config.OIDC.IssuerURL)
	assert.Equal(t, envVars["OIDC_REDIRECT_URL"], config.OIDC.RedirectURL)
	assert.Equal(t, envVars["OIDC_CA_PATH"], config.OIDC.CAPath)

	// test default values
	assert.False(t, config.OIDC.PKCEEnabled)
	assert.Equal(t, []string{"openid"}, config.OIDC.Scopes)
}
