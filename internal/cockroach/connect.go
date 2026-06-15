package cockroach

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	dataSource := fmt.Sprintf("postgresql://%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	pool, err := pgxpool.New(ctx, dataSource)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging pool: %w", err)
	}

	return pool, nil
}
