package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomOTPHandler struct {
	config OtpConfig
	client *http.Client
	Repo   *OTPRepo
}

func NewCustomOTPHandler(config Config, repo *OTPRepo) (OTPHandler, error) {
	var (
		pool = x509.NewCertPool()
		bts  []byte
		err  error
	)

	if bts, err = ioutil.ReadFile(config.RootCA); err != nil {
		logrus.Errorf("failed to read root ca from %s", config.RootCA)
		return nil, err
	}

	if !pool.AppendCertsFromPEM(bts) {
		logrus.Error("failed to append cert from pem")
		return nil, err
	}

	return &CustomOTPHandler{
		config: config.Otp,
		client: &http.Client{
			Timeout: config.Otp.Timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:    pool,
					MinVersion: tls.VersionTLS12,
				},
			},
		},
		Repo: repo,
	}, nil
}

type CustomOtpRequest struct {
	Sub string `json:"sub"`
}

func (c *CustomOTPHandler) Send(_ LoginRequest, _ MFAConsentProvider, to string, _ MFAData) error {
	var (
		bts []byte
		err error
		req *http.Request
		res *http.Response
	)

	// we are not sending any otp request for totp
	if c.isOfflineOTPType() {
		return nil
	}

	if bts, err = json.Marshal(CustomOtpRequest{Sub: to}); err != nil {
		return err
	}

	if req, err = http.NewRequest(http.MethodPost, c.config.RequestURL, bytes.NewBuffer(bts)); err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.config.AuthHeader)

	if res, err = c.client.Do(req); err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return errors.New("otp request not accepted")
	}

	return nil
}

type CustomOTPVerify struct {
	Otp string `json:"otp"`
	Sub string `json:"sub"`
}

type CustomOTPVerifyResp struct {
	Ok bool `json:"ok"`
}

func (c *CustomOTPHandler) Verify(r LoginRequest, login string, otp string) (bool, error) {
	var (
		bts []byte
		err error
		req *http.Request
		res *http.Response
		ver CustomOTPVerifyResp
	)

	if bts, err = json.Marshal(CustomOTPVerify{Sub: login, Otp: otp}); err != nil {
		return false, err
	}

	if req, err = http.NewRequest(http.MethodPost, c.config.VerifyURL, bytes.NewBuffer(bts)); err != nil {
		return false, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.config.AuthHeader)

	if res, err = c.client.Do(req); err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return false, errors.New("otp verification request not accepted")
	}

	if err = json.NewDecoder(res.Body).Decode(&ver); err != nil {
		return false, err
	}

	if ver.Ok {
		if err = c.Repo.Set(OTP{
			ID:         GetOTPID(r),
			OTP:        "-", // we don't want to store that
			Expiration: time.Now().Add(OTPExpiration).Unix(),
			Approved:   true,
		}); err != nil {
			return false, err
		}
	}

	return ver.Ok, nil
}

func (c *CustomOTPHandler) IsApproved(r LoginRequest) (bool, error) {
	return IsOtpApproved(c.Repo, r)
}

func (c *CustomOTPHandler) GetDefaultAction() string {
	if c.isOfflineOTPType() {
		return "request"
	}

	return ""
}

func (c *CustomOTPHandler) isOfflineOTPType() bool {
	typ := strings.ToLower(c.config.Type)

	return typ == "totp" || typ == "hotp"
}

var _ OTPHandler = &CustomOTPHandler{}
