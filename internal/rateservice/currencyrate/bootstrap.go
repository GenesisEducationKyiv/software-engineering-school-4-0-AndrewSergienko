package currencyrate

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/currencyrate/app"
	"go_service/internal/rateservice/infrastructure"
)

func NewApp(settings infrastructure.CurrencyAPISettings) *fiber.App {
	container := app.NewIoC(settings)
	return app.NewWebApp(container)
}
