package main

import (
	"go_service/internal/adapters"
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
	rateMailer := InitRateMailer(emailAdapter, subscriberAdapter, schedulerAdapter, currencyReader)

	// web app
	webApp := InitWebApp(currencyReader, subscriberAdapter)

	// starting services
	go rateMailer.Run()
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
