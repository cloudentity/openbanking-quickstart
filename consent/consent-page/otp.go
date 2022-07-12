package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

type OTPID string

func GetOTPID(r LoginRequest) OTPID {
	return OTPID(fmt.Sprintf("%s_%s_%s", r.ID, r.State, r.ConsentType))
}

type OTP struct {
	ID         OTPID  `json:"id"`
	OTP        string `json:"otp"`
	Expiration int64  `json:"expiration"`
	Approved   bool   `json:"approved"`
}

type OTPRepo struct {
	*bolt.DB
}

var bucket = []byte("otps")

func NewOTPRepo(db *bolt.DB) (*OTPRepo, error) {
	var (
		repo = OTPRepo{}
		err  error
	)

	if err = db.Update(func(tx *bolt.Tx) error {
		if _, err = tx.CreateBucketIfNotExists(bucket); err != nil {
			return errors.Wrapf(err, "failed to create bucket")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	repo.DB = db

	return &repo, nil
}

func (u *OTPRepo) Get(id OTPID) (OTP, error) {
	var (
		otp OTP
		bs  []byte
		k   []byte
		err error
	)

	if err = u.DB.View(func(tx *bolt.Tx) error {
		k, bs = tx.Bucket(bucket).Cursor().Seek([]byte(id))

		if bs == nil || !bytes.Equal(k, []byte(id)) {
			otp = OTP{ID: id}

			return nil
		}

		if err = json.Unmarshal(bs, &otp); err != nil {
			return errors.Wrapf(err, "failed to unmarshal otp")
		}

		return nil
	}); err != nil {
		return otp, errors.Wrapf(err, "failed to get otp")
	}

	return otp, nil
}

func (u *OTPRepo) Set(otp OTP) error {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(&otp); err != nil {
		return errors.Wrapf(err, "failed to marshal otp")
	}

	if err = u.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).Put([]byte(otp.ID), bs)
	}); err != nil {
		return errors.Wrapf(err, "failed to update otp")
	}

	return nil
}

var (
	OTPLength     = 6
	OTPExpiration = time.Minute * 5
)

type OTPHandler interface {
	Send(r LoginRequest, provider MFAConsentProvider, to string, data MFAData) error
	Verify(r LoginRequest, login string, otp string) (bool, error)
	MFAApprovalChecker
	// GetDefaultAction can be used to start mfa flow from different a different point.
	// Example:
	// For TOTP we are skipping sending a request for OTP as it is generate by the user
	GetDefaultAction() string
}

func NewOTPHandler(config Config, otpRepo *OTPRepo, smsClient *SMSClient) (OTPHandler, error) {
	switch config.OTPMode {
	case "mock":
		m := NewMockOTPHandler()
		return &m, nil
	case "custom":
		return NewCustomOTPHandler(config, otpRepo)
	}

	return &DemoOTPHandler{Repo: otpRepo, SMSClient: smsClient}, nil
}

type MockOTPHandler struct {
	Storage map[OTPID]OTP
}

func NewMockOTPHandler() MockOTPHandler {
	return MockOTPHandler{Storage: map[OTPID]OTP{}}
}

func (m *MockOTPHandler) Generate(r LoginRequest) (OTP, error) {
	return OTP{
		ID:         GetOTPID(r),
		OTP:        "111111",
		Expiration: time.Now().Add(OTPExpiration).Unix(),
	}, nil
}

func (m *MockOTPHandler) Store(otp OTP) error {
	m.Storage[otp.ID] = otp
	return nil
}

func (m *MockOTPHandler) Send(r LoginRequest, provider MFAConsentProvider, to string, data MFAData) error {
	var (
		otp OTP
		err error
	)

	if otp, err = m.Generate(r); err != nil {
		return errors.Wrap(err, "failed to generate otp")
	}

	if err = m.Store(otp); err != nil {
		return errors.Wrap(err, "failed to store otp")
	}

	return nil
}

func (m *MockOTPHandler) Verify(r LoginRequest, login string, otp string) (bool, error) {
	var (
		id  = GetOTPID(r)
		ok  bool
		o   OTP
		err error
	)
	if o, ok = m.Storage[id]; !ok {
		return false, nil
	}

	if o.OTP == otp {
		o.Approved = true

		if err = m.Store(o); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (m *MockOTPHandler) IsApproved(r LoginRequest) (bool, error) {
	o, ok := m.Storage[GetOTPID(r)]
	if !ok {
		return false, nil
	}

	return o.Approved, nil
}

func (m *MockOTPHandler) GetDefaultAction() string {
	return ""
}

type DemoOTPHandler struct {
	Repo      *OTPRepo
	SMSClient *SMSClient
}

func (o *DemoOTPHandler) Generate(r LoginRequest) (OTP, error) {
	var (
		otp    OTP
		otpStr string
		err    error
	)

	if otpStr, err = RandomOTP(OTPLength); err != nil {
		return otp, err
	}

	return OTP{
		ID:         GetOTPID(r),
		OTP:        otpStr,
		Expiration: time.Now().Add(OTPExpiration).Unix(),
	}, nil
}

func (o *DemoOTPHandler) Store(otp OTP) error {
	return o.Repo.Set(otp)
}

func (o *DemoOTPHandler) Send(r LoginRequest, provider MFAConsentProvider, to string, data MFAData) error {
	var (
		otp OTP
		err error
	)

	if otp, err = o.Generate(r); err != nil {
		return errors.Wrap(err, "failed to generate otp")
	}

	if err = o.Store(otp); err != nil {
		return errors.Wrap(err, "failed to store otp")
	}

	return o.SMSClient.Send(to, provider.GetSMSBody(data, otp))
}

func (o *DemoOTPHandler) Verify(r LoginRequest, login string, otp string) (bool, error) {
	var (
		storedOtp OTP
		err       error
	)

	if storedOtp, err = o.Repo.Get(GetOTPID(r)); err != nil {
		return false, err
	}

	if storedOtp.OTP == otp && time.Unix(storedOtp.Expiration, 0).After(time.Now()) {
		storedOtp.Approved = true

		if err = o.Repo.Set(storedOtp); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (o *DemoOTPHandler) IsApproved(r LoginRequest) (bool, error) {
	return IsOtpApproved(o.Repo, r)
}

func (o *DemoOTPHandler) GetDefaultAction() string {
	return ""
}

func IsOtpApproved(repo *OTPRepo, r LoginRequest) (bool, error) {
	var (
		storedOtp OTP
		err       error
	)

	if storedOtp, err = repo.Get(GetOTPID(r)); err != nil {
		return false, err
	}

	return storedOtp.Approved, nil
}

var otpChars = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func RandomOTP(length int) (string, error) {
	var err error

	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = otpChars[int(b[i])%len(otpChars)]
	}

	return string(b), nil
}

func MaskMobile(mobile string) string {
	var sb strings.Builder

	for i, m := range mobile {
		if i < 3 || i > 8 {
			sb.WriteRune(m)
		} else {
			sb.WriteString("*")
		}
	}

	return sb.String()
}
