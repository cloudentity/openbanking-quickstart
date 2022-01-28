package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Config struct {
	ClientID     string        `env:"CLIENT_ID,required"`
	ClientSecret string        `env:"CLIENT_SECRET,required"`
	IssuerURL    string        `env:"ISSUER_URL,required"`
	CertFile     string        `env:"CERT_FILE"`
	KeyFile      string        `env:"KEY_FILE"`
	RootCA       string        `env:"ROOT_CA"`
	FailureURL   string        `env:"FAILURE_URL"`
	LogLevel     string        `env:"LOG_LEVEL" envDefault:"info"`
	Port         int           `env:"PORT" envDefault:"8080"`
	Timeout      time.Duration `env:"TIMEOUT" envDefault:"5s"`

	OIDC OidcConfig
}

func (c *Config) AcpClientConfig() (acpclient.Config, error) {
	issuerURL, err := url.Parse(c.IssuerURL)
	if err != nil {
		return acpclient.Config{}, err
	}

	return acpclient.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		IssuerURL:    issuerURL,
		Scopes:       []string{},
		Timeout:      c.Timeout,
		CertFile:     c.CertFile,
		KeyFile:      c.KeyFile,
		RootCA:       c.RootCA,
	}, nil
}

func LoadConfig() (config Config, err error) {
	var level logrus.Level

	if err = env.Parse(&config); err != nil {
		return config, err
	}
	// Anything that may come via a Kubernetes secret, should be trimmed of trailing linefeeds.
	// Values in K8s secrets must be base64-encoded, and linefeeds easily creep in.
	config.ClientID = strings.TrimSuffix(config.ClientID, "\n")
	config.ClientSecret = strings.TrimSuffix(config.ClientSecret, "\n")

	if level, err = logrus.ParseLevel(config.LogLevel); err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	// Log the config with an obscured client secret.
	cf := config
	cf.ClientSecret = cf.ClientSecret[0:4] + "..."
	cf.OIDC.ClientSecret = cf.OIDC.ClientSecret[0:4] + "..."
	logrus.WithField("config", cf).Debug("loaded config")

	return config, err
}

// This interface lets us substitute a mock http.Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Server struct {
	Config     Config
	AcpClient  acpclient.Client
	HTTPClient HTTPClient
	OidcClient OidcClient
}

func NewServer() (Server, error) {
	var (
		acpConfig acpclient.Config
		err       error
		server    = Server{HTTPClient: http.DefaultClient}
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if acpConfig, err = server.Config.AcpClientConfig(); err != nil {
		return server, errors.Wrapf(err, "invalid acp client config")
	}

	if server.AcpClient, err = acpclient.New(acpConfig); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.OidcClient, err = server.Config.OIDC.NewClient(); err != nil {
		return server, errors.Wrapf(err, "failed to init oidc client")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.New()
	r.Use(gin.Recovery())

	base := r.Group("")
	base.GET("/", s.Alive)
	base.GET("/alive", s.Alive)
	base.GET("/health", s.Alive)
	base.GET("/login", BindInput(LoginRequestInput{}), s.Login)
	base.GET("/callback", BindInput(CallbackInput{}), s.Callback)

	if s.Config.CertFile != "" && s.Config.KeyFile != "" {
		return r.RunTLS(fmt.Sprintf(":%d", s.Config.Port), s.Config.CertFile, s.Config.KeyFile)
	}
	return r.Run(fmt.Sprintf(":%d", s.Config.Port))
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
