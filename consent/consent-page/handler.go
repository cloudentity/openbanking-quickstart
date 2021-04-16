package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	ID          string
	State       string
	ConsentType string
}

func NewLoginRequest(c *gin.Context) LoginRequest {
	return LoginRequest{
		ID:          c.Query("login_id"),
		State:       c.Query("login_state"),
		ConsentType: c.Query("consent_type"),
	}
}

func (l *LoginRequest) Validate() error {
	if l.ID == "" || l.State == "" || l.ConsentType == "" {
		return errors.New("login_id / login_state / consent_type missing")
	}

	return nil
}

func (s *Server) WithConsentHandler(c *gin.Context) (SpecificConsentHandler, LoginRequest, error) {
	var (
		loginRequest = NewLoginRequest(c)
		handler      SpecificConsentHandler
		err          error
		ok           bool
	)

	if err = loginRequest.Validate(); err != nil {
		return handler, loginRequest, err
	}

	if handler, ok = s.GetConsentHandler(loginRequest); !ok {
		return handler, loginRequest, fmt.Errorf("invalid consent type %s", loginRequest.ConsentType)
	}

	return handler, loginRequest, nil
}

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			handler      SpecificConsentHandler
			loginRequest LoginRequest
			err          error
		)

		if handler, loginRequest, err = s.WithConsentHandler(c); err != nil {
			RenderInvalidRequestError(c, err)
			return
		}

		handler.GetConsent(c, loginRequest)
	}
}

func (s *Server) Post() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			handler SpecificConsentHandler

			loginRequest LoginRequest
			err          error
		)

		if handler, loginRequest, err = s.WithConsentHandler(c); err != nil {
			RenderInvalidRequestError(c, err)
			return
		}

		s.PostConsent(c, loginRequest, handler)
	}
}

func (s *Server) PostConsent(c *gin.Context, loginRequest LoginRequest, consentHandler SpecificConsentHandler) {
	var (
		redirect string
		err      error
	)

	action := c.PostForm("action")

	switch action {
	case "confirm":
		redirect, err = consentHandler.ConfirmConsent(c, loginRequest)
	case "deny":
		redirect, err = consentHandler.DenyConsent(c, loginRequest)
	default:
		RenderInvalidRequestError(c, fmt.Errorf("invalid form action: %s", action))
		return
	}

	if err != nil {
		RenderInternalServerError(c, errors.Wrapf(err, "failed to %s consent: %+v", c.PostForm("action"), loginRequest.ConsentType))
		return
	}

	logrus.Debugf("redirect to: %s", redirect)

	c.Redirect(http.StatusFound, redirect)
}

func (s *Server) GetConsentHandler(loginRequest LoginRequest) (SpecificConsentHandler, bool) {
	var handler SpecificConsentHandler

	switch loginRequest.ConsentType {
	case "domestic_payment":
		handler = &DomesticPaymentConsentHandler{s, ConsentTools{}}
	case "account_access":
		handler = &AccountAccessConsentHandler{s, ConsentTools{}}
	default:
		return nil, false
	}
	return handler, true
}

type SpecificConsentHandler interface {
	GetConsent(c *gin.Context, loginRequest LoginRequest)
	ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (redirect string, err error)
	DenyConsent(c *gin.Context, loginRequest LoginRequest) (redirect string, err error)
}
