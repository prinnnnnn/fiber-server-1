package database

import (
	"context"
	"fiber-server-1/internal/adapter/config"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx *context.Context, config *config.DB) (*pgxpool.Pool, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Connection, config.Port, config.Name)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	// Set pool size (MaxConns defines the pool size)
	poolConfig.MaxConns = 1                       // Maximum number of connections in the pool
	poolConfig.MinConns = 5                       // Minimum number of idle connections in the pool
	poolConfig.MaxConnIdleTime = 15 * time.Minute // Time before idle connections are closed
	poolConfig.MaxConnLifetime = 2 * time.Hour    // Maximum lifetime for a connection

	// Create the pool
	dbPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return dbPool, nil

}
