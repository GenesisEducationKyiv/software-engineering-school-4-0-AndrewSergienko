package subscriber

import (
	"go_service/internal/notifier/infrastructure/database/models"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func NewSubscriberAdapter(db *gorm.DB) Adapter {
	return Adapter{db: db}
}

func (sa *Adapter) GetAll() []models.Subscriber {
	var subscribers []models.Subscriber
	sa.db.Find(&subscribers)
	return subscribers
}

func (sa *Adapter) GetByEmail(email string) *models.Subscriber {
	var subscribers []models.Subscriber
	sa.db.Find(&subscribers, "email = ?", email)
	if len(subscribers) == 0 {
		return nil
	}
	return &subscribers[0]
}

func (sa *Adapter) Create(email string) error {
	subscriber := models.Subscriber{Email: email}
	return sa.db.Create(&subscriber).Error
}

func (sa *Adapter) Delete(email string) error {
	return sa.db.Where("email = ?", email).Delete(&models.Subscriber{}).Error
}
