package app

import (
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/services"
	"gorm.io/gorm"
)

type IoC struct {
	emailAdapter        adapters.EmailAdapter
	currencyRateAdapter adapters.CurrencyRateAdapter
	subscriberAdapter   adapters.SubscriberAdapter
}

func NewIoC(
	db *gorm.DB,
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
	emailSettings infrastructure.EmailSettings,
) *IoC {
	return &IoC{
		currencyRateAdapter: adapters.NewCurrencyRateAdapter(currencyServiceSettings),
		subscriberAdapter:   adapters.NewSubscriberAdapter(db),
		emailAdapter:        adapters.NewEmailAdapter(emailSettings),
	}
}

func (ioc *IoC) SendNotification() Interactor[
	services.SendNotificationInputDTO,
	services.SendNotificationOutputDTO,
] {
	return services.NewSendNotification(ioc.emailAdapter, &ioc.subscriberAdapter, ioc.currencyRateAdapter)
}
