package main

import (
	"go_service/src"
	"net/http"
)

func setupHandlers(cr src.CurrencyReader) *http.ServeMux {
	sm := http.NewServeMux()
	sm.HandleFunc("/", src.GetCurrencyHandler(cr))
	return sm
}

func main() {
	currencyReader := src.APICurrencyReader{ApiUrl: "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json", CurrencyCode: "USD"}
	sm := setupHandlers(&currencyReader)
	err := http.ListenAndServe(":8080", sm)
	if err != nil {
		return
	}
}
