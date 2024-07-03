package notifier

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/infrastructure"
	"go_service/internal/notifier/adapters"
	"go_service/internal/notifier/app"
	"gorm.io/gorm"
)

func NewTask(
	db *gorm.DB,
	currencyApp *fiber.App,
	subscriberApp *fiber.App,
	emailSettings infrastructure.EmailSettings,
) app.RateMailer {
	schedulerGateway := adapters.NewScheduleDBAdapter(db)
	container := app.NewIoC(currencyApp, subscriberApp, emailSettings)

	return app.NewRateMailer(container, schedulerGateway)
}
