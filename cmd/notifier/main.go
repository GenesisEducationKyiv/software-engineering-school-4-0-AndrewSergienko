package main

import (
	"fmt"
	"go_service/internal/notifier"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/broker"
	"go_service/internal/notifier/infrastructure/database"
	"os"
	"path/filepath"
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

	db := database.New(databaseSettings)

	conn, js := broker.New()
	defer broker.Finalize(conn)
	err = broker.NewStream(js, "events")
	if err != nil {
		panic(err)
	}

	task := notifier.NewTask(db, servicesAPISettings.CurrencyRate, emailSettings, js)
	consumer := notifier.NewConsumer(db, js)

	task.Run()
	consumer.Run()
	select {}
}
