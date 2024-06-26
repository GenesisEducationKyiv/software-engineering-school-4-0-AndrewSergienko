package database

import (
	"fmt"
	"go_service/internal/infrastructure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase(settings infrastructure.DatabaseSettings) *gorm.DB {
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
		log.Fatalf("Database is not available. Error: %s", err)
	}
	return db
}
