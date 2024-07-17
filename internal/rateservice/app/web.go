package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"go_service/internal/rateservice/currencyrate"
	"go_service/internal/rateservice/currencyrate/infrastructure"
	"go_service/internal/rateservice/customers"

	"gorm.io/gorm"
)

func InitWebApp(db *gorm.DB, nc *nats.Conn, apiSettings infrastructure.CurrencyAPISettings) *fiber.App {
	app := fiber.New()

	subscribersApp := customers.NewApp(db, nc)
	currencyRateApp := currencyrate.NewApp(apiSettings)

	app.Mount("/customers/", subscribersApp)
	app.Mount("/rates/", currencyRateApp)

	return app
}
