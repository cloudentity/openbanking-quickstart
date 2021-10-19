package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
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
		return nil, errors.New(fmt.Sprintf("failed read certificate %v", err))
	}

	block, _ = pem.Decode(data)
	k, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse request object signing key")
	}

	return k, nil
}
