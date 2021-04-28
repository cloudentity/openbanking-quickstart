package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Config struct {
	Port             int           `env:"PORT" envDefault:"8080"`
	ClientID         string        `env:"CLIENT_ID,required"`
	ClientSecret     string        `env:"CLIENT_SECRET,required"`
	IssuerURL        *url.URL      `env:"ISSUER_URL,required"`
	Timeout          time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA           string        `env:"ROOT_CA"`
	CertFile         string        `env:"CERT_FILE,required"`
	KeyFile          string        `env:"KEY_FILE,required"`
	BankURL          *url.URL      `env:"BANK_URL"`
	EnableMFAOTP     bool          `env:"ENABLE_MFA_OTP"`
	EnableMFAOkta    bool          `env:"ENABLE_MFA_OKTA"`
	OktaHost         string        `env:"OKTA_HOST"`
	OktaAPIToken     string        `env:"OKTA_API_TOKEN"`
	OktaUser         string        `env:"OKTA_USER"`
	OTPMode          string        `env:"OTP_MODE"` // optional, set to "mock" to use "111111" as otp
	TwilioAccountSid string        `env:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken  string        `env:"TWILIO_AUTH_TOKEN"`
	TwilioFrom       string        `env:"TWILIO_FROM" envDefault:"Cloudentity"`
	DBFile           string        `env:"DB_FILE" envDefault:"./data/my.db"`
	MobileClaim      string        `env:"MOBILE_CLAIM" envDefault:"mobile_verified"`
	LogLevel         string        `env:"LOG_LEVEL" envDefault:"info"`
}

func (c *Config) ClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		IssuerURL:    c.IssuerURL,
		Scopes:       []string{"manage_openbanking_consents"},
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
	Config      Config
	Client      acpclient.Client
	BankClient  BankClient
	SMSClient   *SMSClient
	OTPRepo     *OTPRepo
	OTPHandler  OTPHandler
	OktaHandler OktaHandler
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		db     *bolt.DB
		l      logrus.Level
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if l, err = logrus.ParseLevel(server.Config.LogLevel); err != nil {
		l = logrus.InfoLevel
	}
	logrus.SetLevel(l)

	if server.Client, err = acpclient.New(server.Config.ClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	server.SMSClient = NewSMSClient(server.Config)

	server.BankClient = NewBankClient(server.Config)

	if db, err = InitDB(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init db")
	}

	if server.OTPRepo, err = NewOTPRepo(db); err != nil {
		return server, errors.Wrapf(err, "failed to init otp repo")
	}

	server.OTPHandler = NewOTPHandler(server.Config.OTPMode, server.OTPRepo, server.SMSClient)

	logrus.Debugf("server config: %+v", server.Config)
	if server.Config.EnableMFAOkta {
		server.OktaHandler = NewOktaHandler(server.Config.OktaHost, server.Config.OktaAPIToken)
	}

	return server, nil
}

var mfaPath = "/mfa"

func RequireMFAMiddleware(s *Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, mfaPath) {
			c.Next()

			return
		}
		var (
			approved        bool
			isApprovedFuncs []func(LoginRequest) (bool, error)
			err             error
		)

		if s.Config.EnableMFAOTP {
			isApprovedFuncs = append(isApprovedFuncs, s.OTPHandler.IsApproved)
		}
		if s.Config.EnableMFAOkta {
			isApprovedFuncs = append(isApprovedFuncs, s.OktaHandler.IsApproved)
		}

		for _, f := range isApprovedFuncs {
			if approved, err = f(NewLoginRequest(c)); err != nil {
				RenderInvalidRequestError(c, nil)
				c.Abort()
				return
			}
			if approved {
				break
			}
		}

		if !approved {
			redirect := fmt.Sprintf("%s?%s", mfaPath, c.Request.URL.Query().Encode())

			c.Redirect(http.StatusTemporaryRedirect, redirect)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	if s.Config.EnableMFAOTP || s.Config.EnableMFAOkta {
		r.Use(RequireMFAMiddleware(s))
		r.GET(mfaPath, s.MFAHandler())
		r.POST(mfaPath, s.MFAHandler())
	}

	r.GET("/", s.Get())
	r.POST("/", s.Post())

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
