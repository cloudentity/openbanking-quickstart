package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	Success = CustomOTPVerifyResp{Ok: true}
	Failure = CustomOTPVerifyResp{Ok: false}
)

func (s *Server) DemoTotpVerify(ctx *gin.Context) {
	var (
		decoder = json.NewDecoder(ctx.Request.Body)
		request CustomOTPVerify
		err     error
		auth    = ctx.GetHeader("Authorization")
	)
	if auth != "Bearer this-is-some-test-token" {
		logrus.WithField("auth", auth).Debug("Unauthorized")

		ctx.AbortWithStatus(401)
		return
	}

	if err = decoder.Decode(&request); err != nil {
		logrus.Debug("failed to parse request")

		ctx.AbortWithStatusJSON(400, Failure)
		return
	}

	if request.Otp == "123456" && request.Sub == "user" {
		ctx.JSON(200, Success)
		return
	}
}
