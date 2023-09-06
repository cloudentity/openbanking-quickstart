package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

type ExternalConsentStorage struct {
	URL *url.URL
	HC  *http.Client
}

func NewExternalConsentStorage(config ExternalConsentStorageConfig) (ExternalConsentStorage, error) {
	var (
		pool  *x509.CertPool
		data  []byte
		cert  tls.Certificate
		certs []tls.Certificate
		err   error
	)

	if pool, err = x509.SystemCertPool(); err != nil {
		return ExternalConsentStorage{}, errors.Wrap(err, "failed to load system cert pool")
	}

	if config.RootCA != "" {
		if data, err = os.ReadFile(config.RootCA); err != nil {
			return ExternalConsentStorage{}, errors.Wrap(err, "failed to read root CA")
		}
		pool.AppendCertsFromPEM(data)
	}

	if config.CertFile != "" && config.KeyFile != "" {
		if cert, err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile); err != nil {
			return ExternalConsentStorage{}, errors.Wrap(err, "failed to load tls cert")
		}

		certs = append(certs, cert)
	}

	hc := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				MinVersion:   tls.VersionTLS12,
				Certificates: certs,
			},
		},
	}

	return ExternalConsentStorage{
		URL: config.URL,
		HC:  hc,
	}, nil
}

var _ ConsentStorage = &ExternalConsentStorage{}

type CreateConsentResponse struct {
	ID string `json:"id"`
}

func (e *ExternalConsentStorage) Store(ctx context.Context, sub Subject, data Data) (ConsentID, error) {
	var (
		bs   []byte
		req  *http.Request
		res  *http.Response
		resp CreateConsentResponse
		err  error
	)

	if bs, err = json.Marshal(&data); err != nil {
		return "", errors.Wrap(err, "failed to marshal data")
	}

	u := fmt.Sprintf("%s/consents", e.URL.String())

	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewReader(bs)); err != nil {
		return "", errors.Wrap(err, "failed to create request")
	}

	if res, err = e.HC.Do(req); err != nil {
		return "", errors.Wrapf(err, "failed to send request to: %s", u)
	}
	defer res.Body.Close()

	if bs, err = io.ReadAll(res.Body); err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}

	if res.StatusCode != http.StatusCreated {
		return "", errors.Wrapf(err, "unexpected status code: %d, body: %s", res.StatusCode, string(bs))
	}

	if err = json.Unmarshal(bs, &resp); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal response")
	}

	return ConsentID(resp.ID), nil
}
