package database

import (
	"fmt"
	"go_service/internal/notifier/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func New(settings infrastructure.DatabaseSettings) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		settings.Host,
		settings.User,
		settings.Password,
		settings.Database,
		settings.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error(fmt.Sprintf("Database is not available. Error: %s", err))
		os.Exit(1)
	}
	return db
}
