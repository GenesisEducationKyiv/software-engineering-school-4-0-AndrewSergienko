package createcustomer

import (
	"go_service/internal/rateservice/customers/infrastructure/database/models"
	"go_service/internal/rateservice/customers/services"
)

type InputData struct {
	Email         string
	TransactionID *string
	IsRollback    bool
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

type CreateCustomer struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway, em EventEmitter) *CreateCustomer {
	return &CreateCustomer{subscriberGateway: sg, eventEmitter: em}
}

func (s *CreateCustomer) Handle(data InputData) OutputData {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return OutputData{Err: &services.EmailConflictError{Email: data.Email}}
	}
	err := s.subscriberGateway.Create(data.Email)

	var event string
	if err == nil {
		if data.IsRollback {
			event = "UserCreatedRollback"
		} else {
			event = "UserCreated"
		}
	} else if !data.IsRollback {
		event = "UserCreatedError"
	}
	_ = s.eventEmitter.Emit(event, map[string]interface{}{"Email": data.Email}, data.TransactionID)
	// TODO: Add transactional outbox pattern

	return OutputData{err}
}
