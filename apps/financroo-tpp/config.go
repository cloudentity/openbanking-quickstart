package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var Validator = validator.New()

type FeatureFlags struct {
	Investments bool `env:"investments"`
}

type BankID string

type Spec string

const (
	OBUK    Spec = "obuk"
	OBBR    Spec = "obbr"
	CDR     Spec = "cdr"
	FDX     Spec = "fdx"
	GENERIC Spec = "generic"
)

type Config struct {
	Port                        int          `env:"PORT" envDefault:"8091"`
	DBFile                      string       `env:"DB_FILE" envDefault:"/app/data/my.db"`
	ACPURL                      string       `env:"ACP_URL" validate:"required,url"`
	ACPInternalURL              string       `env:"ACP_MTLS_URL" validate:"required,url"`
	AppHost                     string       `env:"APP_HOST" validate:"required"`
	Tenant                      string       `env:"TENANT" validate:"required"`
	UIURL                       string       `env:"UI_URL" validate:"required,url"`
	CertFile                    string       `env:"CERT_FILE" envDefault:"/certs/tpp_cert.pem"`
	KeyFile                     string       `env:"KEY_FILE" envDefault:"/certs/tpp_key.pem"`
	CookieHashKey               string       `env:"COOKIE_HASH_KEY" envDefault:"secret-key"`
	CookieBlockKey              string       `env:"COOKIE_BLOCK_KEY" envDefault:"this-is-32-len-block-key"`
	FeatureFlags                FeatureFlags `env:"FEATURE_FLAGS"`
	Spec                        Spec         `env:"SPEC" validate:"required"`
	BankURL                     string       `env:"BANK_URL" validate:"required"`
	RootCA                      string       `env:"ROOT_CA" envDefault:"/certs/ca.pem"`
	ClientID                    string       `env:"CLIENT_ID" envDefault:"bugkgm23g9kregtu051g"`
	ClientSecret                string       `env:"CLIENT_SECRET" envDefault:"-TlfoycUiE0qNi-XUBFDfTxMlhHTCjVxOF6pLrWZbQA"` // only required for fdx
	ServerID                    string       `env:"OPENBANKING_SERVER_ID" validate:"required"`
	EnableTLSServer             bool         `env:"ENABLE_TLS_SERVER" envDefault:"true"`
	Currency                    string       `env:"CURRENCY"` // optional custom currency, one of=USD AUD GBP BRL EUR
	AssertionSigningAlg         string       `env:"ASSERTION_SIGNING_ALG" envDefault:"PS256"`
	AssertionSigningKeyFile     string       `env:"ASSERTION_SIGNING_KEY_FILE" envDefault:"/certs/private.ps.pem"`
	RequestObjectSigningAlg     string       `env:"REQUEST_OBJECT_SIGNING_ALG" envDefault:"ES256"`
	RequestObjectSigningKeyFile string       `env:"REQUEST_OBJECT_SIGNING_KEY_FILE" envDefault:"/certs/private.es.pem"`
	EnableDCR                   bool         `env:"ENABLE_DCR" envDefault:"false"`
	BanksConfigFile             string       `env:"BANKS_CONFIG_FILE" envDefault:"/app/banks.json"`

	ClientScopes []string
	Banks        BanksConfig
}

func (c *Config) SetImplicitValues() {
	if c.Currency == "" {
		switch c.Spec {
		case FDX:
			c.Currency = "USD"
		case CDR:
			c.Currency = "AUD"
		case OBBR:
			c.Currency = "BRL"
		case OBUK:
			c.Currency = "GBP"
		case GENERIC:
			c.Currency = "USD"
		}
	}

	switch c.Spec {
	case OBUK:
		c.ClientScopes = []string{"accounts", "payments", "openid", "offline_access"}
	case OBBR:
		c.ClientScopes = []string{"accounts", "payments", "openid", "offline_access", "consents"}
	case CDR:
		c.ClientScopes = []string{"offline_access", "openid", "bank:accounts.basic:read", "bank:accounts.detail:read", "bank:transactions:read", "common:customer.basic:read"}
	case FDX:
		c.ClientScopes = []string{"offline_access", "fdx:accountdetailed:read", "READ_CONSENTS", "fdx:accountbasic:read", "fdx:transactions:read"}
	case GENERIC:
		c.ClientScopes = []string{"openid", "email", "sample", "offline_access"}
	}
}

type BankConfig struct {
	ID             BankID `json:"id" validate:"required"`
	Name           string `json:"name"`
	IconURL        string `json:"icon_url"`
	LogoURL        string `json:"logo_url"`
	URL            string `json:"url" validate:"required,url"`
	ACPURL         string `json:"acp_url" validate:"required,url"`
	ACPInternalURL string `json:"acp_internal_url" validate:"required,url"`
	Tenant         string `json:"tenant" validate:"required"`
	Server         string `json:"server" validate:"required"`
	EnableDCR      bool   `json:"enable_dcr"`
	ClientID       string `json:"client_id"`
}

type RawBankConfig struct {
	Banks BanksConfig `json:"banks" validate:"dive"`
}

type BanksConfig []BankConfig

func (b *BanksConfig) GetIDs() (ids []BankID) {
	for _, x := range *b {
		ids = append(ids, x.ID)
	}

	return ids
}

func LoadConfig() (Config, error) {
	var (
		config = Config{}
		err    error
	)

	if err = env.Parse(&config); err != nil {
		return config, err
	}

	config.SetImplicitValues()

	if config.Banks, err = LoadBanksConfig(config); err != nil {
		return config, err
	}

	return config, nil
}

func LoadBanksConfig(config Config) (BanksConfig, error) {
	var (
		bs  []byte
		c   RawBankConfig
		err error
	)

	if config.BanksConfigFile == "" {
		// compatibility with old config
		return []BankConfig{
			{
				ID:             "gobank",
				URL:            config.BankURL,
				ACPURL:         config.ACPURL,
				ACPInternalURL: config.ACPInternalURL,
				Tenant:         config.Tenant,
				Server:         config.ServerID,
				EnableDCR:      config.EnableDCR,
				ClientID:       config.ClientID,
			},
		}, nil
	}

	if bs, err = ioutil.ReadFile(config.BanksConfigFile); err != nil {
		return BanksConfig{}, errors.Wrapf(err, "failed to read banks config file: %s", config.BanksConfigFile)
	}

	if err = json.Unmarshal(bs, &c); err != nil {
		return BanksConfig{}, errors.Wrapf(err, "failed to unmarshal banks config file: %s", config.BanksConfigFile)
	}

	if err = Validator.Struct(c); err != nil {
		return BanksConfig{}, errors.Wrapf(err, "failed to validate banks config")
	}

	return c.Banks, nil
}
