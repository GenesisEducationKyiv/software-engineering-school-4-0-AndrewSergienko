package main

import (
	globalInfrastructure "go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/notifier"
	"go_service/internal/notifier/infrastructure"
)

func runEmailService() {
	databaseSettings := globalInfrastructure.GetDatabaseSettings()
	emailSettings := infrastructure.GetEmailSettings()

	currencyAPISettings := infrastructure.GetCurrencyServiceAPISettings()
	subscriberAPISettings := infrastructure.GetSubscriberServiceAPISettings()

	db := database.InitDatabase(databaseSettings)
	task := notifier.NewTask(db, currencyAPISettings, subscriberAPISettings, emailSettings)
	go task.Run()
}
