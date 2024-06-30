package main

import (
	"go_service/internal/adapters"
	"go_service/internal/adapters/currencyrate"
	"go_service/internal/app"
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
	subscriberAdapter := adapters.NewSubscribersAdapter(db)
	schedulerAdapter := adapters.NewScheduleDBAdapter(db)
	emailAdapter := adapters.NewEmailAdapter(emailSettings)
	readers := currencyrate.CreateReaders(currencyAPISettings)
	currencyReader := currencyrate.NewAPIReaderFacade(readers)

	// background send mail task
	rateMailer := app.InitRateMailer(emailAdapter, subscriberAdapter, schedulerAdapter, currencyReader)

	// web app
	webApp := app.InitWebApp(*currencyReader, subscriberAdapter)

	// starting services
	go rateMailer.Run()
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
