package presentation

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/common"
)

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
