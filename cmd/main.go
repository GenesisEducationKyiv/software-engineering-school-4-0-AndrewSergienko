package main

import (
	"go_service/internal/rateservice/app"
	currencyRateInfrastructure "go_service/internal/rateservice/currencyrate/infrastructure"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/database"
	"log"
)

func main() {
	// app configuration
	currencyAPISettings := currencyRateInfrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()

	db := database.New(databaseSettings)

	// web app
	webApp := app.InitWebApp(db, currencyAPISettings)
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
