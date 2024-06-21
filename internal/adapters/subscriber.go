package adapters

import (
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type SubscriberAdapter struct {
	db *gorm.DB
}

func NewSubscribersAdapter(db *gorm.DB) *SubscriberAdapter {
	return &SubscriberAdapter{db: db}
}

func (sa *SubscriberAdapter) GetByEmail(email string) *models.Subscriber {
	var subscribers []models.Subscriber
	sa.db.Find(&subscribers, "email = ?", email)
	if len(subscribers) == 0 {
		return nil
	}
	return &subscribers[0]
}

func (sa *SubscriberAdapter) Create(email string) error {
	subscriber := models.Subscriber{Email: email}
	return sa.db.Create(&subscriber).Error
}

func (sa *SubscriberAdapter) Delete(id int) error {
	return sa.db.Delete(&models.Subscriber{}, id).Error
}

func (sa *SubscriberAdapter) GetAll() []models.Subscriber {
	var subscribers []models.Subscriber
	sa.db.Find(&subscribers)
	return subscribers
}
