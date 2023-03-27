package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string `mapstructure:"env" envconfig:"APP_ENV"`
		HTTP        HTTPConfig
		PSQL        PSQLConfig
		AUTH        AuthConfig
	}
	HTTPConfig struct {
		Host               string        `mapstructure:"host" envconfig:"HTTP_HOST"`
		Port               string        `mapstructure:"port" envconfig:"HTTP_PORT"`
		ReadTimeout        time.Duration `mapstructure:"read_timeout" envconfig:"HTTP_READ_TIMEOUT"`
		WriteTimeout       time.Duration `mapstructure:"write_timeout" envconfig:"HTTP_WRITE_TIMEOUT"`
		MaxHeaderMegabytes int           `mapstructure:"max_header_bytes" envconfig:"HTTP_MAX_HEADER_BYTES"`
	}
	PSQLConfig struct {
		Host     string `mapstructure:"host" envconfig:"PSQL_HOST"`
		Port     string `mapstructure:"port" envconfig:"PSQL_PORT"`
		User     string `mapstructure:"user" envconfig:"PSQL_USER"`
		Password string `mapstructure:"password" envconfig:"PSQL_PASSWORD"`
		Database string `mapstructure:"database" envconfig:"PSQL_DATABASE"`
	}
	AuthConfig struct {
		JwtSigningKey   string        `mapstructure:"jwt_signing_key" envconfig:"JWT_SIGNING_KEY"`
		AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl" envconfig:"ACCESS_TOKEN_TTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl" envconfig:"REFRESH_TOKEN_TTL"`
	}
)

func Init(path string) (Config, error) {
	var cfg Config

	if err := unmarshalYML(&cfg, path); err != nil {
		return Config{}, err
	}

	if err := unmarshalENV(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func unmarshalYML(cfg *Config, path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("env", &cfg.Environment); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("psql", &cfg.PSQL); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.AUTH); err != nil {
		return err
	}

	return nil
}

func unmarshalENV(cfg *Config) error {

	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}

	return nil
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.read_timeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.read_timeout", defaultHTTPRWTimeout)
}
