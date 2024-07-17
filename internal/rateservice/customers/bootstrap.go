package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"go_service/internal/rateservice/customers/app"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB, nc *nats.Conn) *fiber.App {
	container := app.NewIoC(db, nc)
	return app.NewWebApp(container)
}
