package main

import (
	"fmt"
	"go_service/src"
	"go_service/src/adapters"
	"go_service/src/api"
	"go_service/src/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func setupHandlers(cr common.CurrencyReader, sa api.SubscriberGateway) *http.ServeMux {
	sm := http.NewServeMux()
	sm.HandleFunc("/", api.GetCurrencyHandler(cr))
	sm.HandleFunc("/subscribe", api.GetSubscribersHandler(sa))
	return sm
}

func fetchEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatal(fmt.Sprintf("Environment variable %s is not set", name))
	}
	return value
}

func main() {
	currencyRateUrl := fetchEnv("CURRENCY_RATE_URL")
	currencyCode := fetchEnv("CURRENCY_CODE")
	dbUser := fetchEnv("POSTGRES_USER")
	dbPassword := fetchEnv("POSTGRES_PASSWORD")
	dbName := fetchEnv("POSTGRES_DB")
	dbPort := fetchEnv("DB_PORT")
	dbHost := fetchEnv("DB_HOST")
	email := fetchEnv("EMAIL")
	emailPassword := fetchEnv("EMAIL_PASSWORD")

	emailAdapter := adapters.EmailAdapter{Username: email, Password: emailPassword}
	emailAdapter.CreateAuth()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	subscriberAdapter := adapters.SubscribersAdapter{Db: db}
	schedulerAdapter := adapters.SchedulerDbAdapter{Db: db}
	currencyReader := adapters.APICurrencyReader{
		ApiUrl:       currencyRateUrl,
		CurrencyCode: currencyCode,
	}
	sm := setupHandlers(&currencyReader, &subscriberAdapter)

	rateMailer := src.RateMailer{Es: emailAdapter, Sr: &subscriberAdapter, Sg: schedulerAdapter, Cr: &currencyReader}
	go rateMailer.Run()

	err = http.ListenAndServe(":8080", sm)
	if err != nil {
		return
	}
}
