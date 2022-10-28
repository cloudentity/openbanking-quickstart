package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Spec string

const (
	OBUK Spec = "obuk"
	CDR  Spec = "cdr"
	OBBR Spec = "obbr"
	FDX  Spec = "fdx"
)

type Config struct {
	SystemClientID         string        `env:"SYSTEM_CLIENT_ID" envDefault:"bv2fe0tpfc67lmeti340"`
	SystemClientSecret     string        `env:"SYSTEM_CLIENT_SECRET" envDefault:"v6yYaApda9juR_DrlI5mpdsdm9u2-D0rQIG9ynakyDE"`
	SystemIssuerURL        *url.URL      `env:"SYSTEM_ISSUER_URL,required"`
	OpenbankingWorkspaceID string        `env:"OPENBANKING_WORKSPACE_ID,required"`
	Timeout                time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA                 string        `env:"ROOT_CA" envDefault:"/ca.pem"`
	CertFile               string        `env:"CERT_FILE" envDefault:"/bank_cert.pem"`
	KeyFile                string        `env:"KEY_FILE" envDefault:"/bank_key.pem"`
	Port                   int           `env:"PORT" envDefault:"8086"`
	IntrospectClientID     string        `env:"INTROSPECT_CLIENT_ID" envDefault:"bv2fddlpfc67lmeti32g"`
	IntrospectClientSecret string        `env:"INTROSPECT_CLIENT_SECRET" envDefault:"RkAIOVRlP6ZLhDvYixT1wBb8DG3bVkXAouoIrX7cwzA"`
	IntrospectIssuerURL    *url.URL      `env:"INTROSPECT_ISSUER_URL,required"`
	Spec                   Spec          `env:"SPEC,required"`
}

func (c *Config) SystemClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:     c.SystemClientID,
		ClientSecret: c.SystemClientSecret,
		IssuerURL:    c.SystemIssuerURL,
		Scopes:       []string{"manage_openbanking_consents", "view_clients"},
		Timeout:      c.Timeout,
		CertFile:     c.CertFile,
		KeyFile:      c.KeyFile,
		RootCA:       c.RootCA,
	}
}

func (c *Config) IntrospectClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:     c.IntrospectClientID,
		ClientSecret: c.IntrospectClientSecret,
		IssuerURL:    c.IntrospectIssuerURL,
		Scopes:       []string{"introspect_tokens"},
		Timeout:      c.Timeout,
		CertFile:     c.CertFile,
		KeyFile:      c.KeyFile,
		RootCA:       c.RootCA,
	}
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config        Config
	Client        acpclient.Client
	ConsentClient ConsentFetchRevoker
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.Client, err = acpclient.New(server.Config.SystemClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.ConsentClient, err = ConsentFetcherFactory(server.Config.Spec, &server); err != nil {
		return server, errors.Wrapf(err, "failed to create server consent client")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())

	r.GET("/clients", s.ListClients())
	r.DELETE("/consents/:id", s.RevokeConsent())
	r.DELETE("/clients/:id", s.RevokeConsentsForClient())

	r.NoRoute(func(c *gin.Context) {
		c.File("web/app/build/index.html")
	})

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
