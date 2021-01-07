package gateway

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sterligov/otus_highload/dating/internal/config"
)

func NewDatabase(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, cfg.Database.Driver, cfg.Database.Addr)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db, nil
}
