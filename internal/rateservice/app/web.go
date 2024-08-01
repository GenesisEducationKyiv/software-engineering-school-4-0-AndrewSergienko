package app

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/prometheus/client_golang/prometheus"
	"go_service/internal/rateservice/currencyrate"
	"go_service/internal/rateservice/customers"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/metrics"
	"gorm.io/gorm"
	"net/http"
)

func InitWebApp(
	ctx context.Context,
	db *gorm.DB,
	conn jetstream.JetStream,
	cacheClient *memcache.Client,
	apiSettings infrastructure.CurrencyAPISettings,
) *fiber.App {
	app := fiber.New()
	// app.Use(swagger.New())

	app.Use(func(c *fiber.Ctx) error {
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			metrics.RequestsDuration.WithLabelValues(c.Path(), c.Method(), http.StatusText(c.Response().StatusCode())).Observe(v)
		}))
		defer timer.ObserveDuration()

		err := c.Next()

		metrics.RequestsTotal.WithLabelValues(c.Path(), c.Method(), http.StatusText(c.Response().StatusCode())).Inc()

		return err
	})

	subscribersApp := customers.NewApp(ctx, db, conn)
	currencyRateApp := currencyrate.NewApp(cacheClient, apiSettings)

	app.Mount("/customers/", subscribersApp)
	app.Mount("/rates/", currencyRateApp)

	return app
}
