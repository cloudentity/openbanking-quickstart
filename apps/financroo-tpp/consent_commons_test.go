package main

import (
	"net/http/httptest"
	"testing"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

func TestCreateConsentResponse(t *testing.T) {
	var (
		s         Server
		ctx       *gin.Context
		bankID    BankID
		user      User
		acpClient acpclient.Client
		consentID string
	)
	ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	bankID = "test_bank_id"
	consentID = "test_consent_id"
	s.Clients.ConsentClient = &FakeConsentClient{}
	s.Clients.AcpAccountsClient = acpClient
	s.LoginURLBuilder = &FakeLoginBuilder{}
	s.SecureCookie = securecookie.New(securecookie.GenerateRandomKey(32), nil)
	s.SecureCookie.Encode("app", AppStorage{
		BankID:   bankID,
		Sub:      user.Sub,
		IntentID: consentID,
	})

	s.CreateConsentResponse(ctx, bankID, user, s.Clients.AcpAccountsClient, consentID)
}

type FakeConsentClient struct {
}

func (f *FakeConsentClient) CreateConsentExplicitly() bool {
	return true
}

func (f *FakeConsentClient) CreateAccountConsent(c *gin.Context) (string, error) {
	return "", nil
}

func (f *FakeConsentClient) CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error) {
	return "", nil
}

func (f *FakeConsentClient) GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error) {
	return nil, nil
}

func (f *FakeConsentClient) UsePAR() bool {
	return true
}

func (f *FakeConsentClient) DoPAR(c *gin.Context) (string, acpclient.CSRF, error) {
	return "", acpclient.CSRF{}, nil
}

func (f *FakeConsentClient) Sign([]byte) (string, error) {
	return "", nil
}

var _ ConsentClient = &FakeConsentClient{}

type FakeLoginBuilder struct {
}

func (f *FakeLoginBuilder) BuildLoginURL(string, acpclient.Client) (string, acpclient.CSRF, error) {
	return "", acpclient.CSRF{}, nil
}
