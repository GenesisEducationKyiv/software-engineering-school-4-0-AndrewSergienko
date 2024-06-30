package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/adapters"
	"go_service/internal/adapters/currencyrate"
	"go_service/internal/presentation"
)

func InitWebApp(
	currencyGateway currencyrate.APIReaderFacade,
	subscriberGateway *adapters.SubscriberAdapter,
) *fiber.App {
	app := fiber.New()

	currencyHandlers := presentation.NewCurrencyHandlers(&currencyGateway)
	subscribersHandles := presentation.NewSubscribersHandlers(subscriberGateway)

	app.Get("/", currencyHandlers.GetCurrency)
	app.Post("/subscribers", subscribersHandles.AddSubscriber)

	return app
}
