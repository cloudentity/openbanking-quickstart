package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	logrus "github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Spec string

const (
	OBUK Spec = "obuk"
	OBBR Spec = "obbr"
)

type Config struct {
	Port      int           `env:"PORT" envDefault:"8070"`
	ClientID  string        `env:"CLIENT_ID,required"`
	IssuerURL *url.URL      `env:"ISSUER_URL,required"`
	Timeout   time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA    string        `env:"ROOT_CA,required"`
	CertFile  string        `env:"CERT_FILE,required"`
	KeyFile   string        `env:"KEY_FILE,required"`
	Spec      Spec          `env:"SPEC,required"`
}

func (c *Config) ClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:  c.ClientID,
		IssuerURL: c.IssuerURL,
		Scopes:    []string{"introspect_openbanking_tokens"},
		Timeout:   c.Timeout,
		CertFile:  c.CertFile,
		KeyFile:   c.KeyFile,
		RootCA:    c.RootCA,
	}
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config  Config
	Client  acpclient.Client
	Storage Storage
	// PaymentQueue PaymentQueue

	GetAccountsLogic         GetEndpointLogic
	GetInternalAccountsLogic GetEndpointLogic

	GetBalancesLogic         GetEndpointLogic
	GetBalancesInternalLogic GetEndpointLogic

	GetTransactionsLogic GetEndpointLogic

	CreatePaymentLogic CreateEndpointLogic
	GetPaymentLogic    GetEndpointLogic
}

// TODO: fix error mapping so that http code is returned
func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.Client, err = acpclient.New(server.Config.ClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.Storage, err = NewUserRepo(); err != nil {
		return server, errors.Wrapf(err, "failed to init repo")
	}

	switch server.Config.Spec {
	case OBUK:
		server.GetAccountsLogic = &OBUKGetAccountsHandler{Server: &server}
		server.GetInternalAccountsLogic = &OBUKGetAccountsInternalHandler{Server: &server}
		server.GetBalancesLogic = &OBUKGetBalancesHandler{Server: &server}
		server.GetBalancesInternalLogic = &OBUKGetBalancesInternalHandler{Server: &server}
		server.GetTransactionsLogic = &OBUKGetTransactionsHandler{Server: &server}
		server.CreatePaymentLogic = &OBUKCreatePaymentHandler{Server: &server}
		server.GetPaymentLogic = &OBUKGetPaymentHandler{Server: &server}
	case OBBR:
		server.GetAccountsLogic = &OBBRGetAccountsHandler{Server: &server}
	default:
		return server, errors.Wrapf(err, "unsupported spec %s", server.Config.Spec)
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.GET("/accounts", s.Get(s.GetAccountsLogic))
	r.GET("/internal/accounts/:sub", s.Get(s.GetInternalAccountsLogic))

	r.GET("/balances", s.Get(s.GetBalancesLogic))
	r.GET("/internal/balances/:sub", s.Get(s.GetBalancesInternalLogic))

	r.GET("/transactions", s.Get(s.GetTransactionsLogic))

	r.POST("/domestic-payments", s.Post(s.CreatePaymentLogic))
	r.GET("/domestic-payments/:DomesticPaymentId", s.Get(s.GetPaymentLogic))

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}

func main() {
	var (
		server Server
		err    error
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}
