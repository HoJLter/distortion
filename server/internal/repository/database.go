package repository

import (
	"context"
	"log/slog"
	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDatabase(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, cfg.Url)
	if err != nil {
		slog.Error("Error while db init", "err", err)
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		slog.Error("Error while db ping", "err", err)
		return nil, err
	}

	return db, nil
}
