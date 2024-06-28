package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type CurrencyGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type CurrencyHandlers struct {
	container dig.Container
}

func NewCurrencyHandlers(container dig.Container) CurrencyHandlers {
	return CurrencyHandlers{container}
}

func (ch *CurrencyHandlers) GetCurrency(c *fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to", "UAH")

	ch.container.Invoke(func() {

	})

	rate, err := ch.currencyGateway.GetCurrencyRate(from, to)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := map[string]interface{}{
		"rate": rate,
	}

	return c.JSON(response)
}
