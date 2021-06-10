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
			name:          "merge multiple templates",
			dirs:          []string{"./testdata/t1", "./testdata/t2"},
			variablesFile: "./testdata/variables.yaml",
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
		      ],
			  "jwks":{
                "keys":[
                   {
                     "alg":"RS256",
                     "e":"AQAB",
                     "kid":"167467200346518873990055631921812347975180003245",
                     "kty":"RSA",
                     "n":"4b_IX1bV29pw6_Ce8DdkoNx4dxJnDD9AyxmTG2z99cvlHG6BJaMF6l09ncGXGbv3dufDKrhftkwfbTBdpUEAeext_ugCmXTV06Fayva6Iq7xCNE8pA6hJT1y3Edsqq3IU8KVivYjYwd_vrSUfCe8pQRsR6K8rqnJ66ryn0yewkTEyCgPIv6pOMbgq1d5iX_2G9rZNhj74miN5y4fy0tsbI3q2RUOzt2d-htkoysqu3Xta6qPA3vEJ2FnQo3dhgw4XSCEvjz-HSGnsTC-XBv6j6jI9SD5jI2UYqnyDcYmRHPJx2sQ_c8aLYHRdZxrxqIxUzulS6g0x74E2m0gBMKF5w",
                     "use":"sig"
                   }
                ]
              }
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
