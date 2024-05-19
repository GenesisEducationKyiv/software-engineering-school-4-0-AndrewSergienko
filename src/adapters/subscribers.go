package adapters

import (
	"go_service/src/models"
	"gorm.io/gorm"
)

type SubscribersAdapter struct {
	Db *gorm.DB
}

func (sa *SubscribersAdapter) GetByEmail(email string) *models.Subscriber {
	var subscribers []models.Subscriber
	sa.Db.Find(&subscribers, "email = ?", email)
	if len(subscribers) == 0 {
		return nil
	}
	return &subscribers[0]
}

func (sa *SubscribersAdapter) Create(email string) error {
	subscriber := models.Subscriber{Email: email}
	result := sa.Db.Create(&subscriber)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sa *SubscribersAdapter) Delete(id int) {
	sa.Db.Delete(&models.Subscriber{}, id)
}

func (sa *SubscribersAdapter) GetAll() []models.Subscriber {
	var subscribers []models.Subscriber
	sa.Db.Find(&subscribers)
	return subscribers
}
