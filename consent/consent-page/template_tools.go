package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RenderInvalidRequestError(c *gin.Context, trans *Trans,  err error) {
	message := trans.OrDefault("invalid_request", "Invalid Request")

	RenderError(c, http.StatusBadRequest, message, err)
}

func RenderInternalServerError(c *gin.Context, trans *Trans, err error) {
	message := trans.OrDefault("internal_server_error", "Internal Server Error")

	RenderError(c, http.StatusInternalServerError, message, err)
}

func RenderError(c *gin.Context, statusCode int, msg interface{}, err error) {
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
