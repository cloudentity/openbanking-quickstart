package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/square/go-jose.v2"

	"github.com/cloudentity/openbanking-quickstart/utils"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Server struct {
	Config                   Config
	Clients                  Clients
	SecureCookie             *securecookie.SecureCookie
	DB                       *bolt.DB
	UserRepo                 UserRepo
	LoginClient              acpclient.Client
	Validator                *validator.Validate
	UserSecureStorage        UserSecureStorage
	LoginURLBuilder          LoginURLBuilder
	SignatureVerificationKey jose.JSONWebKey
}

func NewServer() (Server, error) {
	var (
		server      = Server{}
		dcrResponse DCRClientCreated
		ok          bool
		clientID    string
		err         error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	logrus.WithField("config", server.Config).Info("Config loaded")

	if server.DB, err = InitDB(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init db")
	}

	if server.Config.EnableDCR {
		storage := ClientIDStorage{DB: server.DB}

		if clientID, ok, err = storage.Get(); err != nil {
			return server, errors.Wrapf(err, "failed to fetch client id from db")
		}

		if !ok {
			if dcrResponse, err = RegisterClient(context.Background(), server.Config); err != nil {
				return server, errors.Wrapf(err, "failed to register client")
			}

			if err = storage.Set(dcrResponse.ClientID); err != nil {
				return server, errors.Wrapf(err, "failed to store client id in db")
			}

			server.Config.ClientID = dcrResponse.ClientID

			logrus.Infof("client dynamically registered, id: %s", dcrResponse.ClientID)
		} else {
			logrus.Infof("client already registered, use id: %s", clientID)

			server.Config.ClientID = clientID
		}
	}

	switch server.Config.Spec {
	case OBUK:
		if server.Clients, err = InitClients(server.Config, NewOBUKSigner, NewOBUKClient, NewOBUKConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}
		if server.LoginURLBuilder, err = NewOBUKLoginURLBuilder(); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	case OBBR:
		if server.Clients, err = InitClients(server.Config, NewOBBRSigner, NewOBBRClient, NewOBBRConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}
		if server.LoginURLBuilder, err = NewOBBRLoginURLBuilder(server.Clients.AcpAccountsClient); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	case CDR:
		if server.Clients, err = InitClients(server.Config, nil, NewCDRClient, NewCDRConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}
		if server.LoginURLBuilder, err = NewCDRLoginURLBuilder(server.Config); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	case FDX:
		if server.Clients, err = InitClients(server.Config, nil, NewFDXBankClient, NewFDXConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}

		if server.LoginURLBuilder, err = NewFDXLoginURLBuilder(server.Config); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	case GENERIC:
		if server.Clients, err = InitClients(server.Config, nil, NewGenericBankClient, NewGenericConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}

		if server.LoginURLBuilder, err = NewGenericLoginURLBuilder(server.Config); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	default:
		return server, fmt.Errorf("unsupported spec [%s] in configuration", server.Config.Spec)
	}

	server.SecureCookie = securecookie.New([]byte(server.Config.CookieHashKey), []byte(server.Config.CookieBlockKey))

	server.UserRepo = UserRepo{DB: server.DB}

	server.UserSecureStorage = NewUserSecureStorage(server.SecureCookie)

	if server.SignatureVerificationKey, err = utils.GetServerKey(&server.Clients.AcpAccountsClient, utils.SIG); err != nil {
		return server, errors.Wrapf(err, "failed to retrieve server signature key")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())

	r.POST("/api/connect/:bankId", s.ConnectBank())
	r.GET("/api/callback", s.ConnectBankCallback())
	r.DELETE("/api/disconnect/:bankId", s.DisconnectBank())

	r.POST("/api/domestic-payment-consent", s.CreateDomesticPaymentConsent())
	r.GET("/api/domestic/callback", s.DomesticPaymentCallback())

	r.GET("/api/accounts", s.GetAccounts())
	r.GET("/api/transactions", s.GetTransactions())
	r.GET("/api/balances", s.GetBalances())
	r.GET("/api/banks", s.ConnectedBanks())

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"featureFlags": s.Config.FeatureFlags,
			"spec":         s.Config.Spec,
			"currency":     s.Config.Currency,
		})
	})

	if s.Config.EnableTLSServer {
		logrus.Debugf("running financroo server tls")
		return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
	}
	logrus.Debugf("running financroo server non-tls")
	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}
