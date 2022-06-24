package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type HyprConfig struct {
	Token   string `json:"HYPR_TOKEN"`
	BaseURL string `json:"HYPR_BASE_URL"`
	AppID   string `json:"HYPR_APP_ID"`
}

type HyprStrategy struct {
	HostURL  string
	Client   *http.Client
	APIToken string
	AppID    string
	Storage  map[LoginRequest]bool
}

func NewHyprStrategy(hyprConfig HyprConfig) *HyprStrategy {
	return &HyprStrategy{
		HostURL: hyprConfig.BaseURL,
		AppID:   hyprConfig.AppID,
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		APIToken: hyprConfig.Token,
		Storage:  make(map[LoginRequest]bool),
	}
}

func (h *HyprStrategy) Approve(args map[string]string) *MFAError {
	var (
		devices   UserDevices
		username  string
		requestID string
		ok        bool
		err       error
	)

	if username, ok = args["username"]; !ok {
		return &MFAError{
			Err:     errors.New("missing parameter - username required"),
			Code:    http.StatusBadRequest,
			Message: "missing parameter - username required",
		}
	}

	if devices, err = h.getUserDevices(username); err != nil {
		return &MFAError{
			Err:     err,
			Code:    http.StatusUnauthorized,
			Message: "failed to get user devices",
		}
	}

	if len(devices) < 1 {
		return &MFAError{
			Err:     err,
			Code:    http.StatusBadGateway,
			Message: "no registered devices",
		}
	}

	if requestID, err = h.startAuthentication(username); err != nil {
		return &MFAError{
			Err:     err,
			Code:    http.StatusInternalServerError,
			Message: "failed to start authentication",
		}
	}

	var checkStatus AuthStatusResponse
	if checkStatus, err = h.poll(requestID); err != nil {
		if errors.Is(err, ErrTimeoutWaitingForUser) {
			return &MFAError{
				Err:     err,
				Code:    http.StatusUnauthorized,
				Message: "timeout waiting for user to approve or denyr",
			}
		}
		return &MFAError{
			Err:     err,
			Code:    http.StatusInternalServerError,
			Message: "failed to check auth status",
		}
	}

	if len(checkStatus.State) == 0 {
		return &MFAError{
			Err:     errors.New("failed to check auth status"),
			Code:    http.StatusInternalServerError,
			Message: "invalid state length",
		}
	}

	switch checkStatus.State[len(checkStatus.State)-1].Value {
	case "COMPLETED":
		return nil
	default:
		return &MFAError{
			Err:     errors.New("user denied access"),
			Code:    http.StatusUnauthorized,
			Message: "user denied access",
		}
	}
}

func (h *HyprStrategy) SetStorage(r LoginRequest, approved bool) {
	h.Storage[r] = approved
}

func (h *HyprStrategy) IsApproved(r LoginRequest) (bool, error) {
	approved, ok := h.Storage[r]
	if !ok {
		return false, nil
	}

	return approved, nil
}

func (h *HyprStrategy) startAuthentication(username string) (string, error) {
	var (
		endpoint = "/rp/api/oob/client/authentication/requests"
		resp     *http.Response
		err      error
	)

	deviceReq := DeviceAuthenticationRequest{
		SessionNonce:      GenerateSha256(strconv.Itoa(int(GenerateRandomPin()))),
		DeviceNonce:       GenerateSha256(strconv.Itoa(int(GenerateRandomPin()))),
		ServiceHmac:       GenerateSha256(strconv.Itoa(int(GenerateRandomPin()))),
		ServiceNonce:      GenerateSha256(strconv.Itoa(int(GenerateRandomPin()))),
		AppID:             h.AppID,
		NamedUser:         username,
		Machine:           "WEB",
		MachineID:         GenerateRandomString(6),
		AdditionalDetails: nil,
	}

	if resp, err = h.performRequest(http.MethodPost, fmt.Sprintf("%s%s", h.HostURL, endpoint), deviceReq); err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error starting authentication")
	}

	data := StartAuthResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Response.RequestID, nil
}

func (h *HyprStrategy) poll(requestID string) (AuthStatusResponse, error) {
	var (
		checkStatus  AuthStatusResponse
		pollInterval = time.Tick(time.Duration(2) * time.Second)
		timeout      = time.Tick(time.Duration(120) * time.Second)
		err          error
	)

	for {
		select {
		case <-timeout:
			return checkStatus, ErrTimeoutWaitingForUser
		case <-pollInterval:
			if checkStatus, err = h.performAuthStatusCheck(requestID); err != nil {
				return checkStatus, err
			}
			for i := range checkStatus.State { // no lint
				switch checkStatus.State[i].Value {
				case "COMPLETED", "CANCELED", "FAILED":
					return checkStatus, nil
				}
			}
		}
	}
}

func (h *HyprStrategy) performAuthStatusCheck(requestID string) (AuthStatusResponse, error) {
	var (
		endpoint = fmt.Sprintf("/rp/api/oob/client/authentication/requests/%s", requestID)
		data     = AuthStatusResponse{}
		resp     *http.Response
		err      error
	)

	if resp, err = h.performRequest(http.MethodGet, fmt.Sprintf("%s%s", h.HostURL, endpoint), nil); err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func (h *HyprStrategy) getUserDevices(username string) (UserDevices, error) {
	var (
		endpoint = fmt.Sprintf("/rp/api/oob/client/authentication/%s/%s/devices", h.AppID, username)
		resp     *http.Response
		err      error
	)

	if resp, err = h.performRequest(http.MethodGet, fmt.Sprintf("%s%s", h.HostURL, endpoint), nil); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := UserDevices{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (h *HyprStrategy) performRequest(method string, endpoint string, payload interface{}) (*http.Response, error) {
	var (
		buf bytes.Buffer
		req *http.Request
		err error
	)

	if payload != nil {
		if err = json.NewEncoder(&buf).Encode(payload); err != nil {
			return nil, err
		}
	}

	if req, err = http.NewRequest(method, endpoint, &buf); err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.APIToken))

	return h.Client.Do(req)
}
