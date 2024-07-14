package app

import (
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/services/createsubscriber"
	"go_service/internal/notifier/services/deletesubscriber"
	"go_service/internal/notifier/services/sendnotification"
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

func (ioc *IoC) SendNotification() Interactor[sendnotification.InputData, sendnotification.OutputData] {
	return sendnotification.NewSendNotification(ioc.emailAdapter, &ioc.subscriberAdapter, ioc.currencyRateAdapter)
}

func (ioc *IoC) CreateSubscriber() Interactor[createsubscriber.InputData, createsubscriber.OutputData] {
	return createsubscriber.NewCreateSubscriber(&ioc.subscriberAdapter)
}

func (ioc *IoC) DeleteSubscriber() Interactor[deletesubscriber.InputData, deletesubscriber.OutputData] {
	return deletesubscriber.NewDeleteSubscriber(&ioc.subscriberAdapter)
}
