package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConn(dbUrl string) *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		if err != nil {
			return nil
		}
		os.Exit(1)
	}

	return conn
}
