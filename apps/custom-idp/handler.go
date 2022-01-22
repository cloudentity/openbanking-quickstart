package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ggicci/httpin"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/logins"
	"github.com/cloudentity/acp-client-go/models"
)

func (s *Server) Alive(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

type LoginRequestInput struct {
	ID          string `in:"form=login_id;required"`
	State       string `in:"form=login_state;required"`
	ClientID    string `in:"form=client_id;required"`
	IdpClientID string `in:"form=idp_client_id;required"`
	ServerID    string `in:"form=server_id;required"`
	TenantID    string `in:"form=tenant_id;required"`
	TenantURL   string `in:"form=tenant_url;required"`
	ACRValues   string `in:"form=acr_values"`
}

func (s *Server) Login(c *gin.Context) {
	var input = c.Request.Context().Value(httpin.Input).(*LoginRequestInput)

	// Authenticate the user in the external IDP, using the ACR values.
	user, err := AuthenticateUser(input.ACRValues)

	if err == nil {
		s.AcceptLogin(c, input, user)
	} else {
		s.RejectLogin(c, input)
	}
}

func (s *Server) AcceptLogin(
	c *gin.Context,
	input *LoginRequestInput,
	subject string,
) {
	var err error

	acceptLogin := models.AcceptSession{
		Acr:        "",         // authentication context class reference
		Amr:        []string{}, // authentication methods references
		AuthTime:   strfmt.DateTime(time.Now()),
		ID:         input.ID,
		LoginState: input.State,
		Subject:    subject,
		AuthenticationContext: map[string]interface{}{
			"phone_number": "12015555309",
		},
	}
	if err = acceptLogin.Validate(nil); err != nil {
		logrus.WithError(err).Error("AcceptLogin.Validate failed")
		c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
		return
	}

	res, err := s.AcpClient.Logins.AcceptLoginRequest(
		logins.NewAcceptLoginRequestParams().
			WithContext(c).
			WithLogin(input.ID).
			WithTid(input.TenantID).
			WithAcceptLogin(&acceptLogin),
		nil, // When would this authinfo param be needed?
	)
	if err != nil {
		if payload, ok := ErrorPayload(err); ok {
			logrus.WithError(err).Error(payload.Error)
		} else {
			logrus.WithError(err).Error("AcceptLoginRequest failed")
		}
		c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
		return
	}
	if res.Payload.RedirectTo != "" {
		logrus.WithField("location", res.Payload.RedirectTo).Info("login accepted, redirecting")
		c.Redirect(http.StatusTemporaryRedirect, res.Payload.RedirectTo)
		return
	}
	logrus.WithField("location", res.Payload.RedirectTo).Info("login accepted, OK")
	c.String(http.StatusOK, "AcceptLoginRequest succeeded")
}

func (s *Server) RejectLogin(c *gin.Context, input *LoginRequestInput) {
	var err error

	rejectLogin := models.RejectSession{
		ID:         input.ID,
		LoginState: input.State,
		// There are also fields for Error, ErrorDescription and StatusCode.
	}
	if err = rejectLogin.Validate(nil); err != nil {
		logrus.WithError(err).Error("rejectLogin.Validate failed")
		c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
		return
	}

	res, err := s.AcpClient.Logins.RejectLoginRequest(
		logins.NewRejectLoginRequestParams().
			WithContext(c).
			WithLogin(input.ID).
			WithTid(input.TenantID).
			WithRejectLogin(&rejectLogin),
		nil,
	)
	if err != nil {
		if payload, ok := ErrorPayload(err); ok {
			logrus.WithError(err).Error(payload.Error)
		} else {
			logrus.WithError(err).Error("RejectLoginRequest failed")
		}
		c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
		return
	}
	if res.Payload.RedirectTo != "" {
		logrus.WithField("location", res.Payload.RedirectTo).Info("login rejected, redirecting")
		c.Redirect(http.StatusTemporaryRedirect, res.Payload.RedirectTo)
		return
	}
	logrus.WithField("location", s.Config.FailureURL).Info("login rejected, redirecting")
	c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
	return
}

// BindInput instances an httpin engine for a input struct as a gin middleware.
// See https://github.com/ggicci/httpin/wiki/Integrate-with-gin
func BindInput(inputStruct interface{}) gin.HandlerFunc {
	engine, err := httpin.New(inputStruct)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		input, err := engine.Decode(c.Request)
		if err != nil {
			var invalidFieldError *httpin.InvalidFieldError
			if errors.As(err, &invalidFieldError) {
				c.AbortWithStatusJSON(http.StatusBadRequest, invalidFieldError)
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(c.Request.Context(), httpin.Input, input)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// ErrorPayload returns the *models.Error for errors that have it.
func ErrorPayload(err error) (*models.Error, bool) {
	switch e := err.(type) {
	case *logins.AcceptLoginRequestUnauthorized:
		return e.Payload, true
	case *logins.AcceptLoginRequestForbidden:
		return e.Payload, true
	case *logins.AcceptLoginRequestNotFound:
		return e.Payload, true
	case *logins.RejectLoginRequestUnauthorized:
		return e.Payload, true
	case *logins.RejectLoginRequestForbidden:
		return e.Payload, true
	case *logins.RejectLoginRequestNotFound:
		return e.Payload, true
	default:
		return nil, false
	}
}
