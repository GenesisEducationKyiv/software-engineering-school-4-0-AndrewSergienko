package main

import (
	"context"
	"fmt"
	"go_service/internal/rateservice/app"
	"go_service/internal/rateservice/customers"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/broker"
	"go_service/internal/rateservice/infrastructure/cache"
	"go_service/internal/rateservice/infrastructure/database"
	"go_service/internal/rateservice/infrastructure/metrics"
	"log/slog"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// app configuration
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	databaseSettings := infrastructure.GetDatabaseSettings()
	brokerSettings := infrastructure.GetBrokerSettings()
	cahceSettings := infrastructure.GetCacheSettings()

	ctx := context.Background()

	db, err := database.New(databaseSettings)
	if err != nil {
		slog.Error(fmt.Sprintf("Database is not available. Error: %s", err))
		return
	}

	conn, js, err := broker.New(brokerSettings)
	if err != nil {
		slog.Error(fmt.Sprintf("Message broker is not available. Error: %s", err))
		return
	}

	defer broker.Finalize(conn)
	_, err = broker.NewStream(ctx, js, "events")
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create JetStream stream events. Error: %s", err))
		return
	}

	consumer := customers.NewConsumer(ctx, db, js)
	consumeContext, err := consumer.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("Error starting consumer: %s", err))
		return
	}
	defer func() {
		consumeContext.Stop()
		slog.Info("Consumer stopped")
	}()

	cacheClient := cache.New(cahceSettings)
	defer cacheClient.Close()

	go metrics.RunServer()

	// web app
	webApp := app.InitWebApp(ctx, db, js, cacheClient, currencyAPISettings)
	slog.Error(fmt.Sprintf("App failed with error: %v", webApp.Listen(":8080")))
}
