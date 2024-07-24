package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/presentation/handlers"
	"go_service/internal/rateservice/customers/presentation/handlers/createcustomer"
	"go_service/internal/rateservice/customers/presentation/handlers/deletecustomer"
)

func NewWebApp(container presentation.InteractorFactory, eventGateway adapters.NatsEventEmitter) *fiber.App {
	app := fiber.New()

	subscribeHandler := createcustomer.New(container, eventGateway)
	unsubscribeHandler := deletecustomer.New(container, eventGateway)
	getAllHandler := handlers.NewGetAllHandler(container)

	app.Post("/", subscribeHandler.HandleRequest)
	app.Delete("/", unsubscribeHandler.HandleRequest)
	app.Get("/", getAllHandler.HandleRequest)

	return app
}
