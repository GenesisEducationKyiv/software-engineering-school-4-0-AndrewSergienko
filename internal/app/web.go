package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/adapters"
	"go_service/internal/presentation"
)

func InitWebApp(currencyGateway *adapters.APICurrencyReader, subscriberGateway *adapters.SubscriberAdapter) *fiber.App {
	app := fiber.New()

	currencyHandlers := presentation.InitCurrencyHandlers(currencyGateway)
	subscribersHandles := presentation.InitSubscribersHandlers(subscriberGateway)

	app.Get("/", currencyHandlers.GetCurrency)
	app.Post("/subscribers", subscribersHandles.AddSubscriber)

	return app
}
