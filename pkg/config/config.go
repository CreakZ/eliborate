package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DBUser     = "POSTGRES_USER"
	DBPassword = "POSTGRES_PASSWORD"
	DBHost     = "POSTGRES_HOST"
	DBPort     = "POSTGRES_PORT"
	DBName     = "POSTGRES_DB"

	JWTExpire = "JWT_EXPIRE"
	JWTSecret = "JWT_SECRET"
)

func InitConfig() {
	viper.AddConfigPath("..")
	viper.AddConfigPath("../deploy")

	viper.SetConfigFile("../deploy/test.env")

	viper.SetConfigName("")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config initialization error: %s", err.Error()))
	}
}
