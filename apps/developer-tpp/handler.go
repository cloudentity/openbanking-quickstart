package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	acpclient "github.com/cloudentity/acp-client-go"
)

type AppStorage struct {
	CSRF acpclient.CSRF
}

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, s.GetTemplate("login.tmpl"), gin.H{})
	}
}

type SpecLogicHandler interface {
	AccountsGetter
	LoginURLBuilder
	ConsentCreator
	DoRequestObjectEncryption() bool
	PostAuthenticationAction(*gin.Context, map[string]interface{}) (map[string]interface{}, error)
}

type AccountsGetter interface {
	GetAccounts(*gin.Context, string) (interface{}, error)
}

type LoginURLBuilder interface {
	BuildLoginURL(*gin.Context, string, bool) (string, acpclient.CSRF, error)
}

type ConsentCreator interface {
	CreateConsent(*gin.Context) (interface{}, error)
	GetConsentID(interface{}) string
}

func (s *Server) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			registerResponse   interface{}
			encodedCookieValue string
			storage            AppStorage
			loginURL           string
			u                  *url.URL
			data               = gin.H{}
			err                error
		)

		if registerResponse, err = s.CreateConsent(c); err != nil {
			// todo fix retry mechnism in acp-go-client
			if registerResponse, err = s.CreateConsent(c); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to register consent: %+v", err))
				return
			}
		}

		consentID := s.GetConsentID(registerResponse)

		registerResponseRaw, _ := json.MarshalIndent(registerResponse, "", "  ")
		data["account_access_consent_raw"] = string(registerResponseRaw)

		if loginURL, storage.CSRF, err = s.BuildLoginURL(c, consentID, false); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
			return
		}

		if u, err = url.Parse(loginURL); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to parse login url: %+v", err))
			return
		}

		rp := u.Query()["request"]

		if len(rp) > 0 {
			parser := jwt.Parser{}
			claims := jwt.MapClaims{}
			IDToken, _, _ := parser.ParseUnverified(rp[0], &claims)
			header, _ := json.MarshalIndent(IDToken.Header, "", "  ")
			payload, _ := json.MarshalIndent(claims, "", "  ")

			data["request_raw"] = rp[0]
			data["request_header"] = string(header)
			data["request_payload"] = string(payload)
		}

		if s.DoRequestObjectEncryption() {
			if loginURL, storage.CSRF, err = s.BuildLoginURL(c, consentID, true); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
				return
			}
		}

		// persist csrf in a secure encrypted cookie
		if encodedCookieValue, err = s.SecureCookie.Encode("app", storage); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error while encoding cookie: %+v", err))
			return
		}

		c.SetCookie("app", encodedCookieValue, 0, "/", "", false, true)

		data["intent_id"] = consentID
		data["login_url"] = loginURL

		c.HTML(http.StatusOK, s.GetTemplate("intent_registered.tmpl"), data)
	}
}

func (s *Server) Callback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app              string
			appStorage       = AppStorage{}
			userinfoResponse map[string]interface{}
			additionalData   map[string]interface{}
			code             = c.Query("code")
			token            acpclient.Token
			data             = gin.H{}
			err              error
		)

		if c.Query("error") != "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %s: %s", c.Query("error"), c.Query("error_description")))
			return
		}

		if app, err = c.Cookie("app"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get app cookie: %+v", err))
			return
		}

		if err = s.SecureCookie.Decode("app", app, &appStorage); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode app storage: %+v", err))
			return
		}

		if token, err = s.Client.Exchange(code, c.Query("state"), appStorage.CSRF); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if userinfoResponse, err = s.Client.Userinfo(token.AccessToken); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to introspect access token: %+v", err))
			return
		}

		if token.GrantID != "" {
			data["grant_id"] = token.GrantID
		}

		tokenResponse, _ := json.MarshalIndent(token, "", "  ")
		data["token_response"] = string(tokenResponse)

		data["access_token"] = token.AccessToken

		parser := jwt.Parser{}

		atClaims := jwt.MapClaims{}
		if _, _, err = parser.ParseUnverified(token.AccessToken, &atClaims); err == nil {
			payload, _ := json.MarshalIndent(atClaims, "", "  ")
			data["access_token_payload"] = string(payload)
		}

		userinfoResp, _ := json.MarshalIndent(userinfoResponse, "", "  ")
		data["userinfo"] = string(userinfoResp)

		if token.IDToken != "" {
			claims := jwt.MapClaims{}
			IDToken, _, _ := parser.ParseUnverified(token.IDToken, &claims)
			header, _ := json.MarshalIndent(IDToken.Header, "", "  ")
			payload, _ := json.MarshalIndent(claims, "", "  ")

			data["id_token_raw"] = token.IDToken
			data["id_token_header"] = string(header)
			data["id_token_payload"] = string(payload)
		}

		var accountsResp interface{}

		if accountsResp, err = s.GetAccounts(c, token.AccessToken); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
			return
		}

		accountsRaw, _ := json.MarshalIndent(accountsResp, "", "  ")
		data["accounts_raw"] = string(accountsRaw)

		if additionalData, err = s.PostAuthenticationAction(c, data); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to execute post authentication action: %+v", err))
			return
		}

		for k, v := range additionalData {
			data[k] = v
		}

		c.HTML(http.StatusOK, s.GetTemplate("authenticated.tmpl"), data)
	}
}
