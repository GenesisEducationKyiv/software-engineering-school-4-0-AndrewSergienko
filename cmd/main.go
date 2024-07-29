package main

import (
	"context"
	"go_service/internal/rateservice/app"
	"go_service/internal/rateservice/customers"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/broker"
	"go_service/internal/rateservice/infrastructure/database"
	"log"
	"time"
)

func main() {
	// app configuration
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := database.New(databaseSettings)

	conn, js := broker.New()
	defer broker.Finalize(conn)
	_, err := broker.NewStream(ctx, js, "events")
	if err != nil {
		panic(err)
	}

	consumer := customers.NewConsumer(ctx, db, js)
	consumeContext := consumer.Run()
	defer consumeContext.Stop()

	// web app
	webApp := app.InitWebApp(ctx, db, js, currencyAPISettings)
	log.Printf("App failed with error: %v", webApp.Listen(":8080"))
}
