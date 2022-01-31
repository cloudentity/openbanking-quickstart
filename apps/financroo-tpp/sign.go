package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"strconv"
	"time"

	"github.com/go-jose/go-jose/v3"
)

const (
	iat = "http://openbanking.org.uk/iat"
	iss = "http://openbanking.org.uk/iss"
	tan = "http://openbanking.org.uk/tan"
	kid = "kid"
)

type Signer interface {
	Sign([]byte) (string, error)
}

type OBUKSigner struct {
	privateKey *rsa.PrivateKey
}

func NewOBUKSigner(privateKeyPath string) (Signer, error) {
	var (
		privateKey *rsa.PrivateKey
		err        error
	)

	if privateKey, err = getPrivateKey(privateKeyPath); err != nil {
		return nil, err
	}

	return &OBUKSigner{privateKey: privateKey}, nil
}

func (s *OBUKSigner) Sign(payload []byte) (string, error) {
	var (
		signer        jose.Signer
		jws           *jose.JSONWebSignature
		signerOptions = &jose.SignerOptions{}
		err           error
	)

	signerOptions.WithType("JOSE")
	signerOptions.WithContentType(jose.ContentType("application/json"))
	signerOptions.WithCritical(iat, iss, tan)
	signerOptions.WithHeader(jose.HeaderKey(kid), "167467200346518873990055631921812347975180003245")
	signerOptions.WithHeader(jose.HeaderKey(iat), strconv.FormatInt(time.Now().Unix(), 10))
	signerOptions.WithHeader(jose.HeaderKey(iss), "123456789123456789")
	signerOptions.WithHeader(jose.HeaderKey(tan), "openbanking.org.uk")

	if signer, err = jose.NewSigner(jose.SigningKey{Algorithm: jose.PS256, Key: s.privateKey}, signerOptions); err != nil {
		return "", err
	}

	if jws, err = signer.Sign(payload); err != nil {
		return "", err
	}

	return jws.DetachedCompactSerialize()
}

type OBBRSigner struct {
	privateKey *rsa.PrivateKey
}

func NewOBBRSigner(privateKeyPath string) (Signer, error) {
	var (
		privateKey *rsa.PrivateKey
		err        error
	)

	if privateKey, err = getPrivateKey(privateKeyPath); err != nil {
		return nil, err
	}

	return &OBBRSigner{privateKey: privateKey}, nil
}

func (s *OBBRSigner) Sign(payload []byte) (string, error) {
	var (
		signer        jose.Signer
		jws           *jose.JSONWebSignature
		signerOptions = &jose.SignerOptions{}
		err           error
	)

	signerOptions.WithType("JWT")

	if signer, err = jose.NewSigner(jose.SigningKey{Algorithm: jose.PS256, Key: s.privateKey}, signerOptions); err != nil {
		return "", err
	}

	if jws, err = signer.Sign(payload); err != nil {
		return "", err
	}

	return jws.CompactSerialize()
}

func getPrivateKey(keyFile string) (*rsa.PrivateKey, error) {
	var (
		block *pem.Block
		data  []byte
		k     *rsa.PrivateKey
		err   error
	)

	if data, err = os.ReadFile(keyFile); err != nil {
		return nil, err
	}

	block, _ = pem.Decode(data)
	if k, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	return k, nil
}
