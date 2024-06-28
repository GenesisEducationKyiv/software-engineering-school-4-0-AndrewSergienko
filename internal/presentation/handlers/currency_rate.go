package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/app"
	"go_service/internal/services"
)

type GetCurrencyRate interface {
	Handle(data services.GetCurrencyRateInputDTO) services.GetCurrencyRateOutputDTO
}

type CurrencyHandlers struct {
	container *app.IoC
}

func NewCurrencyHandlers(container *app.IoC) CurrencyHandlers {
	return CurrencyHandlers{container}
}

func (ch *CurrencyHandlers) GetCurrency(c *fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to", "UAH")

	interactor := ch.container.GetCurrencyRate()
	result := interactor.Handle(services.GetCurrencyRateInputDTO{From: from, To: to})

	if result.Err != nil {
		if errors.Is(result.Err, &services.CurrencyNotExistsError{}) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Currency has no found",
			})
		}
		return fiber.ErrInternalServerError
	}

	response := map[string]interface{}{
		"rate": result.Result,
	}

	return c.JSON(response)
}
