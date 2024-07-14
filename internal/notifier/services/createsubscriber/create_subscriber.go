package createsubscriber

import (
	"go_service/internal/notifier/domain"
	"go_service/internal/notifier/infrastructure/database/models"
)

type InputData struct {
	Email string
}

type OutputData struct {
	Err error
}

type SubscriberGateway interface {
	Create(email string) error
	GetByEmail(email string) *models.Subscriber
}

type CreateSubscriber struct {
	subscriberGateway SubscriberGateway
}

func NewCreateSubscriber(sg SubscriberGateway) *CreateSubscriber {
	return &CreateSubscriber{subscriberGateway: sg}
}

func (s *CreateSubscriber) Handle(data InputData) OutputData {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return OutputData{Err: &domain.EmailConflictError{Email: data.Email}}
	}

	return OutputData{s.subscriberGateway.Create(data.Email)}
}
