package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// OidcClient simple HTTP OIDC client.
type OidcClient struct {
	Config     OidcConfig
	HttpClient *http.Client // nolint
}

type OidcConfig struct {
	// ClientID is the ID of our client registered in ACP.
	ClientID string `env:"OIDC_CLIENT_ID"`
	// ClientSecret is the Secret of our client registered in ACP.
	ClientSecret string `env:"OIDC_CLIENT_SECRET"`
	// IssuerURL is the Issuer AuthURL
	IssuerURL string `env:"OIDC_ISSUER_URL"`
	// PKCEEnabled pkce on/off flag.
	PKCEEnabled bool `env:"OIDC_PKCE_ENABLED"`
	// RedirectURL holds information where to redirect the user after successful authentication.
	RedirectURL string `env:"OIDC_REDIRECT_URL"`
	// Scopes must be a subset of scopes assigned to our application in ACP.
	Scopes []string `env:"OIDC_SCOPES" envDefault:"openid"`

	CertPath string `env:"OIDC_CERT_PATH"`
	KeyPath  string `env:"OIDC_KEY_PATH"`
	CAPath   string `env:"OIDC_CA_PATH"`
}

// NewClient creates new instance of the Oidc client.
func (c OidcConfig) NewClient() (client OidcClient, err error) {
	var cert tls.Certificate

	// Set up the certificate HTTP client needs for TLS communication with a server.
	clientCACert, err := os.ReadFile(c.CAPath)
	if err != nil {
		return OidcClient{}, fmt.Errorf("could not open cert file %v: %w", c.CAPath, err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	// Assign a pool with certificates to the HTTP client.
	if cert, err = tls.LoadX509KeyPair(c.CertPath, c.KeyPath); err != nil {
		return OidcClient{}, fmt.Errorf("could not create acp client: %w", err)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			// Assign a pool with certificates to the HTTP client.
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs:      clientCertPool,
			},
		},
	}

	return OidcClient{HttpClient: httpClient, Config: c}, nil
}

// AuthorizeURL builds the URL where the client will redirect the user upon accessing /login endpoint.
// Challenge is a string used only when PKCE is enabled.
func (c OidcConfig) AuthorizeURL(challenge string, state string) string {
	var (
		buf bytes.Buffer

		queryParams = url.Values{
			"response_type": {"code"},
			"client_id":     {c.ClientID},
			"redirect_uri":  {c.RedirectURL},
			"scope":         {strings.Join(c.Scopes, " ")},
			"state":         {state},
		}
	)

	// When PKCE is on, we need to add a code challenge to the authorization request.
	if c.PKCEEnabled {
		queryParams.Add("code_challenge", challenge)
		queryParams.Add("code_challenge_method", "S256")
	}

	buf.WriteString(c.AuthURL())
	if strings.Contains(c.AuthURL(), "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}

	buf.WriteString(queryParams.Encode())
	return buf.String()
}

// AuthURL is an endpoint where ACP verifies the identity of the resource owner, and gain authorization grant.
func (c OidcConfig) AuthURL() string {
	return c.IssuerURL + "/oauth2/authorize"
}

// TokenURL holds information about the endpoint where we can exchange code for an access token.
func (c OidcConfig) TokenURL() string {
	return c.IssuerURL + "/oauth2/token"
}

func (c OidcConfig) GetRedirectURL(values url.Values) string {
	return c.RedirectURL + "?" + values.Encode()
}

func (c OidcClient) Exchange(code string, verifier string) (body []byte, err error) {
	values := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {c.Config.ClientID},
		"client_secret": {c.Config.ClientSecret},
		"redirect_uri":  {c.Config.RedirectURL},
	}

	if c.Config.PKCEEnabled {
		values.Add("code_verifier", verifier)
	}

	response, err := c.HttpClient.PostForm(c.Config.TokenURL(), values)
	if err != nil {
		return []byte{}, fmt.Errorf("error while obtaining token: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("ACP responded with status code: %v", response.Status)
	}

	if body, err = io.ReadAll(response.Body); err != nil {
		return []byte{}, fmt.Errorf("error during decoding exchange body: %w", err)
	}

	return body, nil
}
