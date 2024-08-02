package customers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/app"
	"gorm.io/gorm"
)

func NewApp(ctx context.Context, db *gorm.DB, conn jetstream.JetStream) *fiber.App {
	container := app.NewIoC(ctx, db, conn)
	return app.NewWebApp(container, adapters.NewNatsEventEmitter(ctx, conn))
}

func NewConsumer(ctx context.Context, db *gorm.DB, js jetstream.JetStream) app.Consumer {
	container := app.NewIoC(ctx, db, js)
	return app.NewConsumer(ctx, js, container)
}
