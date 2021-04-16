package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIErrorBody struct {
	Error string `json:"error"`
}

type APIError struct {
	Code    int
	Message string
	Error   error
}

func (e *APIError) Unwrap() error {
	return e.Error
}

func Error(c *gin.Context, e APIError) {
	if e.Error != nil {
		logrus.WithError(e.Error).Errorf(e.Error.Error())
	}

	body := APIErrorBody{Error: e.Message}

	c.JSON(e.Code, &body)
}
