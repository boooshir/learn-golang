package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	var engine = mustache.New("./template", ".mustache")
	app := fiber.New(fiber.Config{
		Views:        engine,
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Prefork:      true,
	})
	app.Use(logger.New())

	app.Use(func(ctx *fiber.Ctx) error {
		fmt.Println("middleware before processing request")
		err := ctx.Next()
		fmt.Println("middleware after processing request")
		return err
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world")
	})

	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("middleware before processing request API")
		err := ctx.Next()
		fmt.Println("middleware after processing request API")
		return err
	})

	app.Get("/api/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world")
	})
	app.Get("/view", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"title":   "Hello title",
			"header":  "Hello header",
			"content": "Hello content",
		})
	})
	err := app.Listen("localhost:3000")

	if err != nil {
		panic(err)
	}
}
