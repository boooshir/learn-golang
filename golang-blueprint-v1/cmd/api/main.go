package main

import (
	"golang-blueprint-v1/config"
	"golang-blueprint-v1/internal/database"
	"golang-blueprint-v1/internal/handler"
	"golang-blueprint-v1/internal/repositories"
	"golang-blueprint-v1/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ladong .env file")
	}
	dbCfg := config.LoadDBConfig()
	db, err := database.New(dbCfg)
	if err != nil {
		log.Fatalf("failed to connect to database : %s", err.Error())
	}
	defer db.Close()

	app := fiber.New()
	app.Use(logger.New())

	userRepo := repositories.NewUserRepositoryImpl(db)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userHandler.UserRoute(app)

	// Start server
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server : %w", err)
	}
}
