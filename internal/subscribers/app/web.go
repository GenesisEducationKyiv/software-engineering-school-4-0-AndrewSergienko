package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/subscribers/presentation"
	"go_service/internal/subscribers/presentation/handlers"
)

func NewWebApp(container presentation.InteractorFactory) *fiber.App {
	app := fiber.New()

	subscribeHandler := handlers.NewSubscriberHandlers(container)
	app.Post("/", subscribeHandler.HandleRequest)

	return app
}

func NewInternalWebApp(container presentation.InteractorFactory) *fiber.App {
	app := fiber.New()

	subscribeHandler := handlers.NewSubscriberHandlers(container)
	getAllHandler := handlers.NewGetAllHandler(container)

	app.Post("/", subscribeHandler.HandleRequest)
	app.Get("/", getAllHandler.HandleRequest)

	return app
}
