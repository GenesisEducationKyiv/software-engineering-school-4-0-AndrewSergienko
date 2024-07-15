package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/presentation/handlers"
)

func NewWebApp(container presentation.InteractorFactory) *fiber.App {
	app := fiber.New()

	subscribeHandler := handlers.NewSubscriberHandlers(container)
	getAllHandler := handlers.NewGetAllHandler(container)

	app.Post("/", subscribeHandler.HandleRequest)
	app.Get("/", getAllHandler.HandleRequest)

	return app
}
