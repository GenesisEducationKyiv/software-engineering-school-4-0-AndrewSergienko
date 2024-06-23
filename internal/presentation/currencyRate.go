package presentation

import (
	"github.com/gofiber/fiber/v2"
)

type CurrencyGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type CurrencyHandlers struct {
	currencyGateway CurrencyGateway
}

func NewCurrencyHandlers(currencyGateway CurrencyGateway) CurrencyHandlers {
	return CurrencyHandlers{currencyGateway}
}

func (ch *CurrencyHandlers) GetCurrency(c *fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to", "UAH")

	rate, err := ch.currencyGateway.GetCurrencyRate(from, to)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := map[string]interface{}{
		"rate": rate,
	}

	return c.JSON(response)
}
