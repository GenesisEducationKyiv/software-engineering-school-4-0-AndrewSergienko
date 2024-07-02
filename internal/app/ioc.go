package app

import (
	"go_service/internal"
	adapters3 "go_service/internal/currency_rate/adapters"
	currencyRatePresentation "go_service/internal/currency_rate/presentation"
	services2 "go_service/internal/currency_rate/services"
	"go_service/internal/infrastructure"
	adapters4 "go_service/internal/notifier/adapters"
	"go_service/internal/notifier/services"
	adapters2 "go_service/internal/subscribers/adapters"
	subscribersPresentation "go_service/internal/subscribers/presentation"
	services3 "go_service/internal/subscribers/services"
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

func (ioc *IoC) Subscribe() subscribersPresentation.Interactor[services3.SubscribeInputDTO, services3.SubscribeOutputDTO] {
	return services3.NewSubscribe(ioc.subscriberAdapter)
}

func (ioc *IoC) GetCurrencyRate() currencyRatePresentation.Interactor[
	services2.GetCurrencyRateInputDTO,
	services2.GetCurrencyRateOutputDTO,
] {
	return services2.NewGetCurrencyRate(ioc.currencyRateFacade)
}

func (ioc *IoC) SendNotification() internal.Interactor[
	services.SendNotificationInputDTO,
	services.SendNotificationOutputDTO,
] {
	return services.NewSendNotification(ioc.emailAdapter, ioc.subscriberAdapter, ioc.currencyRateFacade)
}
