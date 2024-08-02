package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/currencyrate/adapters"
	"go_service/internal/rateservice/currencyrate/presentation"
)

func NewWebApp(container presentation.InteractorFactory, cacheAdapter adapters.CacheRateAdapter) *fiber.App {
	app := fiber.New()

	handlers := presentation.NewCurrencyHandlers(container, cacheAdapter)
	app.Get("/", handlers.GetCurrency)

	return app
}
