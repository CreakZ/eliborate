package storage

import (
	"eliborate/pkg/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewPostgresConn() *sqlx.DB {
	connString := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		viper.GetString(config.PostgresUser),
		viper.GetString(config.PostgresPassword),
		viper.GetString(config.PostgresHost),
		viper.GetInt(config.PostgresPort),
		viper.GetString(config.PostgresDBName))

	db := sqlx.MustConnect("postgres", connString)

	return db
}
