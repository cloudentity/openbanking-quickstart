package main

type MFAError struct {
	err     error
	code    int
	message string
}

type MFAStrategy interface {
	Approve(args map[string]string) *MFAError
	IsApproved(LoginRequest) (bool, error)
	SetStorage(LoginRequest, bool)
}
