package notifier

import (
	"github.com/nats-io/nats.go"
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/app"
	"go_service/internal/notifier/infrastructure"
	"gorm.io/gorm"
)

func NewTask(
	db *gorm.DB,
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
	emailSettings infrastructure.EmailSettings,
) app.RateNotifier {
	schedulerGateway := adapters.NewScheduleAdapter()
	container := app.NewIoC(db, currencyServiceSettings, emailSettings)

	return app.NewRateNotifier(container, schedulerGateway)
}

func NewConsumer(db *gorm.DB, nc *nats.Conn) app.Consumer {
	container := app.NewIoC(db, nil, infrastructure.EmailSettings{})
	return app.NewConsumer(nc, container)
}
