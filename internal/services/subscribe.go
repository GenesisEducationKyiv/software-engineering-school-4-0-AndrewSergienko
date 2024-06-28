package services

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/infrastructure/database/models"
)

type SubscriberGateway interface {
	Create(email string) error
	Delete(id int) error
	GetByEmail(email string) *models.Subscriber
}

type Subscribe struct {
	subscriberGateway SubscriberGateway
}

func NewSubscribe(sg SubscriberGateway) *Subscribe {
	return &Subscribe{subscriberGateway: sg}
}

func (s *Subscribe) Handle(data string) error {
	if sh.subscriberGateway.GetByEmail(requestData.Email) != nil {
		return fiber.ErrConflict
	}

	if sh.subscriberGateway.Create(requestData.Email) != nil {
		return fiber.ErrInternalServerError
	}
	return s.subscriberGateway.Create(email)
}
