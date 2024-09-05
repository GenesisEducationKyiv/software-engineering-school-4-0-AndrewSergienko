package main

import (
	"context"
	"fmt"
	"go_service/internal/notifier"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/broker"
	"go_service/internal/notifier/infrastructure/database"
	"go_service/internal/notifier/infrastructure/metrics"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	databaseSettings := infrastructure.GetDatabaseSettings()
	emailSettings := infrastructure.GetEmailSettings()
	brokerSettings := infrastructure.GetBrokerSettings()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	configPath := filepath.Join(cwd, "conf", "config.toml")
	servicesAPISettings, err := infrastructure.GetServicesAPISettings(configPath)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	db, err := database.New(databaseSettings)
	if err != nil {
		slog.Error(fmt.Sprintf("Database is not available. Error: %s", err))
		return
	}

	conn, js, err := broker.New(brokerSettings)
	if err != nil {
		slog.Error(fmt.Sprintf("NATS is not available. Error: %s", err))
		return
	}

	defer broker.Finalize(conn)
	_, err = broker.NewStream(ctx, js, "events")
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create JetStream stream events. Error: %s", err))
		return
	}

	task := notifier.NewTask(ctx, db, servicesAPISettings.CurrencyRate, emailSettings, js)
	consumer := notifier.NewConsumer(ctx, db, js)

	taskContext := task.Run()
	defer taskContext.Stop()

	consumeContext, err := consumer.Run()
	if err != nil {
		log.Printf("Error starting consumer: %s", err)
		return
	}
	defer func() {
		consumeContext.Stop()
		slog.Info("Consumer stopped")
	}()

	go metrics.RunServer()

	select {}
}
