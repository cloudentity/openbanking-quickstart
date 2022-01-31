package main

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"time"

	acpclient "github.com/cloudentity/acp-client-go"
)

type OidcConfig struct {
	AuthMethod   string        `env:"OIDC_AUTH_METHOD" envDefault:"client_secret_post"`
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
	authMethod, err := c.AuthenticationMethod()
	if err != nil {
		return acpclient.Client{}, err
	}
	issuerURL, err := url.Parse(c.IssuerURL)
	if err != nil {
		return acpclient.Client{}, err
	}
	redirectURL, err := url.Parse(c.RedirectURL)
	if err != nil {
		return acpclient.Client{}, err
	}
	return acpclient.New(acpclient.Config{
		AuthMethod:   authMethod,
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		IssuerURL:    issuerURL,
		RedirectURL:  redirectURL,
		Scopes:       c.Scopes,
		Timeout:      c.Timeout,
		RootCA:       c.CAPath,
	})
}

func (c OidcConfig) AuthenticationMethod() (method acpclient.AuthMethod, err error) {
	switch c.AuthMethod {
	case "client_secret_basic":
		method = acpclient.ClientSecretBasicAuthnMethod
	case "client_secret_post":
		method = acpclient.ClientSecretPostAuthnMethod
	case "client_secret_jwt":
		method = acpclient.ClientSecretJwtAuthnMethod
	case "private_key_jwt":
		method = acpclient.PrivateKeyJwtAuthnMethod
	case "self_signed_tls_client_auth":
		method = acpclient.SelfSignedTLSAuthnMethod
	case "tls_client_auth":
		method = acpclient.TLSClientAuthnMethod
	case "none":
		method = acpclient.NoneAuthnMethod
	default:
		err = fmt.Errorf("unsupported OIDC AuthMethod %q", c.AuthMethod)
	}
	return method, err
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
