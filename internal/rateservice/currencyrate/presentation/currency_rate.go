package presentation

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/currencyrate/adapters"
	"go_service/internal/rateservice/currencyrate/services"
	"log/slog"
)

type GetCurrencyRate interface {
	Handle(data services.GetCurrencyRateInputDTO) services.GetCurrencyRateOutputDTO
}

type CurrencyHandlers struct {
	container    InteractorFactory
	cacheAdapter adapters.CacheRateAdapter
}

func NewCurrencyHandlers(container InteractorFactory, cacheAdapter adapters.CacheRateAdapter) CurrencyHandlers {
	return CurrencyHandlers{container, cacheAdapter}
}

// GetCurrency @Summary Get currency rate
// @Description Get the exchange rate from one currency to another
// @Tags currency
// @Accept json
// @Produce json
// @Param from query string true "Source currency code"
// @Param to query string false "Target currency code" default(UAH)
// @Success 200 {object} map[string]interface{} "rate"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /currency [get]
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

	err := ch.cacheAdapter.SetCurrencyRate(from, to, result.Result)
	if err != nil {
		slog.Warn(fmt.Sprintf("Error set currency rate to cache. Error: %v", err))
	}

	response := map[string]interface{}{
		"rate": result.Result,
	}

	return c.JSON(response)
}
