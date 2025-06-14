package config

import (
	"fmt"

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

func InitConfig() {
	viper.SetConfigFile("./configs/.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config initialization error: %s", err.Error()))
	}
}

func InitTestConfig(cfgPath string) {
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("test config initialization error: %s", err.Error()))
	}
}
