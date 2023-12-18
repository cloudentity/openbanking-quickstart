package main

import (
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
	"golang.org/x/text/language"

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
	Port                             int           `env:"PORT"                envDefault:"8080"`
	ClientID                         string        `env:"CLIENT_ID,required"`
	ClientSecret                     string        `env:"CLIENT_SECRET"       envDefault:"pMPBmv62z3Jt1S4sWl2qRhOhEGPVZ9EcujGL7Xy0-E0"`
	IssuerURL                        *url.URL      `env:"ISSUER_URL,required"`
	Timeout                          time.Duration `env:"TIMEOUT"             envDefault:"5s"`
	RootCA                           string        `env:"ROOT_CA"             envDefault:"/ca.pem"`
	CertFile                         string        `env:"CERT_FILE"           envDefault:"/bank_cert.pem"`
	KeyFile                          string        `env:"KEY_FILE"            envDefault:"/bank_key.pem"`
	BankIDClaim                      string        `env:"BANK_ID_CLAIM"       envDefault:"sub"`
	EnableMFA                        bool          `env:"ENABLE_MFA"`
	MFAProvider                      string        `env:"MFA_PROVIDER"`
	OTPMode                          string        `env:"OTP_MODE"            envDefault:"demo"`
	HyprToken                        string        `env:"HYPR_TOKEN"`
	HyprBaseURL                      string        `env:"HYPR_BASE_URL"`
	HyprAppID                        string        `env:"HYPR_APP_ID"`
	TwilioAccountSid                 string        `env:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken                  string        `env:"TWILIO_AUTH_TOKEN"`
	TwilioFrom                       string        `env:"TWILIO_FROM"         envDefault:"Cloudentity"`
	DBFile                           string        `env:"DB_FILE"             envDefault:"/data/my.db"`
	MFAClaim                         string        `env:"MFA_CLAIM"           envDefault:"mobile_verified"`
	LogLevel                         string        `env:"LOG_LEVEL"           envDefault:"info"`
	DevMode                          bool          `env:"DEV_MODE"`
	DefaultLanguage                  language.Tag  `env:"DEFAULT_LANGUAGE"    envDefault:"en-us"`
	TransDir                         string        `env:"TRANS_DIR"           envDefault:"./translations"`
	Spec                             Spec          `env:"SPEC,required"`
	Otp                              OtpConfig
	EnableTLSServer                  bool     `env:"ENABLE_TLS_SERVER" envDefault:"true"`
	Currency                         Currency `env:"CURRENCY"` // optional custom currency, one of=USD AUD GBP BRL EUR
	BankClientConfig                 BankClientConfig
	ConsentStorageMode               string `env:"CONSENT_STORAGE_MODE"` // only for generic, one of: external | identity
	ExternalConsentStorageConfig     ExternalConsentStorageConfig
	IdentityPoolConsentStorageConfig IdentityPoolConsentStorageConfig
	BankLogo                         string `env:"BANK_LOGO" envDefault:"bank_logo.svg"`
}

type ExternalConsentStorageConfig struct {
	URL      *url.URL `env:"EXTERNAL_CONSENT_STORAGE_URL"` // only for generic spec
	CertFile string   `env:"EXTERNAL_CONSENT_STORAGE_CLIENT_CERT_FILE"`
	KeyFile  string   `env:"EXTERNAL_CONSENT_STORAGE_CLIENT_KEY_FILE"`
	RootCA   string   `env:"EXTERNAL_CONSENT_STORAGE_CLIENT_ROOT_CA"`
}

type IdentityPoolConsentStorageConfig struct {
	IssuerURL    *url.URL `env:"IDENTITY_POOL_CONSENT_STORAGE_ISSUER_URL"`
	ClientID     string   `env:"IDENTITY_POOL_CONSENT_STORAGE_CLIENT_ID"`
	ClientSecret string   `env:"IDENTITY_POOL_CONSENT_STORAGE_CLIENT_SECRET"`
	RootCA       string   `env:"IDENTITY_POOL_CONSENT_STORAGE_ROOT_CA"`
	PoolID       string   `env:"IDENTITY_POOL_CONSENT_STORAGE_POOL_ID"       envDefault:"hyperscalebank-consent-storage"`
}

type Currency string

func (c Currency) ToString() string {
	switch c {
	case "USD":
		return "$"
	case "AUD":
		return "$"
	case "GBP":
		return "£"
	case "EUR":
		return "€"
	case "BRL":
		return "R$"
	default:
		return "$"
	}
}

type OtpConfig struct {
	Type       string        `env:"OTP_TYPE"        envDefault:"otp"`
	RequestURL string        `env:"OTP_REQUEST_URL"`
	VerifyURL  string        `env:"OTP_VERIFY_URL"`
	Timeout    time.Duration `env:"OTP_TIMEOUT"     envDefault:"10s"`
	AuthHeader string        `env:"OTP_AUTH_HEADER"`
}

type BankClientConfig struct {
	URL          *url.URL `env:"BANK_URL,required"`
	AccountsURL  *url.URL `env:"BANK_ACCOUNTS_ENDPOINT"`
	CertFile     string   `env:"BANK_CLIENT_CERT_FILE"`
	KeyFile      string   `env:"BANK_CLIENT_KEY_FILE"`
	TokenURL     string   `env:"BANK_CLIENT_TOKEN_URL"`
	ClientID     string   `env:"BANK_CLIENT_ID"`
	ClientSecret string   `env:"BANK_CLIENT_SECRET"`
	Scopes       []string `env:"BANK_CLIENT_SCOPES"`
}

func (c *Config) ClientConfig(scopes []string) acpclient.Config {
	return acpclient.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		IssuerURL:    c.IssuerURL,
		Scopes:       scopes,
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
