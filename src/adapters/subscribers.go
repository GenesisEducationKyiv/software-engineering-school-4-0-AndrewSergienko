package adapters

import (
	"go_service/src"
	"gorm.io/gorm"
)

type SubscribersAdapter struct {
	Db *gorm.DB
}

func (sa *SubscribersAdapter) GetByEmail(email string) *src.Subscriber {
	var subscribers []src.Subscriber
	sa.Db.Find(&subscribers, "email = ?", email)
	if len(subscribers) == 0 {
		return nil
	}
	return &subscribers[0]
}

func (sa *SubscribersAdapter) Create(email string) error {
	subscriber := src.Subscriber{Email: email}
	result := sa.Db.Create(&subscriber)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sa *SubscribersAdapter) Delete(id int) {
	sa.Db.Delete(&src.Subscriber{}, id)
}

func (sa *SubscribersAdapter) GetAll() []src.Subscriber {
	var subscribers []src.Subscriber
	sa.Db.Find(&subscribers)
	return subscribers
}
