package main

import (
	"github.com/caarlos0/env/v6"
)

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
	ClientScopes                []string
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

	return config, nil
}
