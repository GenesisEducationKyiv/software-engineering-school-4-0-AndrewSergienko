package notifier

import (
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/app"
	"go_service/internal/notifier/infrastructure"
	"gorm.io/gorm"
)

func NewTask(
	db *gorm.DB,
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
	emailSettings infrastructure.EmailSettings,
) app.RateMailer {
	schedulerGateway := adapters.NewScheduleDBAdapter(db)
	container := app.NewIoC(db, currencyServiceSettings, emailSettings)

	return app.NewRateMailer(container, schedulerGateway)
}
