package main

import (
	"fmt"
	globalInfrastructure "go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/notifier"
	"go_service/internal/notifier/infrastructure"
	"os"
	"path/filepath"
)

func main() {
	databaseSettings := globalInfrastructure.GetDatabaseSettings()
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

	db := database.InitDatabase(databaseSettings)
	task := notifier.NewTask(db, servicesAPISettings.CurrencyRate, servicesAPISettings.Subscriber, emailSettings)

	task.Run()
}
