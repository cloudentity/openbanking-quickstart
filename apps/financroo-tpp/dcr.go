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

	"github.com/go-jose/go-jose/v4"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"

	"github.com/cloudentity/openbanking-quickstart/shared"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	"github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type DCRClient struct {
	Name                    string
	URI                     string
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

type GenericDCRHandler struct {
	client acpclient.Client
}

func (d *GenericDCRHandler) Register(ctx context.Context, client DCRClient) (DCRClientCreated, error) {
	var (
		resp    *oauth2.DynamicClientRegistrationCreated
		created DCRClientCreated
		err     error
	)

	if resp, err = d.client.Oauth2.Oauth2.DynamicClientRegistration(
		oauth2.NewDynamicClientRegistrationParamsWithContext(ctx).
			WithClient(&models.DynamicClientRegistrationRequest{
				ClientName:              client.Name,
				ClientURI:               client.URI,
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

var _ DCRHandler = &GenericDCRHandler{}

func NewGenericDCRHandler(cfg Config, bankConfig BankConfig) (DCRHandler, error) {
	var (
		c         = GenericDCRHandler{}
		issuerURL *url.URL
		err       error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", bankConfig.ACPInternalURL, bankConfig.Tenant, bankConfig.Server)); err != nil {
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

func RegisterClient(ctx context.Context, config Config, bankConfig BankConfig) (DCRClientCreated, error) {
	var (
		dcr  DCRHandler
		err  error
		resp DCRClientCreated
		cert *x509.Certificate
		jwks models.ClientJWKs
	)

	if config.Spec != GENERIC {
		return resp, errors.New("DCR can be enabled only for Generic spec")
	}

	if dcr, err = NewGenericDCRHandler(config, bankConfig); err != nil {
		return resp, errors.Wrapf(err, "failed to init DCR")
	}

	if cert, err = loadCertificate(config); err != nil {
		return resp, errors.Wrapf(err, "failed to load certificate")
	}

	if jwks, err = toPublicJWKs(cert); err != nil {
		return resp, errors.Wrapf(err, "failed to convert cert to public jwks")
	}

	if resp, err = dcr.Register(
		ctx,
		DCRClient{
			Name: "Financroo TPP",
			URI:  config.UIURL,
			GrantTypes: []string{
				"authorization_code",
				"refresh_token",
			},
			RedirectURLs: []string{
				config.UIURL + "/api/callback",
			},
			Scopes:                  config.ClientScopes,
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

func loadCertificate(config Config) (*x509.Certificate, error) {
	var (
		data []byte
		pm   *pem.Block
		cert *x509.Certificate
		err  error
	)

	if data, err = os.ReadFile(config.CertFile); err != nil {
		return nil, errors.Wrapf(err, "failed to read cert file")
	}

	pm, _ = pem.Decode(data)

	if pm != nil {
		data = pm.Bytes
	}

	if cert, err = x509.ParseCertificate(data); err != nil {
		return nil, errors.Wrapf(err, "failed to parse x509 certificate")
	}

	return cert, nil
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

var dcrBucket = []byte("dcr")

type DCRClientIDStorage struct {
	shared.DB
}

func NewDCRClientIDStorage(db shared.DB) (storage DCRClientIDStorage, err error) {
	if err = shared.CreateBucket(db, dcrBucket); err != nil {
		return storage, err
	}

	storage.DB = db

	return storage, nil
}

func (c *DCRClientIDStorage) Get(id BankID) (string, bool, error) {
	var (
		clientID string
		exists   bool
		err      error
	)

	if err = c.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(dcrBucket).Get([]byte(id))

		if value != nil {
			clientID = string(value)
			exists = true
		}

		return nil
	}); err != nil {
		return "", false, errors.Wrapf(err, "failed to fetch client_id from db")
	}

	return clientID, exists, nil
}

func (c *DCRClientIDStorage) Set(id BankID, clientID string) error {
	return c.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dcrBucket).Put([]byte(id), []byte(clientID))
	})
}
