package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplates(t *testing.T) {
	testCases := []struct {
		name               string
		dirs               []string
		variablesFile      string
		assertResponse     func(*testing.T, []byte)
		expectedJSONOutput []byte
	}{
		{
			name: "merge templates",
			dirs: []string{"./testdata/t1"},
			assertResponse: func(tt *testing.T, bs []byte) {
				var m map[string]interface{}

				err := json.Unmarshal(bs, &m)
				require.NoError(tt, err)

				require.Equal(tt, 2, len(m["servers"].([]interface{})))
				require.Equal(tt, 1, len(m["clients"].([]interface{})))
			},
		},
		{
			name: "merge multiple templates",
			dirs: []string{"./testdata/t1", "./testdata/t2"},
			assertResponse: func(tt *testing.T, bs []byte) {
				var m map[string]interface{}

				err := json.Unmarshal(bs, &m)
				require.NoError(tt, err)

				require.Equal(tt, 2, len(m["servers"].([]interface{})))
				require.Equal(tt, 2, len(m["clients"].([]interface{})))
			},
		},
		{
			name:          "merge templates using variables file",
			dirs:          []string{"./testdata/t2"},
			variablesFile: "./testdata/variables.yaml",
			assertResponse: func(tt *testing.T, bs []byte) {
				require.JSONEq(tt, string([]byte(`{
		  "clients": [
		    {
		      "tenant_id": "default",
		      "authorization_server_id": "test",
		      "client_id": "cid",
		      "client_name": "Test App",
		      "client_secret": "secret",
		      "redirect_uris": [
		        "https://localhost:8091/callback"
		      ]
		    }
		  ]
		}`)), string(bs))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			var variablesFile *string
			if tc.variablesFile != "" {
				variablesFile = &tc.variablesFile
			}

			templates, err := LoadTemplates(tc.dirs, variablesFile)
			require.NoError(tt, err)

			yamlFile, err := templates.Merge()
			require.NoError(tt, err)

			bs, err := yamlFile.ToJSON()

			require.NoError(tt, err)

			if tc.assertResponse != nil {
				tc.assertResponse(tt, bs)
			}
		})
	}
}
