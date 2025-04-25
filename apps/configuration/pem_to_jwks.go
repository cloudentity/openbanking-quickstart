package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/go-jose/go-jose/v4"
	"github.com/pkg/errors"
)

func ToPublicJWKs(c *x509.Certificate) (jose.JSONWebKeySet, error) {
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
		return jose.JSONWebKeySet{}, errors.New("not supported public key type %v (must be rsa or ecdsa)")
	}

	return jose.JSONWebKeySet{Keys: []jose.JSONWebKey{key}}, nil
}

func ParseCertificate(data []byte) (cert *x509.Certificate, err error) {
	var block *pem.Block

	if block, data = pem.Decode(data); block == nil {
		data = bytes.TrimSpace(data)
		if len(data) > 0 {
			return nil, fmt.Errorf("certificate contains invalid data: %q", string(data))
		}
	}

	if cert, err = x509.ParseCertificate(block.Bytes); err != nil {
		return nil, errors.Wrapf(err, "failed to parse certificate")
	}

	return cert, nil
}
