package subscribe

import (
	"go_service/internal/rateservice/customers/infrastructure/database/models"
	"go_service/internal/rateservice/customers/services"
	"log"
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
	Emit(name string, data map[string]interface{}) error
}

type Subscribe struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway, em EventEmitter) *Subscribe {
	return &Subscribe{subscriberGateway: sg, eventEmitter: em}
}

func (s *Subscribe) Handle(data InputDTO) OutputDTO {
	if s.subscriberGateway.GetByEmail(data.Email) != nil {
		return OutputDTO{Err: &services.EmailConflictError{Email: data.Email}}
	}
	err := s.subscriberGateway.Create(data.Email)
	if err == nil {
		err = s.eventEmitter.Emit("UserCreated", map[string]interface{}{"email": data.Email})
		log.Println("UserCreated event emitted")
	}
	return OutputDTO{err}
}
