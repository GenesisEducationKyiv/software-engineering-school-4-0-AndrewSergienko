package app

import (
	"github.com/gofiber/fiber/v2"
	presentation2 "go_service/internal/currency_rate/presentation"
	subscribers "go_service/internal/subscribers/app"
)

func InitWebApp(container *IoC) *fiber.App {
	app := fiber.New()

	subscribersApp := subscribers.NewApp(container)
	currencyHandlers := presentation2.NewCurrencyHandlers(container)

	app.Get("/", currencyHandlers.GetCurrency)

	app.Mount("/subscribers/", subscribersApp)

	return app
}
