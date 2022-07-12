package utils

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	acpclient "github.com/cloudentity/acp-client-go"
	oauth2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

type KeyUsage string

const (
	SIG KeyUsage = "sig"
	ENC KeyUsage = "enc"
)

func GetServerKey(client *acpclient.Client, keyUse KeyUsage) (jose.JSONWebKey, error) {
	var (
		jwksResponse *oauth2.JwksOK
		encKey       jose.JSONWebKey
		b            []byte
		err          error
	)

	if jwksResponse, err = client.Oauth2.Oauth2.Jwks(
		oauth2.NewJwksParamsWithContext(context.Background())); err != nil {
		return encKey, errors.Wrapf(err, "failed to get jwks from acp server")
	}

	for _, key := range jwksResponse.Payload.Keys {
		if key.Use == string(keyUse) {
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
