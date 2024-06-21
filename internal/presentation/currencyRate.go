package presentation

import (
	"github.com/gofiber/fiber/v2"
)

type CurrencyGateway interface {
	GetUSDCurrencyRate() (float32, error)
}

type CurrencyHandlers struct {
	currencyGateway CurrencyGateway
}

func NewCurrencyHandlers(currencyGateway CurrencyGateway) CurrencyHandlers {
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
