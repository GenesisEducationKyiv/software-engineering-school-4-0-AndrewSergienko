package services

import (
	"go_service/internal/infrastructure/database/models"
)

type SubscribeInputDTO struct {
	Email string
}

type SubscribeOutputDTO struct {
	Err error
}

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

func (s *Subscribe) Handle(data SubscribeInputDTO) SubscribeOutputDTO {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return SubscribeOutputDTO{Err: &EmailConflictError{Email: data.Email}}
	}

	return SubscribeOutputDTO{s.subscriberGateway.Create(data.Email)}
}
