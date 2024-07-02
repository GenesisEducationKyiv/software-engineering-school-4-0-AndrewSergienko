package presentation

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go_service/internal"
	services2 "go_service/internal/currency_rate/services"
)

type GetCurrencyRate interface {
	Handle(data services2.GetCurrencyRateInputDTO) services2.GetCurrencyRateOutputDTO
}

type CurrencyHandlers struct {
	container internal.InteractorFactory
}

func NewCurrencyHandlers(container internal.InteractorFactory) CurrencyHandlers {
	return CurrencyHandlers{container}
}

func (ch *CurrencyHandlers) GetCurrency(c *fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to", "UAH")

	interactor := ch.container.GetCurrencyRate()
	result := interactor.Handle(services2.GetCurrencyRateInputDTO{From: from, To: to})

	if result.Err != nil {
		if errors.Is(result.Err, &internal.CurrencyNotExistsError{}) {
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
