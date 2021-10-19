package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/go-jose/go-jose/v3"
	"io/ioutil"
)

// s.Config.KeyFile
func (s Server) JWSSignature(payload []byte) (string, error) {
	var (
		privateKey       *rsa.PrivateKey
		serializedJWS    string
		signer           jose.Signer
		jsonWebSignature *jose.JSONWebSignature
		err              error
	)

	if privateKey, err = getPrivateKey(s.Config.KeyFile); err != nil {
		return "", err
	}

	if signer, err = jose.NewSigner(jose.SigningKey{Algorithm: jose.PS256, Key: privateKey}, nil); err != nil {
		return "", err
	}

	if jsonWebSignature, err = signer.Sign(payload); err != nil {
		return "", err
	}

	if serializedJWS, err = jsonWebSignature.CompactSerialize(); err != nil {
		return "", err
	}

	return serializedJWS, nil
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
