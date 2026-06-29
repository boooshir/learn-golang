package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/pressly/goose/v3"
)

func RunGooseMigrations(db *sql.DB) error {
	// Get absolute path to migrations
	_, callerPath, _, _ := runtime.Caller(0) // Path to this file
	infraDir := filepath.Dir(callerPath)     // infrastructure/db/
	projectRoot := filepath.Dir(filepath.Dir(infraDir))
	migrationsPath := filepath.Join(projectRoot, "infrastructure", "db", "migrations")
	// set migration
	goose.SetTableName("schema_migrations") // custom migration table
	goose.SetDialect("postgres")
	// run migrations
	if err := goose.Up(db, migrationsPath); err != nil {
		return fmt.Errorf("goose up failed: %w", err)
	}
	return nil
}
