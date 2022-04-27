package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/text/language"

	acpclient "github.com/cloudentity/acp-client-go"
)

// TODO: configure customer_id in acp cdr workspace as a user attribute

type Spec string

const (
	OBUK Spec = "obuk"
	OBBR Spec = "obbr"
	CDR  Spec = "cdr"
	FDX  Spec = "fdx"
)

type Config struct {
	Port             int           `env:"PORT" envDefault:"8080"`
	ClientID         string        `env:"CLIENT_ID" envDefault:"bv0ocudfotn6edhsiu7g"`
	ClientSecret     string        `env:"CLIENT_SECRET" envDefault:"pMPBmv62z3Jt1S4sWl2qRhOhEGPVZ9EcujGL7Xy0-E0"`
	IssuerURL        *url.URL      `env:"ISSUER_URL,required"`
	Timeout          time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA           string        `env:"ROOT_CA" envDefault:"/ca.pem"`
	CertFile         string        `env:"CERT_FILE" envDefault:"/bank_cert.pem"`
	KeyFile          string        `env:"KEY_FILE" envDefault:"/bank_key.pem"`
	BankIDClaim      string        `env:"BANK_ID_CLAIM" envDefault:"sub"`
	EnableMFA        bool          `env:"ENABLE_MFA"`
	OTPMode          string        `env:"OTP_MODE" envDefault:"demo"`
	TwilioAccountSid string        `env:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken  string        `env:"TWILIO_AUTH_TOKEN"`
	TwilioFrom       string        `env:"TWILIO_FROM" envDefault:"Cloudentity"`
	DBFile           string        `env:"DB_FILE" envDefault:"/data/my.db"`
	MFAClaim         string        `env:"MFA_CLAIM" envDefault:"mobile_verified"`
	LogLevel         string        `env:"LOG_LEVEL" envDefault:"info"`
	DevMode          bool          `env:"DEV_MODE"`
	DefaultLanguage  language.Tag  `env:"DEFAULT_LANGUAGE"  envDefault:"en-us"`
	TransDir         string        `env:"TRANS_DIR" envDefault:"./translations"`
	Spec             Spec          `env:"SPEC,required"`

	Otp OtpConfig

	BankClientConfig BankClientConfig
}

type OtpConfig struct {
	Type       string        `env:"OTP_TYPE" envDefault:"otp"`
	RequestURL string        `env:"OTP_REQUEST_URL"`
	VerifyURL  string        `env:"OTP_VERIFY_URL"`
	Timeout    time.Duration `env:"OTP_TIMEOUT" envDefault:"10s"`
	AuthHeader string        `env:"OTP_AUTH_HEADER"`
}

type BankClientConfig struct {
	URL          *url.URL `env:"BANK_URL"`
	CertFile     string   `env:"BANK_CLIENT_CERT_FILE" envDefault:"/tpp_cert.pem"`
	KeyFile      string   `env:"BANK_CLIENT_KEY_FILE" envDefault:"/tpp_key.pem"`
	TokenURL     string   `env:"BANK_CLIENT_TOKEN_URL"`
	ClientID     string   `env:"BANK_CLIENT_ID"`
	ClientSecret string   `env:"BANK_CLIENT_SECRET"`
	Scopes       []string `env:"BANK_CLIENT_SCOPES"`
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

	logrus.WithField("config", config).Debug("loaded config")

	return config, err
}

type Server struct {
	Config                          Config
	Client                          acpclient.Client
	BankClient                      BankClient
	SMSClient                       *SMSClient
	OTPRepo                         *OTPRepo
	OTPHandler                      OTPHandler
	Trans                           *Trans
	PaymentConsentHandler           ConsentHandler
	PaymentMFAConsentProvider       MFAConsentProvider
	AccountAccessConsentHandler     ConsentHandler
	AccountAccessMFAConsentProvider MFAConsentProvider
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		db     *bolt.DB
		l      logrus.Level
		err    error
		trans  []fs.FileInfo
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

	bundle := i18n.NewBundle(server.Config.DefaultLanguage)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	if trans, err = ioutil.ReadDir(server.Config.TransDir); err != nil {
		return server, errors.Wrapf(err, "failed to read dir %s", server.Config.TransDir)
	}

	for _, t := range trans {
		if _, err = bundle.LoadMessageFile(server.Config.TransDir + "/" + t.Name()); err != nil {
			return server, err
		}
	}

	server.Trans = NewTranslations(bundle, server.Config.DefaultLanguage.String())

	server.SMSClient = NewSMSClient(server.Config)

	switch server.Config.Spec {
	case OBUK:
		server.BankClient = NewOBUKBankClient(server.Config)
	case OBBR:
		server.BankClient = NewOBBRBankClient(server.Config)
	case CDR:
		server.BankClient = NewCDRBankClient(server.Config)
	case FDX:
		server.BankClient = NewFDXClient(server.Config)
	default:
		return Server{}, errors.New("invalid SPEC configuration")
	}

	if db, err = InitDB(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init db")
	}

	if server.OTPRepo, err = NewOTPRepo(db); err != nil {
		return server, errors.Wrapf(err, "failed to init otp repo")
	}

	if server.OTPHandler, err = NewOTPHandler(server.Config, server.OTPRepo, server.SMSClient); err != nil {
		return server, errors.Wrapf(err, "failed to init otp handler")
	}

	tools := ConsentTools{Trans: server.Trans, Config: server.Config}
	switch server.Config.Spec {
	case OBUK:
		server.AccountAccessConsentHandler = &OBUKAccountAccessConsentHandler{&server, tools}
		server.AccountAccessMFAConsentProvider = &OBUKAccountAccessMFAConsentProvider{&server, tools}
		server.PaymentConsentHandler = &OBUKDomesticPaymentConsentHandler{&server, tools}
		server.PaymentMFAConsentProvider = &DomesticPaymentMFAConsentProvider{&server, tools}
	case OBBR:
		server.AccountAccessConsentHandler = &OBBRAccountAccessConsentHandler{&server, tools}
		server.AccountAccessMFAConsentProvider = &OBBRAccountAccessMFAConsentProvider{&server, tools}
		server.PaymentConsentHandler = &OBBRPaymentConsentHandler{&server, tools}
		server.PaymentMFAConsentProvider = &OBBRPaymentMFAConsentProvider{&server, tools}
	case CDR:
		server.AccountAccessConsentHandler = &CDRAccountAccessConsentHandler{&server, tools}
	case FDX:
		server.AccountAccessConsentHandler = &FDXAccountAccessConsentHandler{&server, tools}
		server.AccountAccessMFAConsentProvider = &FDXAccountAccessMFAConsentProvider{&server, tools}
	default:
		return server, errors.Wrapf(err, "unsupported spec %s", server.Config.Spec)
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
			approved bool
			err      error
		)

		if approved, err = s.OTPHandler.IsApproved(NewLoginRequest(c)); err != nil {
			RenderInvalidRequestError(c, s.Trans, nil)
			c.Abort()
			return
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
	var err error

	r := gin.Default()

	if err = loadTemplates(r, "./templates"); err != nil {
		return errors.Wrapf(err, "failed to load templates")
	}

	r.Static("/assets", "./assets")

	base := r.Group("")

	if s.Config.EnableMFA {
		base.Use(RequireMFAMiddleware(s))
		base.GET(mfaPath, s.MFAHandler())
		base.POST(mfaPath, s.MFAHandler())
	}

	if s.Config.DevMode {
		demo := r.Group("/demo")
		demo.POST("/totp/verify", s.DemoTotpVerify)
	}

	base.GET("/", s.Get())
	base.POST("/", s.Post())

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

func loadTemplates(r *gin.Engine, dir string) error {
	var (
		baseTemplates   []string
		customTemplates []string
		err             error
	)

	if baseTemplates, err = filepath.Glob(fmt.Sprintf("%s/base/*.tmpl", dir)); err != nil {
		return errors.Wrapf(err, "failed to get base templates")
	}

	if customTemplates, err = filepath.Glob(fmt.Sprintf("%s/custom/*.tmpl", dir)); err != nil {
		return errors.Wrapf(err, "failed to get custom templates")
	}

	templates := append(baseTemplates, customTemplates...)
	r.LoadHTMLFiles(templates...)

	return nil
}
