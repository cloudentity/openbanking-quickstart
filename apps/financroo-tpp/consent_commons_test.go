package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                      string
	fakeConsentClient         FakeConsentClient
	fakeLoginBuilder          FakeLoginBuilder
	securecookie              *securecookie.SecureCookie
	expectedDoPARCount        int
	expectedLoginBuilderCount int
	expectedStatusCode        int
	expectedError             string
}

func TestCreateConsentResponse(t *testing.T) {
	var (
		s         Server
		ctx       *gin.Context
		acpClient acpclient.Client
	)

	tcs := []testCase{
		{
			name: "Success with PAR",
			fakeConsentClient: FakeConsentClient{
				UsePARvar: false,
			},
			fakeLoginBuilder: FakeLoginBuilder{
				LoginURL: "https://test.login.url",
			},
			securecookie: securecookie.New(securecookie.GenerateRandomKey(32), nil),

			expectedDoPARCount:        0,
			expectedLoginBuilderCount: 1,
			expectedStatusCode:        http.StatusOK,
		},
		{
			name: "Success without PAR",
			fakeConsentClient: FakeConsentClient{
				UsePARvar: true,
			},
			fakeLoginBuilder: FakeLoginBuilder{
				LoginURL: "https://test.login.url",
			},
			securecookie: securecookie.New(securecookie.GenerateRandomKey(32), nil),

			expectedDoPARCount:        1,
			expectedLoginBuilderCount: 1,
			expectedStatusCode:        http.StatusOK,
		},
		{
			name: "DoPAR failure",
			fakeConsentClient: FakeConsentClient{
				UsePARvar:  true,
				DoPARError: errors.New("do PAR error"),
			},

			expectedDoPARCount:        1,
			expectedLoginBuilderCount: 0,
			expectedStatusCode:        http.StatusBadRequest,
			expectedError:             "failed to register PAR request: do PAR error",
		},
		{
			name: "DoPAR success LoginURL failure",
			fakeConsentClient: FakeConsentClient{
				UsePARvar: true,
			},
			fakeLoginBuilder: FakeLoginBuilder{
				Err: errors.New("login builder error"),
			},

			expectedDoPARCount:        1,
			expectedLoginBuilderCount: 1,
			expectedStatusCode:        http.StatusInternalServerError,
			expectedError:             "failed to build authorize url: login builder error",
		},
		{
			name: "URL parse failure",
			fakeLoginBuilder: FakeLoginBuilder{
				Err: errors.New("login builder error"),
			},

			expectedLoginBuilderCount: 1,
			expectedStatusCode:        http.StatusInternalServerError,
			expectedError:             "failed to build authorize url: login builder error",
		},
		{
			name: "Invalid login URL",
			fakeLoginBuilder: FakeLoginBuilder{
				LoginURL: "invalid-!@#$%^&&*()()_",
			},

			expectedLoginBuilderCount: 1,
			expectedStatusCode:        http.StatusInternalServerError,
			expectedError:             `failed to parse login url: parse "invalid-!@#$%^&&*()()_": invalid URL escape "%^&"`,
		},
		{
			name:         "Secure cookie Encode failure",
			securecookie: securecookie.New(nil, nil),

			expectedLoginBuilderCount: 1,
			expectedStatusCode:        http.StatusInternalServerError,
			expectedError:             "error while encoding cookie: securecookie: hash key is not set",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			recorder := httptest.NewRecorder()
			ctx, _ = gin.CreateTestContext(recorder)

			s.Clients.ConsentClient = &tc.fakeConsentClient
			s.Clients.AcpAccountsClient = acpClient
			s.LoginURLBuilder = &tc.fakeLoginBuilder
			s.SecureCookie = tc.securecookie

			s.CreateConsentResponse(ctx, "test_bank_id", User{}, s.Clients.AcpAccountsClient, "test_consent_id")
			response := recorder.Result()

			assert.Equal(tt, tc.expectedDoPARCount, tc.fakeConsentClient.DoPARCalls)
			assert.Equal(tt, tc.expectedLoginBuilderCount, tc.fakeLoginBuilder.BuildLoginURLCount)
			assert.Equal(tt, tc.expectedStatusCode, response.StatusCode)

			assertBody(tt, &tc, response)
		})
	}
}

func assertBody(t *testing.T, tc *testCase, response *http.Response) {
	var (
		body    []byte
		err     error
		bodyMap map[string]string
	)
	t.Helper()

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	if tc.expectedError != "" {
		assert.Equal(t, tc.expectedError, string(body))
	} else {
		err = json.Unmarshal(body, &bodyMap)

		assert.NoError(t, err)
		assert.Equal(t, bodyMap["login_url"], tc.fakeLoginBuilder.LoginURL)

		if _, ok := response.Header["Set-Cookie"]; ok {
			assert.Contains(t, response.Header["Set-Cookie"][0], "app=")
		} else {
			t.Fatalf("Set-Cookie header does not exists")
		}
	}
}

type FakeConsentClient struct {
	UsePARvar  bool
	DoPARCalls int
	DoPARError error
}

func (f *FakeConsentClient) CreateConsentExplicitly() bool {
	return false
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
	return f.UsePARvar
}

func (f *FakeConsentClient) DoPAR(c *gin.Context) (string, acpclient.CSRF, error) {
	f.DoPARCalls++
	return "test_request_uri",
		acpclient.CSRF{},
		f.DoPARError
}

func (f *FakeConsentClient) Sign([]byte) (string, error) {
	return "", nil
}

var _ ConsentClient = &FakeConsentClient{}

type FakeLoginBuilder struct {
	BuildLoginURLCount int
	LoginURL           string
	Err                error
}

func (f *FakeLoginBuilder) BuildLoginURL(string, acpclient.Client) (string, acpclient.CSRF, error) {
	f.BuildLoginURLCount++
	return f.LoginURL, acpclient.CSRF{}, f.Err
}
