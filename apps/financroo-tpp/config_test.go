package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	for k, v := range map[string]string{
		"ACP_URL":           "https://localhost:8443",
		"ACP_MTLS_URL":      "https://acp:8443",
		"APP_HOST":          "localhost",
		"UI_URL":            "https://localhost:8091",
		"CERT_FILE":         "cert.pem",
		"KEY_FILE":          "key.pem",
		"TENANT":            "default",
		"SPEC":              "obbr",
		"BANK_URL":          "http://bank-br:8070",
		"BANKS_CONFIG_FILE": "banks.json",
	} {
		os.Setenv(k, v)
	}

	config, err := LoadConfig()
	require.NoError(t, err)

	require.Equal(t, 8091, config.Port)
	require.Equal(t, "/app/data/my.db", config.DBFile)
	require.Equal(t, "https://localhost:8443", config.ACPURL)
	require.Equal(t, "https://acp:8443", config.ACPInternalURL)
	require.Equal(t, "localhost", config.AppHost)
	require.Equal(t, "https://localhost:8091", config.UIURL)
	require.Equal(t, "cert.pem", config.CertFile)
	require.Equal(t, "key.pem", config.KeyFile)
	require.Equal(t, Spec("obbr"), config.Spec)
	require.Equal(t, "http://bank-br:8070", config.BankURL)

	require.Equal(t, 2, len(config.Banks))
	require.Equal(t, []BankID{BankID("gobank"), BankID("hyperscalebank")}, config.Banks.GetIDs())
}

func TestLoadBanksConfig(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		config      Config
		jsonContent string
		expectError bool
		assert      func(*testing.T, BanksConfig)
		expectedLen int
	}{
		{
			name:        "valid banks config file",
			config:      Config{},
			jsonContent: `{"banks": [{"id": "golang","url": "http://bank:7080", "acp_url": "https://acp:8443", "acp_internal_url": "https://acp:8443", "tenant": "default", "server": "generic"}]}`,
			assert: func(tt *testing.T, c BanksConfig) {
				require.Equal(tt, 1, len(c))
			},
		},
		{
			name: "config fallback",
			config: Config{
				ACPURL:          "https://localhost:8443",
				BanksConfigFile: "",
			},
			assert: func(tt *testing.T, c BanksConfig) {
				require.Equal(tt, 1, len(c))
				require.Equal(tt, BankID("gobank"), c[0].ID)
				require.Equal(tt, "https://localhost:8443", c[0].ACPURL)
			},
		},
		{
			name:        "invalid banks config file",
			jsonContent: `{"banks": [{}]}`,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.jsonContent != "" {
				configFile, err := ioutil.TempFile("/tmp", "testbankconfig_*.json")
				require.NoError(t, err)
				defer os.Remove(configFile.Name())

				_, err = configFile.WriteString(tc.jsonContent)
				require.NoError(t, err)
				tc.config.BanksConfigFile = configFile.Name()
			}

			banksConfig, err := LoadBanksConfig(tc.config)

			if tc.expectError {
				require.Error(t, err)
				require.Equal(t, tc.expectedLen, len(banksConfig))
			} else {
				require.NoError(t, err)
				if tc.assert != nil {
					tc.assert(t, banksConfig)
				}
			}
		})
	}
}
