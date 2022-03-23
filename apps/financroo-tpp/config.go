package main

import (
	"github.com/caarlos0/env/v6"
)

type FeatureFlags struct {
	Investments bool `env:"investments"`
}

type BankID string

type Config struct {
	Port           int          `env:"PORT" envDefault:"8091"`
	DBFile         string       `env:"DB_FILE"`
	ACPURL         string       `env:"ACP_URL" validate:"required,url"`
	ACPInternalURL string       `env:"ACP_MTLS_URL" validate:"required,url"`
	AppHost        string       `env:"APP_HOST" validate:"required"`
	Tenant         string       `env:"TENANT" validate:"required"`
	UIURL          string       `env:"UI_URL" validate:"required,url"`
	CertFile       string       `env:"CERT_FILE" validate:"required"`
	KeyFile        string       `env:"KEY_FILE" validate:"required"`
	CookieHashKey  string       `env:"COOKIE_HASH_KEY" envDefault:"secret-key"`
	CookieBlockKey string       `env:"COOKIE_BLOCK_KEY" envDefault:"this-is-32-len-block-key"`
	FeatureFlags   FeatureFlags `env:"FEATURE_FLAGS"`
	Spec           string       `env:"SPEC" validate:"required"`
	BankURL        string       `env:"BANK_URL" validate:"required"`
	RootCA         string       `env:"ROOT_CA" validate:"required"`
	ClientID       string       `env:"CLIENT_ID" validate:"required"`
	ServerID       string       `env:"OPENBANKING_SERVER_ID" validate:"required"`
	ClientScopes   []string
}

func LoadConfig() (Config, error) {
	var (
		config = Config{}
		err    error
	)

	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, nil
}
