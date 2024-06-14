package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/common"
	"net/http"
)

func GetCurrencyHandler(cr common.CurrencyReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err := fmt.Fprintf(w, "Method Not Allowed")
			if err != nil {
				return
			}
		}

		rate, err := cr.GetUSDCurrencyRate()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"rate": rate,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}
}

type CurrencyHandlers struct {
	currencyGateway common.CurrencyReader
}

func InitCurrencyHandlers(currencyGateway common.CurrencyReader) CurrencyHandlers {
	return CurrencyHandlers{currencyGateway}
}

func (ch *CurrencyHandlers) GetCurrency(c *fiber.Ctx) error {
	rate, err := ch.currencyGateway.GetUSDCurrencyRate()

	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := map[string]interface{}{
		"rate": rate,
	}

	return c.JSON(response)
}
