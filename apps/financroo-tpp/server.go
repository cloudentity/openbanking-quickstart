package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
	"gopkg.in/go-playground/validator.v9"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Server struct {
	Config            Config
	Clients           Clients
	SecureCookie      *securecookie.SecureCookie
	DB                *bolt.DB
	UserRepo          UserRepo
	LoginClient       acpclient.Client
	Validator         *validator.Validate
	UserSecureStorage UserSecureStorage
	LoginURLBuilder   LoginURLBuilder
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
	case "obuk":
		server.Config.ClientScopes = []string{"accounts", "payments", "openid", "offline_access"}
		if server.Clients, err = InitClients(server.Config, NewOBUKSigner, NewOBUKClient, NewOBUKConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}
		if server.LoginURLBuilder, err = NewOBUKLoginURLBuilder(); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	case "obbr":
		server.Config.ClientScopes = []string{"accounts", "payments", "openid", "offline_access", "consents"}
		if server.Clients, err = InitClients(server.Config, NewOBBRSigner, NewOBBRClient, NewOBBRConsentClient); err != nil {
			return server, errors.Wrapf(err, "failed to create clients")
		}
		if server.LoginURLBuilder, err = NewOBBRLoginURLBuilder(server.Clients.AcpAccountsClient); err != nil {
			return server, errors.Wrapf(err, "failed to create login url builder")
		}
	default:
		return server, fmt.Errorf("unsupported spec [%s] in configuration", server.Config.Spec)
	}

	server.SecureCookie = securecookie.New([]byte(server.Config.CookieHashKey), []byte(server.Config.CookieBlockKey))

	if server.DB, err = InitDB(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init db")
	}

	if server.UserRepo, err = NewUserRepo(server.DB); err != nil {
		return server, errors.Wrapf(err, "failed to init user repo")
	}

	server.UserSecureStorage = NewUserSecureStorage(server.SecureCookie)

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

	r.GET("/api/configuration", s.GetConfiguration())

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"featureFlags": s.Config.FeatureFlags})
	})

	return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
}
