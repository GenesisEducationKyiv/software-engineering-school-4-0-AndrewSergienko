package app

import (
	adapters3 "go_service/internal/currency_rate/adapters"
	"go_service/internal/infrastructure"
	"go_service/internal/notifier"
	adapters4 "go_service/internal/notifier/adapters"
	"go_service/internal/notifier/services"
	adapters2 "go_service/internal/subscribers/adapters"
	"gorm.io/gorm"
)

type IoC struct {
	subscriberAdapter  *adapters2.SubscriberAdapter
	scheduleAdapter    *adapters4.ScheduleDBAdapter
	emailAdapter       adapters4.EmailAdapter
	currencyRateFacade *adapters3.APIReaderFacade
}

func NewIoC(
	db *gorm.DB,
	emailSettings infrastructure.EmailSettings,
	currencyAPISettings infrastructure.CurrencyAPISettings,
) *IoC {
	readers := adapters3.CreateReaders(currencyAPISettings)
	return &IoC{
		subscriberAdapter:  adapters2.NewSubscriberAdapter(db),
		scheduleAdapter:    adapters4.NewScheduleDBAdapter(db),
		emailAdapter:       adapters4.NewEmailAdapter(emailSettings),
		currencyRateFacade: adapters3.NewAPIReaderFacade(readers),
	}
}

func (ioc *IoC) SendNotification() notifier.Interactor[
	services.SendNotificationInputDTO,
	services.SendNotificationOutputDTO,
] {
	return services.NewSendNotification(ioc.emailAdapter, ioc.subscriberAdapter, ioc.currencyRateFacade)
}
