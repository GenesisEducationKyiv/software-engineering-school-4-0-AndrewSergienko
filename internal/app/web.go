package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/presentation/handlers"
)

func InitWebApp(container *IoC) *fiber.App {
	app := fiber.New()

	currencyHandlers := handlers.NewCurrencyHandlers(container)
	//subscribersHandles := handlers.NewSubscribersHandlers(container)

	app.Get("/", currencyHandlers.GetCurrency)
	//app.Post("/subscribers", subscribersHandles.AddSubscriber)

	return app
}
