package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConn(user, password, dbname, host string, port int) (*sqlx.DB, error) {
	connString := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		return &sqlx.DB{}, fmt.Errorf("failed to connect to db: %w", err)
	}
	return db, nil
}
