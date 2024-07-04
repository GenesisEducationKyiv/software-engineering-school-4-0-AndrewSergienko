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

	db := database.InitDatabase(databaseSettings)

	// web app
	webApp := app.InitWebApp(db, currencyAPISettings)
	log.Fatalf("App failed with error: %v", webApp.Listen(":8080"))
}
