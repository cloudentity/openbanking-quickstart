package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	"github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type DCRClient struct {
	Name                    string
	GrantTypes              []string
	RedirectURLs            []string
	Scopes                  []string
	TokenEndpointAuthMethod string
	TLSClientAuthSubjectDN  string
	ResponseTypes           []string
	Jwks                    *models.ClientJWKs
}

type DCRClientCreated struct {
	ClientID string
	Scopes   []string
}

type DCRHandler interface {
	Register(ctx context.Context, client DCRClient) (DCRClientCreated, error)
}

type OAuth2DCRHandler struct {
	client acpclient.Client
}

func (d *OAuth2DCRHandler) Register(ctx context.Context, client DCRClient) (DCRClientCreated, error) {
	var (
		resp    *oauth2.DynamicClientRegistrationCreated
		created DCRClientCreated
		err     error
	)

	if resp, err = d.client.Oauth2.Oauth2.DynamicClientRegistration(
		oauth2.NewDynamicClientRegistrationParamsWithContext(ctx).
			WithClient(&models.DynamicClientRegistrationRequest{
				ClientName:              client.Name,
				GrantTypes:              client.GrantTypes,
				RedirectUris:            client.RedirectURLs,
				ResponseTypes:           client.ResponseTypes,
				Scopes:                  client.Scopes,
				TLSClientAuthSubjectDn:  client.TLSClientAuthSubjectDN,
				TokenEndpointAuthMethod: client.TokenEndpointAuthMethod,
				Jwks:                    client.Jwks,
			}), nil); err != nil {
		return created, errors.Wrapf(err, "failed to register client: %+v", client)
	}

	created = DCRClientCreated{
		ClientID: resp.Payload.ClientID,
		Scopes:   resp.Payload.Scopes,
	}

	return created, nil
}

var _ DCRHandler = &OAuth2DCRHandler{}

func NewOAuth2DCR(cfg Config) (DCRHandler, error) {
	var (
		c         = OAuth2DCRHandler{}
		issuerURL *url.URL
		err       error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", cfg.ACPInternalURL, cfg.Tenant, cfg.ServerID)); err != nil {
		return nil, err
	}

	config := acpclient.Config{
		ClientID:                   "dummy", // client id is required in acpclient lib
		IssuerURL:                  issuerURL,
		Timeout:                    time.Second * 5,
		CertFile:                   cfg.CertFile,
		KeyFile:                    cfg.KeyFile,
		RootCA:                     cfg.RootCA,
		SkipClientCredentialsAuthn: true,
	}

	if c.client, err = acpclient.New(config); err != nil {
		return nil, err
	}

	return &c, nil
}

func RegisterClient(ctx context.Context, config Config) (DCRClientCreated, error) {
	var (
		dcr  DCRHandler
		err  error
		resp DCRClientCreated
		data []byte
		cert *x509.Certificate
		jwks models.ClientJWKs
		pm   *pem.Block
	)

	if config.Spec != GENERIC {
		return resp, errors.New("DCR can be enabled only for Generic spec")
	}

	if dcr, err = NewOAuth2DCR(config); err != nil {
		return resp, errors.Wrapf(err, "failed to init DCR")
	}

	if data, err = os.ReadFile(config.CertFile); err != nil {
		return resp, errors.Wrapf(err, "failed to read cert file")
	}

	pm, _ = pem.Decode(data)

	if pm != nil {
		data = pm.Bytes
	}

	if cert, err = x509.ParseCertificate(data); err != nil {
		return resp, errors.Wrapf(err, "failed to parse x509 certificate")
	}

	if jwks, err = toPublicJWKs(cert); err != nil {
		return resp, errors.Wrapf(err, "failed to convert cert to public jwks")
	}

	if resp, err = dcr.Register(
		ctx,
		DCRClient{
			Name: "Financroo TPP",
			GrantTypes: []string{
				"authorization_code",
				"refresh_token",
			},
			RedirectURLs: []string{
				config.UIURL + "/api/callback",
			},
			Scopes: []string{
				"openid",
				"email",
				"offline_access",
				"sample",
			},
			TokenEndpointAuthMethod: "tls_client_auth",
			TLSClientAuthSubjectDN:  cert.Subject.ToRDNSequence().String(),
			ResponseTypes: []string{
				"code",
			},
			Jwks: &jwks,
		},
	); err != nil {
		return resp, errors.Wrapf(err, "failed to register client")
	}

	return resp, nil
}

func toPublicJWKs(c *x509.Certificate) (models.ClientJWKs, error) {
	var (
		res models.ClientJWK
		bs  []byte
		err error
	)

	key := jose.JSONWebKey{
		Key:   c.PublicKey,
		Use:   "sig",
		KeyID: c.SerialNumber.String(),
	}

	switch c.PublicKey.(type) {
	case *rsa.PublicKey:
		key.Algorithm = "RS256"
	case *ecdsa.PublicKey:
		key.Algorithm = "ES256"
	default:
		return models.ClientJWKs{}, errors.New("not supported public key type %v (must be rsa or ecdsa)")
	}

	if bs, err = key.MarshalJSON(); err != nil {
		return models.ClientJWKs{}, errors.Wrapf(err, "failed to marshal jwk")
	}

	if err = json.Unmarshal(bs, &res); err != nil {
		return models.ClientJWKs{}, errors.Wrapf(err, "failed to unmarshal jwk")
	}

	return models.ClientJWKs{Keys: []*models.ClientJWK{&res}}, nil
}
