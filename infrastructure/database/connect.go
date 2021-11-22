package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConn(dbUrl string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
