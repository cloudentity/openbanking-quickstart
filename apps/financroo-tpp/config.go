package main

import (
	"time"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

func init() {
	viper.SetDefault("PORT", "8091")
	viper.SetDefault("DB_FILE", "./data/my.db")
	viper.SetDefault("ACP_URL", "")
	viper.SetDefault("ACP_MTLS_URL", "")
	viper.SetDefault("APP_HOST", "")
	viper.SetDefault("UI_URL", "")
	viper.SetDefault("CERT_FILE", "")
	viper.SetDefault("KEY_FILE", "")
	viper.SetDefault("TENANT", "")
	viper.SetDefault("COOKIE_HASH_KEY", []byte("secret-key"))
	viper.SetDefault("COOKIE_BLOCK_KEY", []byte("this-is-32-len-block-key"))
}

type ClientConfig struct {
	ServerID string   `mapstructure:"server_id"`
	ClientID string   `mapstructure:"client_id"`
	Scopes   []string `mapstructure:"scopes"`
}

type HTTPClientConfig struct {
	RootCA   string `mapstructure:"root_ca"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
	Timeout  time.Duration
}

type AcpClient struct {
	ClientConfig     `mapstructure:",squash"`
	HTTPClientConfig `mapstructure:",squash"`
}

type BankID string

type BankConfig struct {
	ID        BankID
	BankType  string `mapstructure:"bank_type"`
	URL       string
	AcpClient AcpClient `mapstructure:"acp_client"`
}

type FeatureFlags struct {
	Investments bool `mapstructure:"investments"`
}

type Config struct {
	Port           int
	DBFile         string `mapstructure:"db_file"`
	ACPURL         string `mapstructure:"acp_url" validate:"required,url"`
	ACPInternalURL string `mapstructure:"ACP_MTLS_URL" validate:"required,url"`
	AppHost        string `mapstructure:"app_host" validate:"required"`
	Tenant         string `mapstructure:"tenant" validate:"required"`
	UIURL          string `mapstructure:"ui_url" validate:"required,url"`
	CertFile       string `mapstructure:"cert_file" validate:"required"`
	KeyFile        string `mapstructure:"key_file" validate:"required"`
	CookieHashKey  []byte `mapstructure:"cookie_hash_key"`
	CookieBlockKey []byte `mapstructure:"cookie_block_key"`
	Banks          []BankConfig
	FeatureFlags   FeatureFlags `mapstructure:"feature_flags"`
}

func (s *Server) GetBankConfigByBankID(bankID BankID) BankConfig {
	for _, config := range s.Config.Banks {
		if bankID == config.ID {
			return config
		}
	}

	return BankConfig{}
}

var Validator = validator.New()

func LoadConfig() (Config, error) {
	var (
		config = Config{}
		err    error
	)

	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./data")

	if err = viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err = Validator.Struct(config); err != nil {
		return config, err
	}

	return config, nil
}
