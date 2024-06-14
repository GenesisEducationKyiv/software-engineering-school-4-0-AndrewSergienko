package main

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/adapters"
	"go_service/internal/api"
	"log"
	"os"
)

//
//func setupHandlers(cr common.CurrencyReader, sa api.SubscriberGateway) *http.ServeMux {
//	sm := http.NewServeMux()
//	sm.HandleFunc("/", api.GetCurrencyHandler(cr))
//	sm.HandleFunc("/subscribe", api.GetSubscribersHandler(sa))
//	return sm
//}

func FetchEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", name)
	}
	return value
}

func main() {
	currencyRateUrl := FetchEnv("CURRENCY_RATE_URL")
	currencyCode := FetchEnv("CURRENCY_CODE")
	//dbUser := fetchEnv("POSTGRES_USER")
	//dbPassword := fetchEnv("POSTGRES_PASSWORD")
	//dbName := fetchEnv("POSTGRES_DB")
	//dbPort := fetchEnv("DB_PORT")
	//dbHost := fetchEnv("DB_HOST")
	//email := fetchEnv("EMAIL")
	//emailPassword := fetchEnv("EMAIL_PASSWORD")

	//emailAdapter := adapters.EmailAdapter{
	//	Username: email,
	//	Auth:     smtp.PlainAuth("", email, emailPassword, "smtp.gmail.com"),
	//}

	//dsn := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//	dbHost,
	//	dbUser,
	//	dbPassword,
	//	dbName,
	//	dbPort,
	//)
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatalf("Database is not available. Error: %s", err)
	//	return
	//}

	//subscriberAdapter := adapters.SubscribersAdapter{Db: db}
	//schedulerAdapter := adapters.SchedulerDbAdapter{Db: db}
	currencyReader := adapters.APICurrencyReader{
		ApiUrl:       currencyRateUrl,
		CurrencyCode: currencyCode,
	}
	//sm := setupHandlers(&currencyReader, &subscriberAdapter)

	//rateMailer := internal.RateMailer{Es: emailAdapter, Sr: &subscriberAdapter, Sg: schedulerAdapter, Cr: &currencyReader}
	//go rateMailer.Run()

	app := fiber.New()

	currencyHandlers := api.NewCurrencyHandlers(&currencyReader)

	app.Get("/", currencyHandlers.GetCurrency)

	if app.Listen(":8080") != nil {
		return
	}
}
