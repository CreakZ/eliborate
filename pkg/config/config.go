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

	RedisHost     = "REDIS_HOST"
	RedisPassword = "REDIS_PASSWORD"
	RedisPort     = "REDIS_PORT"

	JWTAccessExpire  = "JWT_ACCESS_EXPIRE"
	JWTRefreshExpire = "JWT_REFRESH_EXPIRE"
	JWTSecret        = "JWT_SECRET"

	S3AccessKey       = "S3_ACCESS_KEY"
	S3SecretKey       = "S3_SECRET_KEY"
	S3Endpoint        = "S3_ENDPOINT"
	S3Region          = "S3_REGION"
	S3ImageBucketName = "S3_IMAGE_BUCKET_NAME"
)

func InitConfig() {
	viper.SetConfigFile("configs/.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config initialization error: %s", err.Error()))
	}
}
