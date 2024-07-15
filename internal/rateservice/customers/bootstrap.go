package customers

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/customers/app"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB) *fiber.App {
	container := app.NewIoC(db)
	return app.NewWebApp(container)
}
