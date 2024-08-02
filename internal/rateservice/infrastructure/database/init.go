package database

import (
	"fmt"
	"go_service/internal/rateservice/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(settings infrastructure.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		settings.Host,
		settings.User,
		settings.Password,
		settings.Database,
		settings.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
