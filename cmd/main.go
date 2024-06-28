package main

import (
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

	container := app.NewIoC(db, emailSettings, currencyAPISettings)

	// background send mail task
	//rateMailer := app.InitRateMailer(emailAdapter, subscriberAdapter, schedulerAdapter, currencyReader)

	// web app
	webApp := app.InitWebApp(container)

	// starting services
	//go rateMailer.Run()
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
