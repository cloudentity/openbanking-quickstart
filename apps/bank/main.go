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
	OBUK    Spec = "obuk"
	OBBR    Spec = "obbr"
	CDR     Spec = "cdr"
	FDX     Spec = "fdx"
	Generic Spec = "generic"
)

type Config struct {
	Port                int           `env:"PORT" envDefault:"8070"`
	ClientID            string        `env:"CLIENT_ID" envDefault:"bukj5p6k7qdmm5ppbi4g"`
	IssuerURL           *url.URL      `env:"ISSUER_URL,required"`
	Timeout             time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA              string        `env:"ROOT_CA" envDefault:"/ca.pem"`
	CertFile            string        `env:"CERT_FILE" envDefault:"/bank_cert.pem"`
	KeyFile             string        `env:"KEY_FILE" envDefault:"/bank_key.pem"`
	Spec                Spec          `env:"SPEC,required"`
	GINMODE             string        `env:"GIN_MODE"`
	UserIdentifierClaim string        `env:"USER_IDENTIFIER_CLAIM" envDefault:"sub"`
	SeedFilePath        string
}

func (c *Config) ClientConfig(scopes []string) acpclient.Config {
	return acpclient.Config{
		ClientID:  c.ClientID,
		IssuerURL: c.IssuerURL,
		Scopes:    scopes,
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

	switch config.Spec {
	case OBUK:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", OBUK)
	case OBBR:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", OBBR)
	case CDR:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", CDR)
	case FDX:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", FDX)
	case Generic:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", Generic)
	default:
		return config, fmt.Errorf("unknown spec: %s", config.Spec)
	}

	if config.GINMODE == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return config, err
}

type Server struct {
	Config  Config
	Client  acpclient.Client
	Storage Storage
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
	case Generic:
		if server.Client, err = acpclient.New(server.Config.ClientConfig([]string{"introspect_tokens"})); err != nil {
			return server, errors.Wrapf(err, "failed to init acp client")
		}
	default:
		if server.Client, err = acpclient.New(server.Config.ClientConfig([]string{"introspect_openbanking_tokens"})); err != nil {
			return server, errors.Wrapf(err, "failed to init acp client")
		}
	}

	if server.Storage, err = NewUserRepo(server.Config.SeedFilePath); err != nil {
		return server, errors.Wrapf(err, "failed to init repo")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	switch s.Config.Spec {
	case OBUK:
		r.GET("/accounts", s.Get(NewOBUKGetAccountsHandler))
		r.GET("/internal/accounts", s.Get(NewOBUKGetAccountsInternalHandler))
		r.GET("/balances", s.Get(NewOBUKGetBalancesHandler))
		r.GET("/internal/balances", s.Get(NewOBUKGetBalancesInternalHandler))
		r.GET("/transactions", s.Get(NewOBUKGetTransactionsHandler))
		r.POST("/domestic-payments", s.Post(NewOBUKCreatePaymentHandler))
		r.GET("/domestic-payments/:DomesticPaymentId", s.Get(NewOBUKGetPaymentHandler))

	case OBBR:
		r.GET("/accounts/v1/accounts", s.Get(NewOBBRGetAccountsHandler))
		r.GET("/internal/accounts", s.Get(NewOBBRGetAccountsInternalHandler))
		r.GET("/accounts/v1/accounts/:accountID/balances", s.Get(NewOBBRGetBalanceHandler))
		r.GET("/internal/balances", s.Get(NewOBBRGetBalancesInternalHandler))
		r.POST("/payments/v1/pix/payments", s.Post(NewOBBRCreatePaymentHandler))

	case CDR:
		r.POST("/internal/accounts", s.Get(NewCDRGetAccountsInternalHandler))
		r.GET("/banking/accounts", s.Get(NewCDRGetAccountsHandler))
		r.GET("/banking/accounts/:accountId/transactions", s.Get(NewCDRGetTransactionsHandler))
		r.GET("/banking/accounts/balances", s.Get(NewCDRGetBalancesHandler))

	case FDX:
		r.GET("/accounts", s.Get(NewFDXGetAccountsHandler))
		r.GET("/internal/accounts", s.Get(NewFDXGetAccountsInternalHandler))
		r.GET("/accounts/:accountId", s.Get(NewFDXGetBalancesHandler))
		r.GET("/accounts/:accountId/transactions", s.Get(NewFDXGetTransactionsHandler))

	case Generic:
		r.POST("/internal/accounts", s.Get(NewGenericGetAccountsInternalHandler))
		r.GET("/banking/accounts", s.Get(NewGenericGetAccountsHandler))
		r.GET("/banking/accounts/:accountId/transactions", s.Get(NewGenericGetTransactionsHandler))
		r.GET("/banking/accounts/balances", s.Get(NewGenericGetBalancesHandler))

	default:
		return fmt.Errorf("unsupported spec %s", s.Config.Spec)
	}

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
