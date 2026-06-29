package main

import (
	"github.com/joho/godotenv"
	"golang-blueprint-v1/config"
	"golang-blueprint-v1/internal/database"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ladong .env file")
	}

	dbCfg := config.LoadDBConfig()
	if err := database.RunMigrations(dbCfg); err != nil {
		panic("Migration failed: " + err.Error())
	}
}
