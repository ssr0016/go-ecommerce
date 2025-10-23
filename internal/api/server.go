package api

import (
	"go-ecommerce/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", healthCheck)

	app.Listen(config.ServerPort)
}

func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "API is up and running",
	})
}
