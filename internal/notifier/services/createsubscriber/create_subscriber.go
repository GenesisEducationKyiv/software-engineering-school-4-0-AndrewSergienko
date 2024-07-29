package createsubscriber

import (
	"go_service/internal/notifier/domain"
	"go_service/internal/notifier/infrastructure/database/models"
)

type InputData struct {
	Email         string
	TransactionID *string
}

type OutputData struct {
	Err error
}

type SubscriberGateway interface {
	Create(email string) error
	GetByEmail(email string) *models.Subscriber
}

type EventEmitter interface {
	Emit(name string, data map[string]interface{}, transactionID *string) error
}

type CreateSubscriber struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway, em EventEmitter) *CreateSubscriber {
	return &CreateSubscriber{subscriberGateway: sg, eventEmitter: em}
}

func (s *CreateSubscriber) Handle(data InputData) OutputData {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return OutputData{Err: &domain.EmailConflictError{Email: data.Email}}
	}

	err := s.subscriberGateway.Create(data.Email)
	var event string
	if err == nil {
		event = "SubscriberCreated"
	} else {
		event = "SubscriberCreatedError"
	}
	_ = s.eventEmitter.Emit(
		event,
		map[string]interface{}{"email": data.Email},
		data.TransactionID,
	)
	// TODO: Add transactional outbox pattern
	return OutputData{}
}
