package main

import (
	"bytes"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type HyprHandler interface {
	// first get configuration, i.e. appid, machine id, nonces, etc
	// second make the request
	// third poll the request status
	// fourth make decision on status OR handle timeout
	StartAuthentication() (string, error)
	PollHypr(requestID string) (*AuthStatusResponse, error)
	SetStorage(LoginRequest, bool)
	IsApproved(LoginRequest) (bool, error)
}

type DefaultHyprHandler struct {
	HostURL  string
	Client   *http.Client
	APIToken string
	Storage  map[LoginRequest]bool // simple map storage for now
}

// Add setting the host/any other config params
func NewHyprHandler(host string, apiToken string) HyprHandler {
	return &DefaultHyprHandler{
		HostURL: host,
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		APIToken: apiToken,
		Storage:  make(map[LoginRequest]bool),
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

var temptBaseUrl = "https://demo.gethypr.com" //"http://localhost:3031"

func (o *DefaultHyprHandler) StartAuthentication() (string, error) {
	var (
		// TODO move to config
		endpoint = "/rp/api/oob/client/authentication/requests"
		req      *http.Request
		resp     *http.Response
		err      error
	)

	var deviceReq = DeviceAuthenticationRequest{
		SessionNonce:      GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		DeviceNonce:       GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		ServiceHmac:       GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		ServiceNonce:      GenerateSha256(strconv.Itoa(GenerateRandomPin())),
		AppId:             "cloudentity",
		NamedUser:         "billy",
		Machine:           "WEB",
		MachineID:         "1e6680b8a703c7158a8831d472a8969999aaa73b57ea5f2f2ce8d3e7fd48f98d",
		AdditionalDetails: nil,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(deviceReq); err != nil {
		return "", err
	}

	// {{baseUrl}}/rp/api/oob/client/authentication/requests
	if req, err = http.NewRequest("POST", fmt.Sprintf("%s%s", temptBaseUrl, endpoint), &buf); err != nil {
		return "", err
	}

	token := "hypap-125b8ab7-2c86-4855-b77f-ecf5c83c0f05"
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	log.Printf("Request to Hypr: %s", req.URL)
	log.Printf("Request to Hypr: %v", req.Header)
	if resp, err = o.Client.Do(req); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	log.Printf("Status code from Hypr %d", resp.StatusCode)

	data := StartAuthResponse{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	// TODO check Response value and handle
	log.Printf("About to return request ID from data %+v", data.Response)
	return data.Response.RequestID, nil
}

func (o *DefaultHyprHandler) PollHypr(requestID string) (*AuthStatusResponse, error) {
	var (
		endpoint = fmt.Sprintf("/rp/api/oob/client/authentication/requests/%s", requestID)
		req      *http.Request
		resp     *http.Response
		err      error
	)

	// {{baseUrl}}/rp/api/oob/client/authentication/requests/{{requestId}}
	if req, err = http.NewRequest("GET", fmt.Sprintf("%s%s", temptBaseUrl, endpoint), nil); err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	token := "hypap-125b8ab7-2c86-4855-b77f-ecf5c83c0f05"
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if resp, err = o.Client.Do(req); err != nil {
		return nil, err
	}

	data := AuthStatusResponse{}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

type CryptoSource struct{}

func GenerateRandomPin() int {
	var src CryptoSource
	rnd := rand.New(src)
	return rnd.Intn(1000000)
}

func (s CryptoSource) Seed(seed int64) {}

func (s CryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s CryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func GenerateSha256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

type StartAuthResponse struct {
	Status   StatusAuth    `json:"status"`
	Response ResponseStart `json:"response"`
}

type ResponseStart struct {
	RequestID string `json:"requestId"`
}

type StatusAuth struct {
	ResponseCode    int64  `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type AuthStatusResponse struct {
	RequestID string      `json:"requestId"`
	NamedUser string      `json:"namedUser"`
	Machine   string      `json:"machine"`
	Device    DeviceAuth  `json:"device"`
	State     []StateAuth `json:"state"`
}

type DeviceAuth struct {
	DeviceID       string      `json:"deviceId"`
	HmacDeviceKey  interface{} `json:"hmacDeviceKey"`
	HmacSessionKey interface{} `json:"hmacSessionKey"`
}

type StateAuth struct {
	Value     string `json:"value"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

////// --- Old Below

// auth info
type ResponseBase struct {
	Status       string         `json:"status"`
	RawJSON      string         `json:"-"`
	HTTPResponse *http.Response `json:"-"`
}

type DeviceAuthenticationResponse struct {
	ResponseBase
	DeviceAuthenticationResponse *DeviceAuthenticationResponseDTO `json:"deviceAuthenticationResponse,omitempty"`
}

type DeviceAuthenticationResponseDTO struct {
	Status   Status   `json:"status"`
	Response Response `json:"response"`
}

type Status struct {
	ResponseCode    int64  `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Response struct {
	Version   int64    `json:"version"`
	Session   *Session `json:"session"`
	Device    *Device  `json:"device"`
	RequestId *string  `json:"requestId"`
	Links     []Link   `json:"links"`
}

type Session struct {
	SessionId string `json:"sessionId"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Device struct {
	DeviceKey         string `json:"deviceKey"`
	AuthenticationKey string `json:"authenticationKey"`
	FriendlyName      string `json:"friendlyName"`
	DeviceId          string `json:"deviceId"`
	ModelNumber       string `json:"modelNumber"`
}

type DeviceAuthenticationRequest struct {
	SessionNonce      string            `json:"sessionNonce"`
	DeviceNonce       string            `json:"deviceNonce"`
	ServiceHmac       string            `json:"serviceHmac"`
	ServiceNonce      string            `json:"serviceNonce"`
	MachineID         string            `json:"machineId"`
	AppId             string            `json:"appId"`
	NamedUser         string            `json:"namedUser"`
	Machine           string            `json:"machine"`
	AdditionalDetails map[string]string `json:"additionalDetails,omitempty"`
}

// poll response
type DeviceAuthenticationRequestStatusResponse struct {
	ResponseBase
	DeviceAuthenticationRequestStatus *DeviceAuthenticationRequestStatusDTO `json:"deviceAuthenticationRequestStatus,omitempty"`
}

type DeviceAuthenticationRequestStatusDTO struct {
	Device    Device  `json:"device"`
	Links     []Link  `json:"links"`
	Machine   string  `json:"machine"`
	NamedUser string  `json:"namedUser"`
	RequestId string  `json:"requestId"`
	State     []State `json:"state"`
}

type State struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Value     string `json:"value"`
}
