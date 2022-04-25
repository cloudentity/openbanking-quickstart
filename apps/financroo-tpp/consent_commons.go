package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	acpclient "github.com/cloudentity/acp-client-go"
	oauth2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

type LoginURLBuilder interface {
	BuildLoginURL(string, acpclient.Client) (string, acpclient.CSRF, error)
}

type OBBRLoginURLBuilder struct {
	key jose.JSONWebKey
}

func NewOBBRLoginURLBuilder(client acpclient.Client) (LoginURLBuilder, error) {
	var (
		key jose.JSONWebKey
		err error
	)

	if key, err = getEncryptionKey(client); err != nil {
		return nil, err
	}

	return &OBBRLoginURLBuilder{key: key}, nil
}

func (o *OBBRLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	var err error

	config := client.Config
	config.Scopes = append(config.Scopes, "consent:"+consentID)

	if client, err = acpclient.New(config); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create new acp client")
	}

	return client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:brasil:openbanking:loa2"}),
		acpclient.WithRequestObjectEncryption(o.key),
		acpclient.WithPKCE(),
	)
}

type OBUKLoginURLBuilder struct{}

func NewOBUKLoginURLBuilder() (LoginURLBuilder, error) {
	return &OBUKLoginURLBuilder{}, nil
}

func (o *OBUKLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	return client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:openbanking:psd2:sca"}),
		acpclient.WithPKCE())
}

type CDRLoginURLBuilder struct {
	signingKey interface{}
}

func NewCDRLoginURLBuilder(config Config) (LoginURLBuilder, error) {
	/*
		https://authorization.cloudentity.com:8443/default/cdr/oauth2/authorize?client_id=c9gpbd93kfoeqags5qqg&response_type=code id_token&scope=openid profile bank:accounts.basic:read bank:accounts.detail:read bank:transactions:read common:customer.basic:read introspect_tokens revoke_tokens&response_mode=form_post&request=eyJhbGciOiJQUzI1NiIsImtpZCI6IkI1NDhDOTE0QTAyNzg3QTNCNUYxNTU4M0M4RUIwMzBEOTRCQzI0MjQiLCJ4NXQiOiJ0VWpKRktBbmg2TzE4VldEeU9zRERaUzhKQ1EiLCJ0eXAiOiJKV1QifQ.eyJyZXNwb25zZV90eXBlIjoiY29kZSBpZF90b2tlbiIsImNsaWVudF9pZCI6ImM5Z3BiZDkza2ZvZXFhZ3M1cXFnIiwicmVkaXJlY3RfdXJpIjoiaHR0cHM6Ly9kYXRhcmVjaXBpZW50Lm1vY2s6OTAwMS9jb25zZW50L2NhbGxiYWNrIiwicmVzcG9uc2VfbW9kZSI6ImZvcm1fcG9zdCIsInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgYmFuazphY2NvdW50cy5iYXNpYzpyZWFkIGJhbms6YWNjb3VudHMuZGV0YWlsOnJlYWQgYmFuazp0cmFuc2FjdGlvbnM6cmVhZCBjb21tb246Y3VzdG9tZXIuYmFzaWM6cmVhZCBpbnRyb3NwZWN0X3Rva2VucyByZXZva2VfdG9rZW5zIiwic3RhdGUiOiI3YTFlNzdjMC1lNzIzLTQxZTItOTY0Mi02NTg5MTJjMjA2OWIiLCJub25jZSI6IjNlNTFlMjc0LTI4N2QtNDgyYS05M2JjLWI0OGZjNTc5YWRkYyIsImNsYWltcyI6eyJzaGFyaW5nX2R1cmF0aW9uIjpudWxsLCJjZHJfYXJyYW5nZW1lbnRfaWQiOm51bGwsInVzZXJpbmZvIjp7ImdpdmVuX25hbWUiOm51bGwsImZhbWlseV9uYW1lIjpudWxsfSwiaWRfdG9rZW4iOnsiYWNyIjp7ImVzc2VudGlhbCI6dHJ1ZSwidmFsdWVzIjpbInVybjpjZHMuYXU6Y2RyOjMiXX19fSwiZXhwIjoxNjUwNTYyODM1LCJpYXQiOjE2NTA1NjI1MzUsImlzcyI6ImM5Z3BiZDkza2ZvZXFhZ3M1cXFnIiwiYXVkIjoiaHR0cHM6Ly9hdXRob3JpemF0aW9uLmNsb3VkZW50aXR5LmNvbTo4NDQzL2RlZmF1bHQvY2RyIn0.F9-0fLpHSqn_nSAp3RdhqcfTRLSmiCr6p168b4qCcT9jHSDiTEpocbFx-QenWNpYAI-pCKAsa04U7jTElmOCQm2vEznf_vcyg-HbEOL_CikNg0K3BgsVXMgrPbPXE5tNf1Rrxh5ei8-aEnVjiZOTmerG-NSfWQFPY9Zqz11ZB9VAnupTVKMpmhRNvG5EC1guSvstOWefWSnuTzQq0jNQLQKkP9bobMwq32sp0ChejMfuBfLVxLuzTKP6DgBkBRNGRbPR40zBomLN5FuC2essCwmJYBKmJFKFmE6Gy55Uyhh1vh1a2vgvX2SeRLPLDUfIkC0uh_1CLOhgXyrNn5SwiQ
	*/

	var bs []byte
	var err error
	var signingKey interface{}

	if bs, err = os.ReadFile(config.KeyFile); err != nil {
		return nil, errors.Wrapf(err, "failed to read request object signing key")
	}

	block, _ := pem.Decode(bs)

	if signingKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return nil, errors.Wrapf(err, "failed to parse request object signing key")
	}

	return &CDRLoginURLBuilder{
		signingKey,
	}, nil
}

// TODO: AUT-5813

type ClaimRequests struct {
	Userinfo map[string]*ClaimRequest `json:"userinfo"`
	IDToken  map[string]*ClaimRequest `json:"id_token"`
}

type ClaimRequest struct {
	Essential bool     `json:"essential"`
	Value     string   `json:"value"`
	Values    []string `json:"values"`
}

func (o *CDRLoginURLBuilder) BuildLoginURL(arrangementID string, client acpclient.Client) (authorizeURL string, csrf acpclient.CSRF, err error) {
	var (
		parsed      *url.URL
		signedToken string
	)

	if authorizeURL, csrf, err = client.AuthorizeURL(acpclient.WithPKCE()); err != nil {
		return authorizeURL, csrf, err
	}

	if parsed, err = url.Parse(authorizeURL); err != nil {
		return authorizeURL, csrf, err
	}

	claims := jwt.MapClaims{
		"exp":          time.Now().Add(time.Minute).Unix(),
		"nonce":        csrf.Nonce,
		"state":        csrf.State,
		"nbf":          time.Now().Unix(),
		"redirect_uri": client.Config.RedirectURL.String(),
		"scope":        strings.Join(client.Config.Scopes, " "),
		"claims": ClaimRequests{
			IDToken: map[string]*ClaimRequest{
				"acr": {
					Essential: true,
					Value:     "urn:cds.au:cdr:2",
				},
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	if signedToken, err = token.SignedString(o.signingKey); err != nil {
		return authorizeURL, csrf, err
	}

	params := parsed.Query()
	params.Set("request", signedToken)
	parsed.RawQuery = params.Encode()

	return parsed.String(), csrf, nil
	// 	return fmt.Sprintf("%s?%s", c.Config.GetAuthorizeURL(), values.Encode()), csrf, nil
}

func (s *Server) CreateConsentResponse(
	c *gin.Context, bankID BankID,
	consentID string,
	user User,
	client acpclient.Client,
	loginURLBuilder LoginURLBuilder) {
	var (
		loginURL           string
		err                error
		encodedCookieValue string
		app                = AppStorage{
			BankID:   bankID,
			IntentID: consentID,
			Sub:      user.Sub,
		}
		data = gin.H{}
	)

	if loginURL, app.CSRF, err = loginURLBuilder.BuildLoginURL(consentID, client); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
		return
	}

	if _, err = url.Parse(loginURL); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to parse login url: %+v", err))
		return
	}

	// persist verifier and nonce in a secure encrypted cookie
	if encodedCookieValue, err = s.SecureCookie.Encode("app", app); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error while encoding cookie: %+v", err))
		return
	}

	c.SetCookie("app", encodedCookieValue, 0, "/", "", false, true)

	data["login_url"] = loginURL

	c.JSON(http.StatusOK, data)
}

func getEncryptionKey(client acpclient.Client) (jose.JSONWebKey, error) {
	var (
		jwksResponse *oauth2.JwksOK
		encKey       jose.JSONWebKey
		b            []byte
		err          error
	)

	ctx := gin.Context{}

	if jwksResponse, err = client.Oauth2.Oauth2.Jwks(
		oauth2.NewJwksParamsWithContext(&ctx),
	); err != nil {
		return encKey, errors.Wrapf(err, "failed to get jwks from acp server")
	}

	for _, key := range jwksResponse.Payload.Keys {
		if key.Use == "enc" {
			if b, err = json.Marshal(key); err != nil {
				return encKey, errors.Wrapf(err, "failed to marshal key")
			}

			if err = encKey.UnmarshalJSON(b); err != nil {
				return encKey, errors.Wrapf(err, "failed to unmarshal jwk")
			}

			break
		}
	}

	return encKey, nil
}
