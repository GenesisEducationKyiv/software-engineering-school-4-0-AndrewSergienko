package app

import (
	"go_service/internal/adapters"
	"go_service/internal/adapters/currencyrate"
	"go_service/internal/infrastructure"
	"go_service/internal/services"
	"gorm.io/gorm"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type IoC struct {
	subscriberAdapter  *adapters.SubscriberAdapter
	scheduleAdapter    *adapters.ScheduleDBAdapter
	emailAdapter       adapters.EmailAdapter
	currencyRateFacade *currencyrate.APIReaderFacade
}

func NewIoC(
	db *gorm.DB,
	emailSettings infrastructure.EmailSettings,
	currencyAPISettings infrastructure.CurrencyAPISettings,
) *IoC {
	readers := currencyrate.CreateReaders(currencyAPISettings)
	return &IoC{
		subscriberAdapter:  adapters.NewSubscribersAdapter(db),
		scheduleAdapter:    adapters.NewScheduleDBAdapter(db),
		emailAdapter:       adapters.NewEmailAdapter(emailSettings),
		currencyRateFacade: currencyrate.NewAPIReaderFacade(readers),
	}
}

func (ioc *IoC) Subscribe() Interactor[services.SubscribeInputDTO, services.SubscribeOutputDTO] {
	return services.NewSubscribe(ioc.subscriberAdapter)
}

func (ioc *IoC) GetCurrencyRate() *services.GetCurrencyRate {
	return services.NewGetCurrencyRate(ioc.currencyRateFacade)
}
