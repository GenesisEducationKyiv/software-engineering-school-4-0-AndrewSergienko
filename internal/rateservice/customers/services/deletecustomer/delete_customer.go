package deletecustomer

import (
	"go_service/internal/rateservice/customers/infrastructure/database/models"
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
	GetByEmail(email string) *models.Subscriber
	DeleteByEmail(email string) error
}

type EventEmitter interface {
	Emit(name string, data map[string]interface{}, transactionID *string) error
}

type DeleteCustomer struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway, em EventEmitter) *DeleteCustomer {
	return &DeleteCustomer{subscriberGateway: sg, eventEmitter: em}
}

func (s *DeleteCustomer) Handle(data InputData) OutputData {
	err := s.subscriberGateway.DeleteByEmail(data.Email)

	var event string
	if err == nil {
		if data.IsRollback {
			event = "UserDeletedRollback"
		} else {
			event = "UserDeleted"
		}
	} else if !data.IsRollback {
		event = "UserDeletedError"
	}
	_ = s.eventEmitter.Emit(event, map[string]interface{}{"Email": data.Email}, data.TransactionID)
	// TODO: Add transactional outbox pattern

	return OutputData{err}
}
