package db

import (
	"context"
	"fmt"

	"github.com/MerchLeti/service/internal/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "shop"
	password = "shop"
	dbname   = "shop"
)

func New(ctx context.Context) (*Database, error) {
	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.Get("POSTGRES_HOST", host),
			env.Get("POSTGRES_PORT", port),
			env.Get("POSTGRES_USER", user),
			env.Get("POSTGRES_PASSWORD", password),
			env.Get("POSTGRES_DB", dbname),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't open pool: %w", err)
	}
	return newDatabase(pool), nil
}
