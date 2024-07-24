package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/app"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB, conn nats.JetStreamContext) *fiber.App {
	container := app.NewIoC(db, conn)
	return app.NewWebApp(container, adapters.NewNatsEventEmitter(conn))
}
