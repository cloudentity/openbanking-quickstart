package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
	"math/rand"

	"github.com/pkg/errors"
)

var ErrTimeoutWaitingForUser = errors.New("Request expired before response")

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

type UserDevices []UserDevice

type UserDevice struct {
	ID                interface{}  `json:"id"`
	DeviceID          string       `json:"deviceId"`
	DeviceType        string       `json:"deviceType"`
	ProtocolVersion   interface{}  `json:"protocolVersion"`
	FriendlyName      string       `json:"friendlyName"`
	ModelNumber       string       `json:"modelNumber"`
	CreateDate        int64        `json:"createDate"`
	LastLoginDate     int64        `json:"lastLoginDate"`
	PushID            string       `json:"pushId"`
	FbDeviceInfo      FbDeviceInfo `json:"fbDeviceInfo"`
	DeviceAttributes  interface{}  `json:"deviceAttributes"`
	MachineID         string       `json:"machineId"`
	DeviceKey         string       `json:"deviceKey"`
	AuthenticationKey string       `json:"authenticationKey"`
	FidoUsername      string       `json:"fidoUsername"`
	NamedUser         string       `json:"namedUser"`
	RegisteredUser    interface{}  `json:"registeredUser"`
}

type FbDeviceInfo struct {
	ID                  string `json:"id"`
	NotificationKey     string `json:"notificationKey"`
	RegistrationToken   string `json:"registrationToken"`
	Brand               string `json:"brand"`
	NotificationKeyName string `json:"notificationKeyName"`
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
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
