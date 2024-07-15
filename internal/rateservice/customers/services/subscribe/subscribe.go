package subscribe

import (
	"go_service/internal/rateservice/customers/infrastructure/database/models"
	"go_service/internal/rateservice/customers/services"
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

type EventEmitter interface {
	Emit(event string, data map[string]interface{}) error
}

type Subscribe struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway) *Subscribe {
	return &Subscribe{subscriberGateway: sg}
}

func (s *Subscribe) Handle(data InputDTO) OutputDTO {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return OutputDTO{Err: &services.EmailConflictError{Email: data.Email}}
	}
	err := s.subscriberGateway.Create(data.Email)
	if err == nil {
		err = s.eventEmitter.Emit("UserCreated", map[string]interface{}{"email": data.Email})
	}
	return OutputDTO{err}
}
