package main

import (
	"bytes"
	"net/url"
	"strings"
	"time"

	acpclient "github.com/cloudentity/acp-client-go"
)

type OidcConfig struct {
	AuthStyle    string        `env:"OIDC_AUTH_STYLE"` // Enum: [client_secret_basic client_secret_post tls_client_auth ]
	ClientID     string        `env:"OIDC_CLIENT_ID"`
	ClientSecret string        `env:"OIDC_CLIENT_SECRET"`
	IssuerURL    string        `env:"OIDC_ISSUER_URL"`
	PKCEEnabled  bool          `env:"OIDC_PKCE_ENABLED"`
	RedirectURL  string        `env:"OIDC_REDIRECT_URL"`
	Scopes       []string      `env:"OIDC_SCOPES" envDefault:"openid"`
	Timeout      time.Duration `env:"OIDC_TIMEOUT" envDefault:"5s"`
	CAPath       string        `env:"OIDC_CA_PATH"`
}

func (c OidcConfig) Client() (acpclient.Client, error) {
	issuerURL, err := url.Parse(c.IssuerURL)
	if err != nil {
		return acpclient.Client{}, err
	}
	redirectURL, err := url.Parse(c.RedirectURL)
	if err != nil {
		return acpclient.Client{}, err
	}
	return acpclient.New(acpclient.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		IssuerURL:    issuerURL,
		RedirectURL:  redirectURL,
		Scopes:       c.Scopes,
		Timeout:      c.Timeout,
		RootCA:       c.CAPath,
	})
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

	authURL := c.IssuerURL + "/oauth2/authorize"
	buf.WriteString(authURL)
	if strings.Contains(authURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}

	buf.WriteString(queryParams.Encode())
	return buf.String()
}
