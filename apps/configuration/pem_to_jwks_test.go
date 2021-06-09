package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gopkg.in/square/go-jose.v2"
)

func TestPemToJWKS(t *testing.T) {
	bs, err := ioutil.ReadFile("../../data/tpp_cert.pem")
	require.NoError(t, err)
	require.NotNil(t, bs)

	cert, err := ParseCertificate(bs)
	require.NoError(t, err)

	key, err := ToPublicJWKs(cert)
	require.NoError(t, err)

	jsonBytes, err := json.MarshalIndent(&key, "", "    ")
	require.NoError(t, err)

	t.Logf("%s", jsonBytes)
	t.Fail()
}

func ToPublicJWKs(c *x509.Certificate) (jose.JSONWebKey, error) {
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
		return key, errors.New("not supported public key type %v (must be rsa or ecdsa)")
	}

	return key, nil
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
