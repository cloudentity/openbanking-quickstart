package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/openbanking-quickstart/shared"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Server struct {
	Config                          Config
	Client                          acpclient.Client
	BankClient                      BankClient
	SMSClient                       *SMSClient
	OTPRepo                         *OTPRepo
	OTPHandler                      OTPHandler
	MFAApprovalChecker              MFAApprovalChecker
	Trans                           *Trans
	PaymentConsentHandler           ConsentHandler
	PaymentMFAConsentProvider       MFAConsentProvider
	AccountAccessConsentHandler     ConsentHandler
	AccountAccessMFAConsentProvider MFAConsentProvider
	MFAStrategy                     MFAStrategy
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		db     shared.DB
		l      logrus.Level
		trans  []fs.DirEntry
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if l, err = logrus.ParseLevel(server.Config.LogLevel); err != nil {
		l = logrus.InfoLevel
	}
	logrus.SetLevel(l)

	logrus.WithField("config", server.Config).Infof("app config")

	switch server.Config.Spec {
	case Generic:
		if server.Client, err = acpclient.New(server.Config.ClientConfig([]string{"manage_scope_grants"})); err != nil {
			return server, errors.Wrapf(err, "failed to init acp client")
		}
	default:
		if server.Client, err = acpclient.New(server.Config.ClientConfig([]string{"manage_openbanking_consents"})); err != nil {
			return server, errors.Wrapf(err, "failed to init acp client")
		}
	}

	bundle := i18n.NewBundle(server.Config.DefaultLanguage)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	if trans, err = os.ReadDir(server.Config.TransDir); err != nil {
		return server, errors.Wrapf(err, "failed to read dir %s", server.Config.TransDir)
	}

	if server.Config.EnableMFA && server.Config.MFAProvider != "" {
		switch server.Config.MFAProvider {
		case "hypr":
			hyprConfig := HyprConfig{
				Token:   server.Config.HyprToken,
				BaseURL: server.Config.HyprBaseURL,
				AppID:   server.Config.HyprAppID,
			}

			server.MFAStrategy = NewHyprStrategy(hyprConfig)
			server.MFAApprovalChecker = server.MFAStrategy
		default:
			return server, errors.New("unknown MFA provider")
		}
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
	case Generic:
		server.BankClient = NewGenericBankClient(server.Config)
	default:
		return Server{}, errors.New("invalid SPEC configuration")
	}

	if server.Config.EnableMFA && server.Config.MFAProvider == "" {
		logrus.Debugf("mfa is enabled... loading otp db")
		if db, err = shared.InitDB(server.Config.DBFile); err != nil {
			return server, errors.Wrapf(err, "failed to init db")
		}

		if server.OTPRepo, err = NewOTPRepo(db); err != nil {
			return server, errors.Wrapf(err, "failed to init otp repo")
		}

		if server.OTPHandler, err = NewOTPHandler(server.Config, server.OTPRepo, server.SMSClient); err != nil {
			return server, errors.Wrapf(err, "failed to init otp handler")
		}
		server.MFAApprovalChecker = server.OTPHandler
	} else {
		logrus.Debugf("mfa is disabled... skipping otp db load")
	}

	switch server.Config.Spec {
	case OBUK:
		tools := OBUKConsentTools{Trans: server.Trans, Config: server.Config}
		server.AccountAccessConsentHandler = &OBUKAccountAccessConsentHandler{&server, tools}
		server.AccountAccessMFAConsentProvider = &OBUKAccountAccessMFAConsentProvider{&server, tools}
		server.PaymentConsentHandler = &OBUKDomesticPaymentConsentHandler{&server, tools}
		server.PaymentMFAConsentProvider = &DomesticPaymentMFAConsentProvider{&server, tools}
	case OBBR:
		tools := OBBRConsentTools{Trans: server.Trans, Config: server.Config}
		server.AccountAccessConsentHandler = &OBBRAccountAccessConsentHandler{&server, tools}
		server.AccountAccessMFAConsentProvider = &OBBRAccountAccessMFAConsentProvider{&server, tools}
		server.PaymentConsentHandler = &OBBRPaymentConsentHandler{&server, tools}
		server.PaymentMFAConsentProvider = &OBBRPaymentMFAConsentProvider{&server, tools}
	case CDR:
		tools := CDRConsentTools{Trans: server.Trans, Config: server.Config}
		server.AccountAccessConsentHandler = &CDRAccountAccessConsentHandler{&server, tools}
	case FDX:
		tools := FDXConsentTools{Trans: server.Trans, Config: server.Config}
		server.AccountAccessConsentHandler = &FDXAccountAccessConsentHandler{&server, tools}
		server.AccountAccessMFAConsentProvider = &FDXAccountAccessMFAConsentProvider{&server, tools}
	case Generic:
		tools := GenericConsentTools{Trans: server.Trans, Config: server.Config}

		var consentStorage ConsentStorage

		switch server.Config.ConsentStorageMode {
		case "external":
			if consentStorage, err = NewExternalConsentStorage(server.Config.ExternalConsentStorageConfig); err != nil {
				return server, errors.Wrapf(err, "failed to init external consent storage")
			}
		case "identity":
			if consentStorage, err = NewIdentityPoolConsentStorage(server.Config.IdentityPoolConsentStorageConfig); err != nil {
				return server, errors.Wrapf(err, "failed to init identity pool consent storage")
			}
		default:
			return server, fmt.Errorf("unsupported consent storage mode: %s", server.Config.ConsentStorageMode)
		}

		server.AccountAccessConsentHandler = &GenericAccountAccessConsentHandler{
			&server,
			tools,
			consentStorage,
		}
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

		if approved, err = s.MFAApprovalChecker.IsApproved(NewLoginRequest(c)); err != nil {
			s.RenderInvalidRequestError(c, err)
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

	if s.Config.EnableTLSServer {
		logrus.Debugf("running consent page server tls")
		return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
	}
	logrus.Debugf("running consent page server non-tls")
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
