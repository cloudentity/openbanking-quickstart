package main

import (
	"net/http"

	"github.com/pkg/errors"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"-"`
}

func (e *Error) WithError(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Error:   errors.WithStack(err),
	}
}

func (e *Error) WithMessage(message string) *Error {
	return &Error{
		Code:    e.Code,
		Message: message,
		Error:   e.Error,
	}
}

func (e *Error) WithCode(code int) *Error {
	return &Error{
		Code:    code,
		Message: e.Message,
		Error:   e.Error,
	}
}

var (
	ErrNotFound = &Error{
		Code:    http.StatusNotFound,
		Message: "entity not found",
	}
	ErrForbidden = &Error{
		Code:    http.StatusForbidden,
		Message: "operation forbidden",
	}
	ErrAlreadyExists = &Error{
		Code:    http.StatusConflict,
		Message: "entity already exists",
	}
	ErrBadRequest = &Error{
		Code: http.StatusBadRequest,
	}
	ErrUnprocessableEntity = &Error{
		Code: http.StatusUnprocessableEntity,
	}
	ErrInternalServer = &Error{
		Code:    http.StatusInternalServerError,
		Message: "unknown error",
	}
)
