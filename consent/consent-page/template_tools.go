package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RenderInvalidRequestError(c *gin.Context, err error) {
	RenderError(c, http.StatusBadRequest, "Invalid request", err)
}

func RenderInternalServerError(c *gin.Context, err error) {
	RenderError(c, http.StatusInternalServerError, "Internal Server Error", err)
}

func RenderError(c *gin.Context, statusCode int, msg string, err error) {
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
	}

	c.HTML(statusCode, "error.tmpl", gin.H{
		"msg": msg,
	})
}

func Render(c *gin.Context, templateName string, data map[string]interface{}) {
	c.HTML(http.StatusOK, templateName, data)
}
