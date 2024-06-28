package app

import (
	"go.uber.org/dig"
	"go_service/internal/adapters"
	"go_service/internal/adapters/currencyrate"
	"go_service/internal/infrastructure"
	"gorm.io/gorm"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(InputDTO) OutputDTO
}

func SetupContainer(
	db *gorm.DB,
	emailSettings infrastructure.EmailSettings,
	currencyAPISettings infrastructure.CurrencyAPISettings,
) {
	container := dig.New()

	container.Provide(func() *adapters.SubscriberAdapter {
		return adapters.NewSubscribersAdapter(db)
	})

	container.Provide(func() *adapters.ScheduleDBAdapter {
		return adapters.NewScheduleDBAdapter(db)
	})

	container.Provide(func() adapters.EmailAdapter {
		return adapters.NewEmailAdapter(emailSettings)
	})

	container.Provide(func() *currencyrate.APIReaderFacade {
		readers := currencyrate.CreateReaders(currencyAPISettings)
		return currencyrate.NewAPIReaderFacade(readers)
	})

	container.Provide(func() Interactor[string, float32] {})
}
