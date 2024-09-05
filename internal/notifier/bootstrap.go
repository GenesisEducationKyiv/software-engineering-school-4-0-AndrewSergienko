package notifier

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/notifier/adapters/scheduler"
	"go_service/internal/notifier/app"
	"go_service/internal/notifier/infrastructure"
	"gorm.io/gorm"
)

func NewTask(
	ctx context.Context,
	db *gorm.DB,
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
	emailSettings infrastructure.EmailSettings,
	conn jetstream.JetStream,
) app.RateNotifier {
	schedulerGateway := scheduler.NewScheduleAdapter(nil)
	container := app.NewIoC(ctx, db, currencyServiceSettings, emailSettings, conn)

	return app.NewRateNotifier(container, schedulerGateway)
}

func NewConsumer(ctx context.Context, db *gorm.DB, js jetstream.JetStream) app.Consumer {
	container := app.NewIoC(ctx, db, nil, infrastructure.EmailSettings{}, js)
	return app.NewConsumer(ctx, js, container)
}
