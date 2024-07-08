package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/currencyrate"
	"go_service/internal/currencyrate/infrastructure"
	"go_service/internal/subscribers"
	"gorm.io/gorm"
)

func InitWebApp(db *gorm.DB, apiSettings infrastructure.CurrencyAPISettings) *fiber.App {
	app := fiber.New()

	subscribersApp := subscribers.NewApp(db)
	currencyRateApp := currencyrate.NewApp(apiSettings)

	app.Mount("/subscribers/", subscribersApp)
	app.Mount("/rates/", currencyRateApp)

	return app
}
