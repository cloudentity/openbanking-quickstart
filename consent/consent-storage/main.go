package main

import (
	"fmt"
	"strconv"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	RootCA          string `env:"ROOT_CA" envDefault:"/ca.pem"`
	CertFile        string `env:"CERT_FILE" envDefault:"/bank_cert.pem"`
	KeyFile         string `env:"KEY_FILE" envDefault:"/bank_key.pem"`
	Port            int    `env:"PORT" envDefault:"8084"`
	EnableTLSServer bool   `env:"ENABLE_TLS_SERVER" envDefault:"true"`
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config Config
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	r.POST("/consents", createConsent)
	r.GET("/consents", listConsents)

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
