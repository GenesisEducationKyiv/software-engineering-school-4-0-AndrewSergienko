package main

import (
	"context"
	"fmt"
	"go_service/internal/notifier"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/broker"
	"go_service/internal/notifier/infrastructure/database"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	databaseSettings := infrastructure.GetDatabaseSettings()
	emailSettings := infrastructure.GetEmailSettings()

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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := database.New(databaseSettings)

	conn, js := broker.New()
	defer broker.Finalize(conn)
	_, err = broker.NewStream(ctx, js, "events")
	if err != nil {
		panic(err)
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
	defer consumeContext.Stop()
	select {}
}
