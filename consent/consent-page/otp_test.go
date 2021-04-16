package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOTP(t *testing.T) {
	db, err := InitDB(Config{DBFile: "./data/test.db"})
	require.NoError(t, err)

	otpRepo, err := NewOTPRepo(db)
	require.NoError(t, err)

	h := DefaultOTPHandler{Repo: otpRepo}
	r := LoginRequest{
		ID:          "123",
		State:       "abc",
		ConsentType: "account_access",
	}
	otp, err := h.Generate(r)
	require.NoError(t, err)

	require.NotEmpty(t, otp.OTP)

	err = h.Store(otp)
	require.NoError(t, err)

	approved, err := h.IsApproved(r)
	require.NoError(t, err)
	require.Equal(t, false, approved)

	valid, err := h.Verify(r, "invalid")
	require.NoError(t, err)
	require.Equal(t, false, valid)

	valid, err = h.Verify(r, otp.OTP)
	require.NoError(t, err)
	require.Equal(t, true, valid)

	approved, err = h.IsApproved(r)
	require.NoError(t, err)
	require.Equal(t, true, approved)
}

func TestMaskMobile(t *testing.T) {
	require.Equal(t, "+48******321", MaskMobile("+48987654321"))
}
