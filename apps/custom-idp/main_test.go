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
		"ISSUER_URL":    "issuer url",
		"CERT_FILE":     "cert file",
		"KEY_FILE":      "key file",
		"ROOT_CA":       "root ca",
		"FAILURE_URL":   "failure url",
		"LOG_LEVEL":     "warning",
		"PORT":          "8443",
		"TIMEOUT":       "30s",
	}
	for k, v := range envVars {
		t.Setenv(k, v)
	}

	config, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, envVars["CLIENT_ID"], config.ClientID)
	assert.Equal(t, envVars["CLIENT_SECRET"], config.ClientSecret)
	assert.NotNil(t, config.IssuerURL)
	assert.Equal(t, envVars["ROOT_CA"], config.RootCA)
	assert.Equal(t, envVars["CERT_FILE"], config.CertFile)
	assert.Equal(t, envVars["KEY_FILE"], config.KeyFile)
	assert.Equal(t, envVars["FAILURE_URL"], config.FailureURL)
	assert.Equal(t, envVars["LOG_LEVEL"], config.LogLevel)
	assert.Equal(t, 8443, config.Port)
	assert.Equal(t, time.Second*30, config.Timeout)

	acpClient := config.AcpClientConfig()
	assert.Equal(t, config.ClientID, acpClient.ClientID)
	assert.Equal(t, config.ClientSecret, acpClient.ClientSecret)
	assert.Equal(t, config.IssuerURL, acpClient.IssuerURL)
	assert.Equal(t, config.CertFile, acpClient.CertFile)
	assert.Equal(t, config.KeyFile, acpClient.KeyFile)
	assert.Equal(t, config.RootCA, acpClient.RootCA)
	assert.Equal(t, config.Timeout, acpClient.Timeout)
}
