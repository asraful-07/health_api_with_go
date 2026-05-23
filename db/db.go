package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	connStr := os.Getenv("DB_STRING")

	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
}
