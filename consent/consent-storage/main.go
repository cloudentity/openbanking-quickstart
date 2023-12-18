package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/cloudentity/openbanking-quickstart/shared"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

var Validator = validator.New()

type Config struct {
	RootCA          string   `env:"ROOT_CA"           envDefault:"/ca.pem"`
	CertFile        string   `env:"CERT_FILE"         envDefault:"/consent_storage_cert.pem"`
	KeyFile         string   `env:"KEY_FILE"          envDefault:"/consent_storage_key.pem"`
	Port            int      `env:"PORT"              envDefault:"8084"`
	EnableTLSServer bool     `env:"ENABLE_TLS_SERVER" envDefault:"true"`
	DBFile          string   `env:"DB_FILE"           envDefault:"/app/data/my.db"`
	IssuerURL       *url.URL `env:"ISSUER_URL"        validate:"required,url"`
	ServerID        string   `env:"SERVER"            validate:"required"`
	ClientID        string   `env:"CLIENT_ID"         validate:"required"`
	ClientSecret    string   `env:"CLIENT_SECRET"     validate:"required"`
}

func (c *Config) ClientConfig(scopes []string) acpclient.Config {
	return acpclient.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		IssuerURL:    c.IssuerURL,
		Scopes:       scopes,
		Timeout:      time.Second * 5,
		CertFile:     c.CertFile,
		KeyFile:      c.KeyFile,
		RootCA:       c.RootCA,
	}
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	if err = Validator.Struct(config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config Config
	DB     shared.DB
	Repo   ConsentRepo
	Client acpclient.Client
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	logrus.WithField("config", server.Config).Infof("config")

	if server.Client, err = acpclient.New(server.Config.ClientConfig([]string{"revoke_tokens"})); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.DB, err = shared.InitDB(server.Config.DBFile); err != nil {
		return server, errors.Wrapf(err, "failed to init db")
	}

	if server.Repo, err = NewConsentRepo(server.DB); err != nil {
		return server, errors.Wrapf(err, "failed to init consent repo")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "./assets")

	r.POST("/consents", s.CreateConsent)
	r.GET("/", s.ListConsentsHTML)
	r.GET("/:id", s.DeleteConsentHTML)

	if s.Config.EnableTLSServer {
		logrus.Debugf("running consent storage server tls")
		return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
	}

	logrus.Debugf("running consent storage server non-tls")
	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
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
