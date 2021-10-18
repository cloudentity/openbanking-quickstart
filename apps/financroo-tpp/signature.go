package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type OpenbankingHeaderName string

const (
	typ  OpenbankingHeaderName = "typ"
	cty  OpenbankingHeaderName = "cty"
	alg  OpenbankingHeaderName = "alg"
	kid  OpenbankingHeaderName = "kid"
	iat  OpenbankingHeaderName = "http://openbanking.org.uk/iat"
	iss  OpenbankingHeaderName = "http://openbanking.org.uk/iss"
	tan  OpenbankingHeaderName = "http://openbanking.org.uk/tan"
	crit OpenbankingHeaderName = "crit"
)

type ContextKey string

const (
	ContentTypeKey ContextKey = "Content-Type"
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

func NewSignatureHeader(signature string) (SignatureHeader, error) {
	var (
		err             error
		decoded         []byte
		signatureHeader SignatureHeader
	)

	if decoded, err = validateAndDecodeSignature(signature); err != nil {
		return SignatureHeader{}, err
	}

	if err = json.Unmarshal(decoded, &signatureHeader); err != nil {
		return signatureHeader, errors.New("failed to serialize header to json")
	}

	return signatureHeader, nil
}

func validateAndDecodeSignature(signature string) (decoded []byte, err error) {
	parts := strings.Split(signature, ".")

	if signature == "" {
		return decoded, errors.New("signature missing")
	}

	if len(parts) != 3 {
		return decoded, errors.New("signature must contain three parts")
	}

	if decoded, err = base64.RawURLEncoding.DecodeString(parts[0]); err != nil {
		return decoded, errors.New("failed to base64 decode header")
	}

	return decoded, nil
}
