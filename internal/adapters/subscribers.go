package adapters

import (
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type SubscriberAdapter struct {
	Db *gorm.DB
}

func GetSubscribersAdapter(db *gorm.DB) *SubscriberAdapter {
	return &SubscriberAdapter{Db: db}
}

func (sa *SubscriberAdapter) GetByEmail(email string) *models.Subscriber {
	var subscribers []models.Subscriber
	sa.Db.Find(&subscribers, "email = ?", email)
	if len(subscribers) == 0 {
		return nil
	}
	return &subscribers[0]
}

func (sa *SubscriberAdapter) Create(email string) error {
	subscriber := models.Subscriber{Email: email}
	return sa.Db.Create(&subscriber).Error
}

func (sa *SubscriberAdapter) Delete(id int) error {
	return sa.Db.Delete(&models.Subscriber{}, id).Error
}

func (sa *SubscriberAdapter) GetAll() []models.Subscriber {
	var subscribers []models.Subscriber
	sa.Db.Find(&subscribers)
	return subscribers
}
