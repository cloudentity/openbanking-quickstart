package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type CustomOTPHandler struct {
	config OtpConfig
	client *http.Client
	Repo   *OTPRepo
}

func NewCustomOTPHandler(config OtpConfig, repo *OTPRepo) OTPHandler {
	return &CustomOTPHandler{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
		},
		Repo: repo,
	}
}

type CustomOtpRequest struct {
	Sub string `json:"sub"`
}

func (c *CustomOTPHandler) Send(r LoginRequest, provider MFAConsentProvider, to string, data MFAData) error {
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.config.AuthHeader)

	if res, err = c.client.Do(req); err != nil {
		return false, err
	}

	if res.StatusCode >= 400 {
		return false, errors.New("otp verification request not accepted")
	}

	if err = json.NewDecoder(res.Body).Decode(&ver); err != nil {
		return false, err
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