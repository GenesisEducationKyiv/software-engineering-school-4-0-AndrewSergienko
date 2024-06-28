package main

import (
	"go.uber.org/dig"
	"go_service/internal/adapters"
	"go_service/internal/adapters/currencyrate"
	"go_service/internal/app"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"log"
)

func main() {
	container := dig.New()

	// app configuration
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()
	emailSettings := infrastructure.GetEmailSettings()

	db := database.InitDatabase(databaseSettings)

	container.Provide(func() *adapters.SubscriberAdapter {
		return adapters.NewSubscribersAdapter(db)
	})

	container.Provide(func() *adapters.ScheduleDBAdapter {
		return adapters.NewScheduleDBAdapter(db)
	})

	container.Provide(func() adapters.EmailAdapter {
		return adapters.NewEmailAdapter(emailSettings)
	})

	container.Provide(func() *currencyrate.APIReaderFacade {
		readers := currencyrate.CreateReaders(currencyAPISettings)
		return currencyrate.NewAPIReaderFacade(readers)
	})

	// background send mail task
	//rateMailer := app.InitRateMailer(emailAdapter, subscriberAdapter, schedulerAdapter, currencyReader)

	// web app
	webApp := app.InitWebApp(*currencyReader, subscriberAdapter)

	// starting services
	go rateMailer.Run()
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
