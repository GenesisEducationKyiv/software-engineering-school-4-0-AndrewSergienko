package currencyrate

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/currencyrate/adapters"
	"go_service/internal/rateservice/currencyrate/app"
	"go_service/internal/rateservice/infrastructure"
)

func NewApp(cacheClient *memcache.Client, settings infrastructure.CurrencyAPISettings) *fiber.App {
	cacheAdapter := adapters.NewCacheRateAdapter(cacheClient)
	container := app.NewIoC(cacheAdapter, settings)
	return app.NewWebApp(container, cacheAdapter)
}
