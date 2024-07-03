package main

import (
	"go_service/internal/app"
	"go_service/internal/currency_rate"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/notifier"
	"go_service/internal/subscribers"
	"log"
)

func main() {
	// app configuration
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()
	emailSettings := infrastructure.GetEmailSettings()

	db := database.InitDatabase(databaseSettings)

	currencyApp := currency_rate.NewApp(currencyAPISettings)
	subscriberApp := subscribers.NewInternalApp(db)

	// background send mail task
	notifierTask := notifier.NewTask(db, currencyApp, subscriberApp, emailSettings)

	// web app
	webApp := app.InitWebApp(db, currencyAPISettings)

	// starting services
	go notifierTask.Run()
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
