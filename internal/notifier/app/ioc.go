package app

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/infrastructure"
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/services"
)

type IoC struct {
	emailAdapter        adapters.EmailAdapter
	currencyRateAdapter adapters.CurrencyRateAdapter
	subscriberAdapter   adapters.SubscriberAdapter
}

func NewIoC(currencyApp *fiber.App, subscriberApp *fiber.App, emailSettings infrastructure.EmailSettings) *IoC {
	return &IoC{
		currencyRateAdapter: adapters.NewCurrencyRateAdapter(currencyApp),
		subscriberAdapter:   adapters.NewSubscriberAdapter(subscriberApp),
		emailAdapter:        adapters.NewEmailAdapter(emailSettings),
	}
}

func (ioc *IoC) SendNotification() Interactor[
	services.SendNotificationInputDTO,
	services.SendNotificationOutputDTO,
] {
	return services.NewSendNotification(ioc.emailAdapter, ioc.subscriberAdapter, ioc.currencyRateAdapter)
}
