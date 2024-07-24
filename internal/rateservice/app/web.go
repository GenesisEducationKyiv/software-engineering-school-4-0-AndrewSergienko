package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"go_service/internal/rateservice/currencyrate"
	"go_service/internal/rateservice/customers"
	"go_service/internal/rateservice/infrastructure"

	"gorm.io/gorm"
)

func InitWebApp(db *gorm.DB, conn nats.JetStreamContext, apiSettings infrastructure.CurrencyAPISettings) *fiber.App {
	app := fiber.New()

	subscribersApp := customers.NewApp(db, conn)
	currencyRateApp := currencyrate.NewApp(apiSettings)

	app.Mount("/customers/", subscribersApp)
	app.Mount("/rates/", currencyRateApp)

	return app
}
