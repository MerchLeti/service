package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context) (*Database, error) {
	pool, err := pgxpool.New(
		ctx,
		DSNFromEnv(),
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't open pool: %w", err)
	}
	return newDatabase(pool), nil
}
