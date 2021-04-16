package main

import "fmt"

type ErrNotFound struct {
	resourceName string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("unable to find resource %s", e.resourceName)
}

type ErrAlreadyExists struct {
	resourceName string
}

func (e ErrAlreadyExists) Error() string {
	return fmt.Sprintf("resource %s already exists", e.resourceName)
}
