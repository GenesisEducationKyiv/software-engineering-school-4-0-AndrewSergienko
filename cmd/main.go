package main

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal"
	"go_service/internal/adapters"
	"go_service/internal/api"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"log"
)

func main() {
	// app configuration
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()
	emailSettings := infrastructure.GetEmailSettings()

	db := database.InitDatabase(databaseSettings)

	// initializing adapters
	subscriberAdapter := adapters.GetSubscribersAdapter(db)
	schedulerAdapter := adapters.GetSchedulerDbAdapter(db)
	emailAdapter := adapters.GetEmailAdapter(emailSettings)
	currencyReader := adapters.GetAPICurrencyReader(currencyAPISettings)

	// background send mail task
	rateMailer := internal.RateMailer{
		Es: emailAdapter,
		Sr: subscriberAdapter,
		Sg: schedulerAdapter,
		Cr: currencyReader,
	}

	// web app
	app := fiber.New()
	currencyHandlers := api.InitCurrencyHandlers(currencyReader)
	app.Get("/", currencyHandlers.GetCurrency)

	// starting services
	go rateMailer.Run()
	log.Fatalf("App failed with error: %v", app.Listen(":8080"))
}
