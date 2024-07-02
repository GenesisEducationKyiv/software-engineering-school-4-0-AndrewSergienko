package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/notifier"
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/services"
)

type IoC struct {
	currencyRateAdapter adapters.SubscriberAdapter
	emailAdapter        adapters.EmailAdapter
}

func NewIoC(currencyApp *fiber.App) *IoC {
	return &IoC{currencyRateAdapter: adapters.NewCurrencyRateAdapter(currencyApp)}
}

func (ioc *IoC) SendNotification() notifier.Interactor[
	services.SendNotificationInputDTO,
	services.SendNotificationOutputDTO,
] {
	return services.NewSendNotification(ioc.emailAdapter, ioc.subscriberAdapter, ioc.currencyRateAdapter)
}
