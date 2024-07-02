package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/subscribers/presentation"
)

func NewApp(container presentation.InteractorFactory) *fiber.App {
	app := fiber.New()

	handlers := presentation.NewSubscribersHandlers(container)
	app.Post("/subscribers", handlers.AddSubscriber)

	return app
}
