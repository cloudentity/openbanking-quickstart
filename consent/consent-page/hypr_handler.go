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
	BaseUrl string `json:"HYPR_BASE_URL"`
	AppID   string `json:"HYPR_APP_ID"`
}

type DefaultHyprHandler struct {
	HostURL  string
	Client   *http.Client
	APIToken string
	AppId    string
	Storage  map[LoginRequest]bool
}

func NewHyprHandler(hyprConfig HyprConfig) *DefaultHyprHandler {
	return &DefaultHyprHandler{
		HostURL: hyprConfig.BaseUrl,
		AppId:   hyprConfig.AppID,
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		APIToken: hyprConfig.Token,
		Storage:  make(map[LoginRequest]bool),
	}
}

func (h *DefaultHyprHandler) Approve(args map[string]string) *MFAError {
	var (
		devices   UserDevices
		username  string
		requestId string
		ok        bool
		err       error
	)

	if username, ok = args["username"]; !ok {
		return &MFAError{
			err:     errors.New("missing parameter - username required"),
			code:    http.StatusBadRequest,
			message: "missing parameter - username required",
		}
	}

	if devices, err = h.getUserDevices(username); err != nil {
		return &MFAError{
			err:     err,
			code:    http.StatusUnauthorized,
			message: "failed to get user devices",
		}
	}

	if len(devices) < 1 {
		return &MFAError{
			err:     err,
			code:    http.StatusBadGateway,
			message: "no registered devices",
		}
	}

	if requestId, err = h.startAuthentication(username); err != nil {
		return &MFAError{
			err:     err,
			code:    http.StatusInternalServerError,
			message: "failed to start authentication",
		}
	}

	var checkStatus *AuthStatusResponse
	if checkStatus, err = h.poll(requestId); err != nil {
		if errors.Is(err, ErrTimeoutWaitingForUser) {
			return &MFAError{
				err:     err,
				code:    http.StatusUnauthorized,
				message: "timeout waiting for user to approve or denyr",
			}
		}
		return &MFAError{
			err:     err,
			code:    http.StatusInternalServerError,
			message: "failed to check auth status",
		}
	}

	if len(checkStatus.State) == 0 {
		return &MFAError{
			err:     errors.New("failed to check auth status"),
			code:    http.StatusInternalServerError,
			message: "invalid state length",
		}
	}

	switch checkStatus.State[len(checkStatus.State)-1].Value {
	case "COMPLETED":
		return nil
	default:
		return &MFAError{
			err:     errors.New("user denied access"),
			code:    http.StatusUnauthorized,
			message: "user denied access",
		}
	}
}

func (o *DefaultHyprHandler) SetStorage(r LoginRequest, approved bool) {
	o.Storage[r] = approved
}

func (o *DefaultHyprHandler) IsApproved(r LoginRequest) (bool, error) {
	approved, ok := o.Storage[r]
	if !ok {
		return false, nil
	}

	return approved, nil
}

func (h *DefaultHyprHandler) startAuthentication(username string) (string, error) {
	var (
		endpoint = "/rp/api/oob/client/authentication/requests"
		resp     *http.Response
		err      error
	)

	var deviceReq = DeviceAuthenticationRequest{
		SessionNonce:      GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		DeviceNonce:       GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		ServiceHmac:       GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		ServiceNonce:      GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		AppId:             h.AppId,
		NamedUser:         username,
		Machine:           "WEB",
		MachineID:         GenerateRandomString(6),
		AdditionalDetails: nil,
	}

	if resp, err = h.performRequest(http.MethodPost, fmt.Sprintf("%s%s", h.HostURL, endpoint), deviceReq); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := StartAuthResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Response.RequestID, nil
}

func (h *DefaultHyprHandler) poll(requestID string) (*AuthStatusResponse, error) {
	var (
		checkStatus  *AuthStatusResponse
		pollInterval = time.Tick(time.Duration(2) * time.Second)
		timeout      = time.Tick(time.Duration(120) * time.Second)
		err          error
	)

	for {
		select {
		case <-timeout:
			return nil, ErrTimeoutWaitingForUser
		case <-pollInterval:
			if checkStatus, err = h.performAuthStatusCheck(requestID); err != nil {
				return nil, err
			}
			for i, _ := range checkStatus.State {
				switch checkStatus.State[i].Value {
				case "COMPLETED", "CANCELED", "FAILED":
					return checkStatus, nil
				}
			}
		}
	}
}

func (h *DefaultHyprHandler) performAuthStatusCheck(requestID string) (*AuthStatusResponse, error) {
	var (
		endpoint = fmt.Sprintf("/rp/api/oob/client/authentication/requests/%s", requestID)
		resp     *http.Response
		err      error
	)

	if resp, err = h.performRequest(http.MethodGet, fmt.Sprintf("%s%s", h.HostURL, endpoint), nil); err != nil {
		return nil, err
	}

	data := AuthStatusResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (h *DefaultHyprHandler) getUserDevices(username string) (UserDevices, error) {
	var (
		endpoint = fmt.Sprintf("/rp/api/oob/client/authentication/%s/%s/devices", "cloudentity", username)
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

func (h *DefaultHyprHandler) performRequest(method string, endpoint string, payload interface{}) (*http.Response, error) {
	var (
		buf bytes.Buffer
		req *http.Request
		err error
	)

	if payload != nil {
		if err := json.NewEncoder(&buf).Encode(payload); err != nil {
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
