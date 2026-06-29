package db

import (
	"database/sql"
	"fmt"
	"note-clean-code/internal/core/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.DBConfig) (*gorm.DB, *sql.DB, error) {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
	)
	gormDb, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to connect to database : %w", err)
	}

	sqlDB, err := gormDb.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return gormDb, sqlDB, nil
}
