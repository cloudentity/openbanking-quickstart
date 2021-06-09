package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPemToPublicJwks(t *testing.T) {
	pem := "-----BEGIN CERTIFICATE-----\nMIIEVjCCAz6gAwIBAgIUHVV9LeNtt81MoMzosAABgejAp60wDQYJKoZIhvcNAQEL\nBQAwVzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcT\nB1NlYXR0bGUxFDASBgNVBAoTC0Nsb3VkZW50aXR5MQswCQYDVQQLEwJDQTAeFw0y\nMTAxMjAwOTQ4MDBaFw0yNjAxMTkwOTQ4MDBaMIGQMQswCQYDVQQGEwJVUzEUMBIG\nA1UECBMLV2FzaGluZ2h0b24xEDAOBgNVBAcTB1NlYXR0bGUxFDASBgNVBAoTC0Ns\nb3VkZW50aXR5MRYwFAYDVQQLEw1BdXRob3JpemF0aW9uMSswKQYDVQQDEyJjaWQx\nLmF1dGhvcml6YXRpb24uY2xvdWRlbnRpdHkuY29tMIIBIjANBgkqhkiG9w0BAQEF\nAAOCAQ8AMIIBCgKCAQEA4b/IX1bV29pw6/Ce8DdkoNx4dxJnDD9AyxmTG2z99cvl\nHG6BJaMF6l09ncGXGbv3dufDKrhftkwfbTBdpUEAeext/ugCmXTV06Fayva6Iq7x\nCNE8pA6hJT1y3Edsqq3IU8KVivYjYwd/vrSUfCe8pQRsR6K8rqnJ66ryn0yewkTE\nyCgPIv6pOMbgq1d5iX/2G9rZNhj74miN5y4fy0tsbI3q2RUOzt2d+htkoysqu3Xt\na6qPA3vEJ2FnQo3dhgw4XSCEvjz+HSGnsTC+XBv6j6jI9SD5jI2UYqnyDcYmRHPJ\nx2sQ/c8aLYHRdZxrxqIxUzulS6g0x74E2m0gBMKF5wIDAQABo4HfMIHcMA4GA1Ud\nDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0T\nAQH/BAIwADAdBgNVHQ4EFgQUt+G5I2aokd3gFFMty1rsIT8yMuwwHwYDVR0jBBgw\nFoAUylkHhGQUSf1DtT8uxH02jqlbo20wXQYDVR0RBFYwVIIJbG9jYWxob3N0giJj\naWQxLmF1dGhvcml6YXRpb24uY2xvdWRlbnRpdHkuY29tgh1hdXRob3JpemF0aW9u\nLmNsb3VkZW50aXR5LmNvbYcEfwAAATANBgkqhkiG9w0BAQsFAAOCAQEAjKXSkhqk\nL7tpjk5KbQe8OIvpsR0EjO3Za2KsvzUGDQzE+3tqL3jDvAhqWpICGGzv0kkroveb\n0kXc2Ltc5a+EQtm5l100N0kMB8f11/B1tcLFWQQnqEPB4RTkesGv70e1C3LbjmgJ\nbE62cLR8X2dXr20HxUMZAzmlDdRZS/80YnXSDgjcWxDiFVitFbFeUyYF/oh4RmO5\nHQKvMEd3XIO/hWKp0Jv2G4B5IwtWm2ZodaM6zwMgX0BBp1LUsSF5OYDjuJ2Tq8da\nEGmUj5Y/+CroeW7nFDgAwt+1x3M76uT8fo+rP+UHweR8TCSq28dY3+N8UYrD5YIB\nZD/0ro2b2KVvgg==\n-----END CERTIFICATE-----\n"
	res, err := PemToPublicJwks(pem)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	m := map[string]interface{}{}
	err = json.Unmarshal([]byte(res), &m)
	require.NoError(t, err)
}
