package main

import (
	"go_service/internal/rateservice/app"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/broker"
	"go_service/internal/rateservice/infrastructure/database"
	"log"
)

func main() {
	// app configuration
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()

	db := database.New(databaseSettings)

	conn, js := broker.New()
	defer broker.Finalize(conn)
	err := broker.NewStream(js, "events")
	if err != nil {
		panic(err)
	}

	// web app
	webApp := app.InitWebApp(db, js, currencyAPISettings)
	log.Printf("App failed with error: %v", webApp.Listen(":8080"))
}
