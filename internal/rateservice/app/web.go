package app

import (
	"context"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/rateservice/currencyrate"
	"go_service/internal/rateservice/customers"
	"go_service/internal/rateservice/infrastructure"
	"gorm.io/gorm"
)

func InitWebApp(
	ctx context.Context,
	db *gorm.DB,
	conn jetstream.JetStream,
	apiSettings infrastructure.CurrencyAPISettings,
) *fiber.App {
	app := fiber.New()
	app.Use(swagger.New())

	subscribersApp := customers.NewApp(ctx, db, conn)
	currencyRateApp := currencyrate.NewApp(apiSettings)

	app.Mount("/customers/", subscribersApp)
	app.Mount("/rates/", currencyRateApp)

	return app
}
