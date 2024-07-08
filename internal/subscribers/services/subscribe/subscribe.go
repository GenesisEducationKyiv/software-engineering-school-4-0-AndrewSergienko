package subscribe

import (
	"go_service/internal/subscribers/infrastructure/database/models"
	"go_service/internal/subscribers/services"
)

type InputDTO struct {
	Email string
}

type OutputDTO struct {
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

func (s *Subscribe) Handle(data InputDTO) OutputDTO {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return OutputDTO{Err: &services.EmailConflictError{Email: data.Email}}
	}

	return OutputDTO{s.subscriberGateway.Create(data.Email)}
}
