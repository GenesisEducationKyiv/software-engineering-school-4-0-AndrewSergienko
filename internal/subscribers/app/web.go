package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/subscribers/presentation"
	"go_service/internal/subscribers/presentation/handlers"
)

func NewWebApp(container presentation.InteractorFactory) *fiber.App {
	app := fiber.New()

	handlers := handlers.NewSubscriberHandlers(container)
	app.Post("/", handlers.HandleRequest)

	return app
}
