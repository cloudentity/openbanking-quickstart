package main

import (
	"errors"
	"fmt"
)

var (
	errNotFound            ErrNotFound
	errAlreadyExists       ErrAlreadyExists
	errForbidden           ErrForbidden
	errInternalServer      ErrInternalServer
	errUnprocessableEntity ErrUnprocessableEntity
)

type ErrNotFound struct {
	resourceName string
}

func NewErrNotFound(name string) ErrNotFound {
	return ErrNotFound{name}
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("unable to find resource %s", e.resourceName)
}

type ErrAlreadyExists struct {
	resourceName string
}

func NewErrAlreadyExists(name string) ErrAlreadyExists {
	return ErrAlreadyExists{name}
}

func (e ErrAlreadyExists) Error() string {
	return fmt.Sprintf("resource %s already exists", e.resourceName)
}

type ErrForbidden struct {
	err error
}

func NewErrForbidden(message string) ErrForbidden {
	return ErrForbidden{errors.New(message)}
}

func (e ErrForbidden) Error() string {
	return e.err.Error()
}

type ErrInternalServer struct {
	err error
}

func NewErrInternalServer(message string) ErrInternalServer {
	return ErrInternalServer{errors.New(message)}
}

func (e ErrInternalServer) Error() string {
	return e.err.Error()
}

type ErrUnprocessableEntity struct {
	err error
}

func NewErrUnprocessableEntity(message string) ErrUnprocessableEntity {
	return ErrUnprocessableEntity{errors.New(message)}
}

func (e ErrUnprocessableEntity) Error() string {
	return e.err.Error()
}
