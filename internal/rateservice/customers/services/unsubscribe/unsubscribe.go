package unsubscribe

import "go_service/internal/rateservice/customers/infrastructure/database/models"

type InputDTO struct {
	Email string
}

type OutputDTO struct {
	Err error
}

type SubscriberGateway interface {
	GetByEmail(email string) *models.Subscriber
	DeleteByEmail(email string) error
}

type EventEmitter interface {
	Emit(name string, data map[string]interface{}) error
}

type Unsubscribe struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway, em EventEmitter) *Unsubscribe {
	return &Unsubscribe{subscriberGateway: sg, eventEmitter: em}
}

func (s *Unsubscribe) Handle(data InputDTO) OutputDTO {
	err := s.subscriberGateway.DeleteByEmail(data.Email)
	if err == nil {
		err = s.eventEmitter.Emit("UserDeleted", map[string]interface{}{"email": data.Email})
	}
	return OutputDTO{err}
}
