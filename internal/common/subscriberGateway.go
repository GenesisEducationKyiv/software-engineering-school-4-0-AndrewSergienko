package common

import (
	"go_service/internal/infrastructure/database/models"
)

type SubscriberReader interface {
	GetByEmail(email string) *models.Subscriber
}

type SubscriberListReader interface {
	GetAll() []models.Subscriber
}

type SubscriberWriter interface {
	Create(email string) error
}

type SubscriberDeleter interface {
	Delete(id int) error
}
