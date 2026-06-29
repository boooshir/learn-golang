package handler

import (
	"fmt"
	"golang-blueprint-v1/internal/models"
	"golang-blueprint-v1/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) UserRoute(app *fiber.App) {
	auth := app.Group("/api/auth")
	auth.Post("/register", h.Register)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error parse": err})
	}
	if err := h.service.Create(c.Context(), &req); err != nil {
		fmt.Printf(" error at handler %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"service": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Success": "user is created"})
}
