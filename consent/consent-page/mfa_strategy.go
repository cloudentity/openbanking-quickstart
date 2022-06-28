package main

type MFAError struct {
	Err     error
	Code    int
	Message string
}

type MFAApprovalChecker interface {
	IsApproved(r LoginRequest) (bool, error)
}

type MFAStrategy interface {
	Approve(args map[string]string) *MFAError
	IsApproved(LoginRequest) (bool, error)
	SetStorage(LoginRequest, bool)
	MFAApprovalChecker
}
