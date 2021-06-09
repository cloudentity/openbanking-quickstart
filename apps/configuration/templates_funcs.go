package main

import (
	"crypto/x509"
	"encoding/json"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
	"gopkg.in/square/go-jose.v2"
)

func CustomTemplateFuncs() template.FuncMap {
	extra := template.FuncMap{
		"pem_to_public_jwks": MustPemToPublicJwks,
	}

	// merge with sprig
	f := sprig.TxtFuncMap()
	for k, v := range extra {
		f[k] = v
	}

	return f
}

func PemToPublicJwks(v string) (string, error) {
	var (
		c    *x509.Certificate
		jwks jose.JSONWebKeySet
		bs   []byte
		err  error
	)

	if c, err = ParseCertificate([]byte(v)); err != nil {
		return "", err
	}

	if jwks, err = ToPublicJWKs(c); err != nil {
		return "", err
	}

	if bs, err = json.Marshal(&jwks); err != nil {
		return "", err
	}

	return string(bs), nil
}

func MustPemToPublicJwks(v string) string {
	res, err := PemToPublicJwks(v)
	if err != nil {
		return ""
	}

	return res
}
