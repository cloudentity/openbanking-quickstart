package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudentity/openbanking-quickstart/client/accounts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

type AppStorage struct {
	CSRF acpclient.CSRF
}

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	}
}

func (s *Server) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			registerResponse   *openbanking.CreateAccountAccessConsentRequestCreated
			encodedCookieValue string
			storage            AppStorage
			loginURL           string
			u                  *url.URL
			data               = gin.H{}
			err                error
		)

		if registerResponse, err = s.Client.Openbanking.CreateAccountAccessConsentRequest(
			openbanking.NewCreateAccountAccessConsentRequestParams().
				WithTid(s.Client.TenantID).
				WithAid(s.Client.ServerID).
				WithRequest(&models.AccountAccessConsentRequest{
					Data: &models.OBReadConsent1Data{
						Permissions: c.PostFormArray("permissions"),
					},
					Risk: map[string]interface{}{},
				}),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register account access consent: %+v", err))
			return
		}

		registerResponseRaw, _ := json.MarshalIndent(registerResponse, "", "  ")
		data["account_access_consent_raw"] = string(registerResponseRaw)

		if loginURL, storage.CSRF, err = s.Client.AuthorizeURL(
			acpclient.WithOpenbankingIntentID(*registerResponse.Payload.Data.ConsentID, []string{"urn:openbanking:psd2:sca"}),
			acpclient.WithPKCE(),
		); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
			return
		}

		if u, err = url.Parse(loginURL); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to parse login url: %+v", err))
			return
		}

		// persist csrf in a secure encrypted cookie
		if encodedCookieValue, err = s.SecureCookie.Encode("app", storage); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error while encoding cookie: %+v", err))
			return
		}

		c.SetCookie("app", encodedCookieValue, 0, "/", "", false, true)

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

		data["intent_id"] = registerResponse.Payload.Data.ConsentID
		data["login_url"] = loginURL

		c.HTML(http.StatusOK, "intent_registered.tmpl", data)
	}
}

func (s *Server) Callback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app              string
			appStorage       = AppStorage{}
			userinfoResponse map[string]interface{}
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

		data["access_token"] = token.AccessToken
		userinfoResp, _ := json.MarshalIndent(userinfoResponse, "", "  ")
		data["userinfo"] = string(userinfoResp)

		if token.IDToken != "" {
			parser := jwt.Parser{}
			claims := jwt.MapClaims{}
			IDToken, _, _ := parser.ParseUnverified(token.IDToken, &claims)
			header, _ := json.MarshalIndent(IDToken.Header, "", "  ")
			payload, _ := json.MarshalIndent(claims, "", "  ")

			data["id_token_raw"] = IDToken
			data["id_token_header"] = string(header)
			data["id_token_payload"] = string(payload)
		}

		var accountsResp *accounts.GetAccountsOK

		if accountsResp, err = s.BankClient.Accounts.GetAccounts(accounts.NewGetAccountsParams().WithAuthorization(token.AccessToken), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
			return
		}

		accountsRaw, _ := json.MarshalIndent(accountsResp.Payload, "", "  ")
		data["accounts_raw"] = string(accountsRaw)

		c.HTML(http.StatusOK, "authenticated.tmpl", data)
	}
}
