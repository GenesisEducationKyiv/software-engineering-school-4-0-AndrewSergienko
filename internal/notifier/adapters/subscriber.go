package adapters

import (
	"go_service/internal/notifier/infrastructure/database/models"
	"gorm.io/gorm"
)

type SubscriberAdapter struct {
	db *gorm.DB
}

func NewSubscriberAdapter(db *gorm.DB) SubscriberAdapter {
	return SubscriberAdapter{db: db}
}

func (sa *SubscriberAdapter) GetAll() []models.Subscriber {
	var subscribers []models.Subscriber
	sa.db.Find(&subscribers)
	return subscribers
}
