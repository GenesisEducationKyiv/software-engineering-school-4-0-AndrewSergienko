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
	emailAdapter := adapters.EmailAdapter{Username: "sergienkoandre9922@gmail.com", Password: "ydbf qumr oxjx lhyh"}
	emailAdapter.CreateAuth()

	currencyRateUrl := fetchEnv("CURRENCY_RATE_URL")
	currencyCode := fetchEnv("CURRENCY_CODE")

	dsn := "host=localhost user=postgres password=postgres dbname=my-db port=5432 sslmode=disable"
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
