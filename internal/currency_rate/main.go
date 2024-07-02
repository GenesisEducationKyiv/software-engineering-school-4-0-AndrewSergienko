package currency_rate

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/currency_rate/app"
	"go_service/internal/infrastructure"
)

func NewApp(settings infrastructure.CurrencyAPISettings) *fiber.App {
	container := app.NewIoC(settings)
	return app.NewWebApp(container)
}
