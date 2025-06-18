package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
)

const (
	PostgresUser     = "POSTGRES_USER"
	PostgresPassword = "POSTGRES_PASSWORD"
	PostgresHost     = "POSTGRES_HOST"
	PostgresPort     = "POSTGRES_PORT"
	PostgresDBName   = "POSTGRES_DB"

	RedisHost     = "REDIS_HOST"
	RedisPassword = "REDIS_PASSWORD"
	RedisPort     = "REDIS_PORT"

	JWTAccessExpire  = "JWT_ACCESS_EXPIRE"
	JWTRefreshExpire = "JWT_REFRESH_EXPIRE"
	JWTSecret        = "JWT_SECRET"

	MeiliHost      = "MEILI_HOST"
	MeiliPort      = "MEILI_PORT"
	MeiliMasterKey = "MEILI_MASTER_KEY"
)

type CorsConfig struct {
	AccessControlAllowOrigin  string `toml:"access_control_allow_origin"`
	AccessControlAllowMethods string `toml:"access_control_allow_methods"`
	AccessControlAllowHeaders string `toml:"access_control_allow_headers"`
}

func InitConfig() *CorsConfig {
	cfg, err := newCorsConfig()
	if err != nil {
		panic(fmt.Errorf("cors initialization error: %w", err))
	}

	viper.SetConfigFile("./configs/.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("global initialization error: %w", err))
	}

	return cfg
}

func newCorsConfig() (*CorsConfig, error) {
	cfg := &CorsConfig{}

	_, err := toml.DecodeFile("./configs/cors.config.toml", cfg)
	if err != nil {
		return &CorsConfig{}, err
	}

	return cfg, nil
}

func InitTestConfig(cfgPath string) {
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("test config initialization error: %s", err.Error()))
	}
}
