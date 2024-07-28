package adapters

import (
	"go_service/internal/rateservice/customers/infrastructure/database/models"
	"gorm.io/gorm"
)

type SubscriberAdapter struct {
	db *gorm.DB
}

func NewSubscriberAdapter(db *gorm.DB) *SubscriberAdapter {
	return &SubscriberAdapter{db: db}
}

func (sa *SubscriberAdapter) GetByEmail(email string) *models.Customer {
	var subscribers []models.Customer
	sa.db.Find(&subscribers, "email = ?", email)
	if len(subscribers) == 0 {
		return nil
	}
	return &subscribers[0]
}

func (sa *SubscriberAdapter) Create(email string) error {
	subscriber := models.Customer{Email: email}
	err := sa.db.Create(&subscriber).Error
	return err
}

func (sa *SubscriberAdapter) DeleteByEmail(email string) error {
	return sa.db.Where("email = ?", email).Delete(&models.Customer{}).Error
}
