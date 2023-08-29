package main

import (
	"context"
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

type OBBRLoginURLBuilder struct{}

func NewOBBRLoginURLBuilder() (LoginURLBuilder, error) {
	return &OBBRLoginURLBuilder{}, nil
}

func (o *OBBRLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	var (
		err error
		key jose.JSONWebKey
	)

	config := client.Config
	config.Scopes = append(config.Scopes, "consent:"+consentID)

	if key, err = o.getEncryptionKey(client); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to get encryption key")
	}

	if client, err = acpclient.New(config); err != nil {
		return "", acpclient.CSRF{}, errors.Wrapf(err, "failed to create new acp client")
	}

	return client.AuthorizeURL(
		acpclient.WithResponseType("code"),
		acpclient.WithOpenbankingIntentID(consentID, []string{"urn:brasil:openbanking:loa2"}),
		acpclient.WithRequestObjectEncryption(key),
		acpclient.WithPKCE(),
		acpclient.WithResponseMode("jwt"),
	)
}

func (o *OBBRLoginURLBuilder) getEncryptionKey(client acpclient.Client) (jose.JSONWebKey, error) {
	var (
		jwksResponse *oauth2.JwksOK
		encKey       jose.JSONWebKey
		b            []byte
		err          error
	)

	ctx, cancel := context.WithTimeout(context.Background(), client.Config.Timeout)
	defer cancel()

	if jwksResponse, err = client.Oauth2.Oauth2.Jwks(
		oauth2.NewJwksParamsWithContext(ctx),
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

type OBUKLoginURLBuilder struct{}

func NewOBUKLoginURLBuilder() (LoginURLBuilder, error) {
	return &OBUKLoginURLBuilder{}, nil
}

func (o *OBUKLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (string, acpclient.CSRF, error) {
	return client.AuthorizeURL(
		acpclient.WithResponseType("code"),
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
		acpclient.WithResponseType("code"),
		acpclient.WithResponseMode("jwt"),
		acpclient.WithPAR(client.Config.ClientID, arrangementID),
	)
}

func (s *Server) CreateConsentResponse(
	c *gin.Context,
	bankID BankID,
	user User,
	consentClient ConsentClient,
	client acpclient.Client,
	consentID string,
) {
	var (
		loginURL           string
		err                error
		encodedCookieValue string
		app                = AppStorage{
			BankID:   bankID,
			Sub:      user.Sub,
			IntentID: consentID,
		}
		data          = gin.H{}
		parRequestURI string
	)

	if consentClient.UsePAR() {
		if parRequestURI, app.CSRF, err = consentClient.DoPAR(c); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register PAR request: %+v", err))
			return
		}

		if loginURL, _, err = s.LoginURLBuilder.BuildLoginURL(parRequestURI, client); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
			return
		}
	} else {
		if loginURL, app.CSRF, err = s.LoginURLBuilder.BuildLoginURL(consentID, client); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
			return
		}
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

type FDXLoginURLBuilder struct{}

func NewFDXLoginURLBuilder(config Config) (LoginURLBuilder, error) {
	return &FDXLoginURLBuilder{}, nil
}

func (f *FDXLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (authorizeURL string, csrf acpclient.CSRF, err error) {
	return client.AuthorizeURL(
		acpclient.WithResponseType("code"),
		acpclient.WithPAR(client.Config.ClientID, consentID),
	)
}

type GenericLoginURLBuilder struct{}

func NewGenericLoginURLBuilder(config Config) (LoginURLBuilder, error) {
	return &GenericLoginURLBuilder{}, nil
}

func (f *GenericLoginURLBuilder) BuildLoginURL(consentID string, client acpclient.Client) (authorizeURL string, csrf acpclient.CSRF, err error) {
	return client.AuthorizeURL(
		acpclient.WithResponseType("code"),
		acpclient.WithPAR(client.Config.ClientID, consentID),
	)
}
