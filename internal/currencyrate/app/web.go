package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/currencyrate/presentation"
)

func NewWebApp(container presentation.InteractorFactory) *fiber.App {
	app := fiber.New()

	handlers := presentation.NewCurrencyHandlers(container)
	app.Get("/", handlers.GetCurrency)

	return app
}