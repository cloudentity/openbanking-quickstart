package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
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

func (s Server) JWSSignature(payload []byte) (string, error) {
	var (
		privateKey  *rsa.PrivateKey
		detachedJWS string
		signer      jose.Signer
		jws         *jose.JSONWebSignature
		err         error
	)

	if privateKey, err = getPrivateKey(s.Config.KeyFile); err != nil {
		return "", err
	}

	if signer, err = jose.NewSigner(jose.SigningKey{Algorithm: jose.PS256, Key: privateKey}, getSignerOptions()); err != nil {
		return "", err
	}

	if jws, err = signer.Sign(payload); err != nil {
		return "", err
	}

	if detachedJWS, err = jws.DetachedCompactSerialize(); err != nil {
		return "", err
	}

	return detachedJWS, nil
}

func getPrivateKey(keyFile string) (*rsa.PrivateKey, error) {
	var (
		block *pem.Block
		data  []byte
		k     *rsa.PrivateKey
		err   error
	)

	if data, err = ioutil.ReadFile(keyFile); err != nil {
		return nil, err
	}

	block, _ = pem.Decode(data)
	if k, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	return k, nil
}

func getSignerOptions() *jose.SignerOptions {
	signerOptions := &jose.SignerOptions{}
	signerOptions.WithType("JOSE")
	signerOptions.WithContentType(jose.ContentType("application/json"))
	signerOptions.WithCritical(iat, iss, tan)

	signerOptions.WithHeader(jose.HeaderKey(kid), "167467200346518873990055631921812347975180003245")
	signerOptions.WithHeader(jose.HeaderKey(iat), strconv.FormatInt(time.Now().Unix(), 10))
	signerOptions.WithHeader(jose.HeaderKey(iss), "123456789123456789")
	signerOptions.WithHeader(jose.HeaderKey(tan), "openbanking.org.uk")

	return signerOptions
}
