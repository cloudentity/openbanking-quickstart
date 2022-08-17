package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
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
		acpclient.WithResponseMode("jwt"),
	)
}

type OBUKLoginURLBuilder struct{}

func NewOBUKLoginURLBuilder() (LoginURLBuilder, error) {
	return &OBUKLoginURLBuilder{}, nil
}

func (o *OBUKLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	return client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:openbanking:psd2:sca"}),
		acpclient.WithPKCE(),
		acpclient.WithResponseMode("jwt"),
	)
}

type CDRLoginURLBuilder struct{}

func NewCDRLoginURLBuilder(config Config) (LoginURLBuilder, error) {
	return &CDRLoginURLBuilder{}, nil
}

func (o *CDRLoginURLBuilder) BuildLoginURL(arrangementID string, client acpclient.Client) (authorizeURL string, csrf acpclient.CSRF, err error) {
	return client.AuthorizeURL(
		acpclient.WithPKCE(),
		acpclient.WithOpenbankingACR([]string{"urn:cds.au:cdr:2"}),
		acpclient.WithResponseMode("jwt"),
	)
}

func (s *Server) CreateConsentResponse(
	c *gin.Context, bankID BankID,
	consentID string,
	user User,
	client acpclient.Client,
	loginURLBuilder LoginURLBuilder,
) {
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

type FDXLoginURLBuilder struct{}

func NewFDXLoginURLBuilder(config Config) (LoginURLBuilder, error) {
	return &CDRLoginURLBuilder{}, nil
}

func (f *FDXLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (authorizeURL string, csrf acpclient.CSRF, err error) {
	var (
		u string
	)

	if u, err = client.AuthorizeURLWithPAR(consentID); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create authorize url with par")
	}

	return u, acpclient.CSRF{}, nil
}
