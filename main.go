package main

import (
	"go_service/src"
	"go_service/src/adapters"
	"go_service/src/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func setupHandlers(cr src.CurrencyReader, sa api.SubscriberGateway) *http.ServeMux {
	sm := http.NewServeMux()
	sm.HandleFunc("/", api.GetCurrencyHandler(cr))
	sm.HandleFunc("/subscribe", api.GetSubscribersHandler(sa))
	return sm
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=my-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	subscriberAdapter := adapters.SubscribersAdapter{Db: db}
	currencyReader := adapters.APICurrencyReader{ApiUrl: "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json", CurrencyCode: "USD"}
	sm := setupHandlers(&currencyReader, &subscriberAdapter)
	err = http.ListenAndServe(":8080", sm)
	if err != nil {
		return
	}
}
