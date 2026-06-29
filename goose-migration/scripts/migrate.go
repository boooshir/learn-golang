package main

import (
	"log"
	"note-clean-code/internal/core/config"
	"note-clean-code/internal/infrastructure/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ladong .env file")
	}
	// Load config from environment variables
	dbCfg := config.LoadDBConfig()
	// Connect to DB
	_, sqlDB, err := db.NewPostgresDB(dbCfg)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	defer sqlDB.Close()

	if err := db.RunGooseMigrations(sqlDB); err != nil {
		panic("Migration failed: " + err.Error())
	}
}
