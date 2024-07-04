package notifier

import (
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/app"
	"go_service/internal/notifier/infrastructure"
	"gorm.io/gorm"
)

func NewTask(
	db *gorm.DB,
	currencyServiceSettings infrastructure.CurrencyServiceAPISettings,
	subscriberServiceSettings infrastructure.SubscriberServiceAPISettings,
	emailSettings infrastructure.EmailSettings,
) app.RateMailer {
	schedulerGateway := adapters.NewScheduleDBAdapter(db)
	container := app.NewIoC(currencyServiceSettings, subscriberServiceSettings, emailSettings)

	return app.NewRateMailer(container, schedulerGateway)
}
