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
	"github.com/cloudentity/acp-client-go/client/oauth2"
)

type LoginURLBuilder interface {
	BuildLoginURL(string, acpclient.Client) (string, acpclient.CSRF, error)
}

type OBBRLoginURLBuilder struct {
	key jose.JSONWebKey
}

func NewOBBRLoginURLBuilder(c *gin.Context, client acpclient.Client) (LoginURLBuilder, error) {
	var (
		key jose.JSONWebKey
		err error
	)

	if key, err = getEncryptionKey(c, client); err != nil {
		return nil, err
	}

	return &OBBRLoginURLBuilder{key: key}, nil
}

func (o *OBBRLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	var (
		err error
	)

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

type OBUKLoginURLBuilder struct {
	key jose.JSONWebKey
}

func NewOBUKLoginURLBuilder() (LoginURLBuilder, error) {
	return &OBUKLoginURLBuilder{}, nil
}

func (o *OBUKLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	return client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:openbanking:psd2:sca"}),
		acpclient.WithPKCE())
}

func (s *Server) CreateConsentResponse(
	c *gin.Context, bankID BankID,
	consentID string,
	user User,
	client acpclient.Client,
	loginUrlBuilder LoginURLBuilder) {
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

	if loginURL, app.CSRF, err = loginUrlBuilder.BuildLoginURL(consentID, client); err != nil {
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

func getEncryptionKey(c *gin.Context, client acpclient.Client) (jose.JSONWebKey, error) {
	var (
		jwksResponse *oauth2.JwksOK
		encKey       jose.JSONWebKey
		b            []byte
		err          error
	)

	if jwksResponse, err = client.Oauth2.Jwks(
		oauth2.NewJwksParamsWithContext(c).
			WithTid(client.TenantID).
			WithAid(client.ServerID),
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
