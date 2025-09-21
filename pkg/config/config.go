package config

import (
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
	MeiliIndex     = "MEILI_INDEX"
	MeiliMasterKey = "MEILI_MASTER_KEY"

	AccessControlAllowOrigin  = ""
	AccessControlAllowMethods = ""
	AccessControlAllowHeaders = ""
)

func InitConfig() {
	viper.AutomaticEnv()
}
