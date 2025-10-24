package api

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	restHandler := &rest.RestHandler{
		App: app,
	}

	setupRoutes(restHandler)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// User handler
	handlers.SetupUserRoutes(rh)

	// transaction handler
	// catalog handler
}
