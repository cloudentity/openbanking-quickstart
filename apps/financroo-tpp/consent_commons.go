package main

import (
	"fmt"
	"net/http"
	"net/url"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateConsentResponse(c *gin.Context, bankID BankID, consentID string, user User, client acpclient.Client) {
	var (
		loginURL string
		err      error
		encodedCookieValue string
		app =    AppStorage{
			BankID:   bankID,
			IntentID: consentID,
			Sub:      user.Sub,
		}
		data = gin.H{}
	)

	if loginURL, app.CSRF, err = client.AuthorizeURL(
		acpclient.WithOpenbankingIntentID(app.IntentID, []string{"urn:openbanking:psd2:sca"}),
		acpclient.WithPKCE(),
	); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
		return
	}

	if _, err = url.Parse(loginURL); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to parse login url: %+v", err))
		return
	}

	// persist verifier and nonce in a secure encrypted cookie
	if encodedCookieValue, err = s.SecureCookie.Encode("app", app); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error while encoding cookie: %+v", err))
		return
	}

	c.SetCookie("app", encodedCookieValue, 0, "/", "", false, true)

	data["login_url"] = loginURL

	c.JSON(http.StatusOK, data)
}

