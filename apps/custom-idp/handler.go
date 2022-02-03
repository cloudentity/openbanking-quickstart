package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ggicci/httpin"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/system/client/logins"
	"github.com/cloudentity/acp-client-go/clients/system/models"
)

func (s *Server) Alive(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// LoginState holds the session keys that we will encode into the OAuth state.
type LoginState struct {
	ID       string `json:"login_id"`
	State    string `json:"login_state"`
	TenantID string `json:"tenant_id"`
}

// LoginRequestInput holds the parameters from the request URL.
type LoginRequestInput struct {
	ID          string `in:"form=login_id;required"`
	State       string `in:"form=login_state;required"`
	ClientID    string `in:"form=client_id;required"`
	IdpClientID string `in:"form=idp_client_id;required"`
	ServerID    string `in:"form=server_id;required"`
	TenantID    string `in:"form=tenant_id;required"`
	TenantURL   string `in:"form=tenant_url;required"`
}

// Login is the handler for the /login endpoint.
func (s *Server) Login(c *gin.Context) {
	input, _ := c.Request.Context().Value(httpin.Input).(*LoginRequestInput)

	state, err := json.Marshal(LoginState{
		ID:       input.ID,
		State:    input.State,
		TenantID: input.TenantID,
	})
	if err != nil {
		logrus.WithError(err).Error("marshal state")
		return
	}

	authorizeURL := s.Config.OIDC.AuthorizeURL("", string(state))
	logrus.WithField("location", authorizeURL).Info("/login redirecting to external oauth/authorize")

	c.Redirect(http.StatusTemporaryRedirect, authorizeURL)
}

// CallbackInput holds the parameters from the request URL.
type CallbackInput struct {
	Code  string `in:"form=code;required"`
	State string `in:"form=state;required"`
}

// Callback handles the redirect from the external IDP to /callback.
func (s *Server) Callback(c *gin.Context) {
	var (
		tokens acpclient.Token
		state  LoginState
		err    error
	)
	input, _ := c.Request.Context().Value(httpin.Input).(*CallbackInput)

	if err = json.Unmarshal([]byte(input.State), &state); err != nil {
		logrus.WithError(err).WithField("state", input.State).Error("unmarshal state")
		return
	}

	// Exchange code for access and ID tokens.
	if tokens, err = s.OidcClient.Exchange(input.Code, "", acpclient.CSRF{}); err != nil {
		logrus.WithError(err).Error("Exchange code for tokens")
		return
	}

	// At this point you have the OAuth Access and ID tokens in `tokens`.
	// So you can interact with your system, before accepting the login in ACP.
	if err = DoMyCustomStuff(s, c, tokens); err != nil {
		logrus.WithError(err).Error("DoMyCustomStuff failed")
		return
	}

	if err == nil {
		s.AcceptLogin(c, state, tokens)
	} else {
		s.RejectLogin(c, state)
	}
}

// DoMyCustomStuff can be used to implement your own interactions.
// The gin.Context can access the http.Request and ResponseWriter.
func DoMyCustomStuff(s *Server, c *gin.Context, tokens acpclient.Token) error {
	return nil
}

func (s *Server) AcceptLogin(c *gin.Context, login LoginState, tokens acpclient.Token) {
	var (
		parser jwt.Parser
		claims jwt.MapClaims
		err    error
	)

	if _, _, err = parser.ParseUnverified(tokens.IDToken, &claims); err != nil {
		logrus.WithError(err).Error("parsing id token")
	}
	logrus.Debugf("claims: %#v", claims)

	acceptLogin := models.AcceptSession{
		Acr:        ClaimAsString(claims, "acr"),      // authentication context class reference
		Amr:        ClaimAsStringSlice(claims, "amr"), // authentication methods references
		AuthTime:   strfmt.DateTime(time.Now()),
		ID:         login.ID,
		LoginState: login.State,
		Subject:    ClaimAsString(claims, "sub"),
		AuthenticationContext: map[string]interface{}{
			"access_token": tokens.AccessToken,
			"id_token":     tokens.IDToken,
		},
	}
	if err = acceptLogin.Validate(nil); err != nil {
		logrus.WithError(err).Error("AcceptLogin.Validate failed")
		c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
		return
	}

	res, err := s.AcpClient.System.Logins.AcceptLoginRequest(
		logins.NewAcceptLoginRequestParams().
			WithContext(c).
			WithLogin(login.ID).
			WithAcceptLogin(&acceptLogin),
		nil,
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
		logrus.WithField("location", res.Payload.RedirectTo).Info("acp login accepted, redirecting")
		c.Redirect(http.StatusTemporaryRedirect, res.Payload.RedirectTo)
		return
	}
	logrus.WithField("location", res.Payload.RedirectTo).Info("login accepted, OK")
	c.String(http.StatusOK, "AcceptLoginRequest succeeded")
}

func (s *Server) RejectLogin(c *gin.Context, login LoginState) {
	var err error

	rejectLogin := models.RejectSession{
		ID:         login.ID,
		LoginState: login.State,
		// There are also fields for Error, ErrorDescription and StatusCode.
	}
	if err = rejectLogin.Validate(nil); err != nil {
		logrus.WithError(err).Error("rejectLogin.Validate failed")
		c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
		return
	}

	res, err := s.AcpClient.System.Logins.RejectLoginRequest(
		logins.NewRejectLoginRequestParams().
			WithContext(c).
			WithLogin(login.ID).
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
		logrus.WithField("location", res.Payload.RedirectTo).Info("acp login rejected, redirecting")
		c.Redirect(http.StatusTemporaryRedirect, res.Payload.RedirectTo)
		return
	}
	logrus.WithField("location", s.Config.FailureURL).Info("login rejected, redirecting")
	c.Redirect(http.StatusTemporaryRedirect, s.Config.FailureURL)
}

// BindInput instances an httpin engine for a input struct as a gin middleware.
// See github.com/ggicci/httpin/wiki/Integrate-with-gin
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

func ClaimAsString(claims jwt.MapClaims, name string) string {
	if val, ok := claims[name].(string); ok {
		return val
	}
	logrus.Errorf("claim %q is not a string", name)
	return ""
}

func ClaimAsStringSlice(claims jwt.MapClaims, name string) []string {
	if vals, ok := claims[name].([]interface{}); ok {
		strs := make([]string, 0, len(vals))
		for _, val := range vals {
			if str, ok := val.(string); ok {
				strs = append(strs, str)
			}
		}
		return strs
	}
	logrus.Errorf("claim %q is not a slice", name)
	return []string{}
}

// ErrorPayload returns the *models.Error for errors that have it.
func ErrorPayload(err error) (*models.Error, bool) {
	switch e := err.(type) { // nolint
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
