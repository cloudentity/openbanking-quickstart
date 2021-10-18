package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/pkg/errors"
)
type SignatureHeader struct {
	Type        string          `json:"typ,omitempty"`
	Kid         string          `json:"kid,omitempty"`
	Alg         string          `json:"alg,omitempty"`
	Ctype       string          `json:"cty,omitempty"`
	Issuer      string          `json:"http://openbanking.org.uk/iss,omitempty"`
	IssuedAt    decimal.Decimal `json:"http://openbanking.org.uk/iat,omitempty"`
	TrustAnchor string          `json:"http://openbanking.org.uk/tan,omitempty"`
	B64         *bool           `json:"b64,omitempty"`
	Critical    []string        `json:"crit,omitempty"`
}
func Sign(payload []byte, key jose.JSONWebKey, signingOpts *jose.SignerOptions, detached bool) (string, error) {
	var (
		err        error
		signer     jose.Signer
		jws        *jose.JSONWebSignature
		serialized string
	)

	// algorithm must be ps256 as required by the 3.1.6 specification
	if signer, err = jose.NewSigner(
		jose.SigningKey{Algorithm: jose.PS256, Key: key.Key},
		signingOpts,
	); err != nil {
		return "", errors.Wrapf(err, fmt.Sprintf("failed to create signer with key %+v", key))
	}

	if jws, err = signer.Sign(payload); err != nil {
		return serialized, errors.Wrapf(err, "failed to sign data")
	}

	if detached {
		return jws.DetachedCompactSerialize()
	}

	return jws.CompactSerialize()
}

// TODO: this will need to take in a configuration for iss and tan
func GetOpenbankingUKSignerOptions(contentType string, key jose.JSONWebKey) *jose.SignerOptions {
	signerOptions := &jose.SignerOptions{}
	signerOptions.WithType("JOSE")
	signerOptions.WithContentType(jose.ContentType(contentType))
	signerOptions.WithCritical(string(iat), string(iss), string(tan))

	signerOptions.WithHeader(jose.HeaderKey(kid), key.KeyID)
	signerOptions.WithHeader(jose.HeaderKey(iat), strconv.FormatInt(time.Now().Unix(), 10))
	signerOptions.WithHeader(jose.HeaderKey(iss), "123456789123456789")
	signerOptions.WithHeader(jose.HeaderKey(tan), "openbanking.org.uk")

	return signerOptions
}

func GetOpenbankingBRSignerOptions(contentType string, key jose.JSONWebKey) *jose.SignerOptions {
	signerOptions := &jose.SignerOptions{}
	signerOptions.WithType("JWT")
	signerOptions.WithContentType(jose.ContentType(cty))
	signerOptions.WithHeader("kid", key.KeyID)
	return signerOptions
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
