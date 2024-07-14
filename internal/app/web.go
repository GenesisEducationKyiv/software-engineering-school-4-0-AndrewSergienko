package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/currencyrate"
	"go_service/internal/currencyrate/infrastructure"
	"go_service/internal/customers"
	"gorm.io/gorm"
)

func InitWebApp(db *gorm.DB, apiSettings infrastructure.CurrencyAPISettings) *fiber.App {
	app := fiber.New()

	subscribersApp := customers.NewApp(db)
	currencyRateApp := currencyrate.NewApp(apiSettings)

	app.Mount("/customers/", subscribersApp)
	app.Mount("/rates/", currencyRateApp)

	return app
}
