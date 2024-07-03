package currencyrate

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/currencyrate/app"
	"go_service/internal/infrastructure"
)

func NewApp(settings infrastructure.CurrencyAPISettings) *fiber.App {
	container := app.NewIoC(settings)
	return app.NewWebApp(container)
}
