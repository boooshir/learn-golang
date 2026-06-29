package database

import (
	"context"
	"fmt"
	"golang-blueprint-v1/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.DBConfig) (*DB, error) {
	ctx := context.Background()

	// Parse config intoo pgxpool.Config
	poolCfg, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DB config: %s", err)
	}
	poolCfg.MaxConns = int32(cfg.MaxOpen)
	poolCfg.MinConns = int32(cfg.MaxIdle)
	poolCfg.MaxConnLifetime, _ = time.ParseDuration(cfg.MaxLifetime)

	// connect
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("DB ping failed, error : %s", err)
	}
	fmt.Println("DB connection established")
	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
