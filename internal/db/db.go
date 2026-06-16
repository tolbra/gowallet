package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, DatabaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, DatabaseURL)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
