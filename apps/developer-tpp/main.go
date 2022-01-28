package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"

	acpclient "github.com/cloudentity/acp-client-go"
	oauth2 "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
)

type Spec string

const (
	OBUK Spec = "obuk"
	OBBR Spec = "obbr"
)

type Config struct {
	Port         int           `env:"PORT" envDefault:"8090"`
	ClientID     string        `env:"CLIENT_ID,required"`
	AuthorizeURL *url.URL      `env:"AUTHORIZE_URL,required"`
	TokenURL     *url.URL      `env:"TOKEN_URL,required"`
	IssuerURL    *url.URL      `env:"ISSUER_URL,required"`
	UserinfoURL  *url.URL      `env:"USERINFO_URL,required"`
	RedirectURL  *url.URL      `env:"REDIRECT_URL,required"`
	Timeout      time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA       string        `env:"ROOT_CA,required"`
	CertFile     string        `env:"CERT_FILE,required"`
	KeyFile      string        `env:"KEY_FILE,required"`
	BankURL      *url.URL      `env:"BANK_URL,required"`
	Spec         Spec          `env:"SPEC,required"`
	ClientScopes []string
}

func (c Config) ExtendConsentScope(consentID string) *Config {
	for idx, scope := range c.ClientScopes {
		if strings.HasPrefix(scope, "consent:") {
			c.ClientScopes[idx] = "consent:" + consentID
			break
		}
	}
	return &c
}

func (c *Config) ClientConfig() acpclient.Config {
	requestObjectExpiration := time.Minute * 10
	return acpclient.Config{
		ClientID:                    c.ClientID,
		IssuerURL:                   c.IssuerURL,
		TokenURL:                    c.TokenURL,
		AuthorizeURL:                c.AuthorizeURL,
		UserinfoURL:                 c.UserinfoURL,
		RedirectURL:                 c.RedirectURL,
		RequestObjectSigningKeyFile: c.KeyFile,
		RequestObjectExpiration:     &requestObjectExpiration,
		Scopes:                      c.ClientScopes,
		Timeout:                     c.Timeout,
		CertFile:                    c.CertFile,
		KeyFile:                     c.KeyFile,
		RootCA:                      c.RootCA,
	}
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config       Config
	Client       acpclient.Client
	BankClient   OpenbankingClient
	SecureCookie *securecookie.SecureCookie
	SpecLogicHandler
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	switch server.Config.Spec {
	case OBUK:
		server.Config.ClientScopes = []string{"openid", "accounts"}
	case OBBR:
		server.Config.ClientScopes = []string{"openid", "consents", "consent:*"}
	}

	if server.Client, err = acpclient.New(server.Config.ClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	server.BankClient = NewOpenbankingClient(server.Config)

	server.SecureCookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

	switch server.Config.Spec {
	case OBUK:
		server.SpecLogicHandler = &OBUKLogic{Server: &server}
	case OBBR:
		server.SpecLogicHandler = &OBBRLogic{Server: &server}
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*/*")
	r.Static("/assets", "./assets")

	r.GET("/", s.Get())
	r.POST("/login", s.Login())
	r.GET("/callback", s.Callback())

	return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
}

func main() {
	var (
		server Server
		err    error
	)

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}

func (s *Server) GetTemplate(name string) string {
	switch s.Config.Spec {
	case OBUK:
		return string(OBUK) + "-" + name
	case OBBR:
		return string(OBBR) + "-" + name
	default:
		return ""
	}
}

func (s *Server) GetEncryptionKey(ctx context.Context) (jose.JSONWebKey, error) {
	var (
		jwksResponse *oauth2.JwksOK
		encKey       jose.JSONWebKey
		b            []byte
		err          error
	)

	if jwksResponse, err = s.Client.Oauth2.Oauth2.Jwks(
		oauth2.NewJwksParamsWithContext(ctx)); err != nil {
		return encKey, errors.Wrapf(err, "failed to get jwks from acp server")
	}

	for _, key := range jwksResponse.Payload.Keys {
		if key.Use == "enc" {
			if b, err = json.Marshal(key); err != nil {
				return encKey, errors.Wrapf(err, "failed to marshal key")
			}

			if err = encKey.UnmarshalJSON(b); err != nil {
				return encKey, errors.Wrapf(err, "failed to unmarshal jwk")
			}

			break
		}
	}

	return encKey, nil
}
