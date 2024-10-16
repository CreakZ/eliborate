package database

import (
	"fmt"
	"yurii-lib/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func ConnectDB() *sqlx.DB {
	connString := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		viper.GetString(config.DBUser),
		viper.GetString(config.DBPassword),
		viper.GetString(config.DBHost),
		viper.GetInt(config.DBPort),
		viper.GetString(config.DBName))

	db := sqlx.MustConnect("postgres", connString)

	return db
}
