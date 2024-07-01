package services

import (
	"go_service/internal/infrastructure/database/models"
	"log"
)

type SendNotificationSubscriberGateway interface {
	GetAll() []models.Subscriber
}

type SendNotificationEmailGateway interface {
	Send(target string, rate float32) error
}

type SendNotificationCurrencyGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type SendNotificationInputDTO struct {
	From string
	To   string
}

type SendNotificationOutputDTO struct {
	Err error
}

type SendNotification struct {
	emailGateway      SendNotificationEmailGateway
	subscriberGateway SendNotificationSubscriberGateway
	currencyGateway   SendNotificationCurrencyGateway
}

func NewSendNotification(
	emailGateway SendNotificationEmailGateway,
	subscriberGateway SendNotificationSubscriberGateway,
	currencyGateway SendNotificationCurrencyGateway,
) *SendNotification {
	return &SendNotification{
		emailGateway:      emailGateway,
		subscriberGateway: subscriberGateway,
		currencyGateway:   currencyGateway,
	}
}

func (s *SendNotification) Handle(data SendNotificationInputDTO) SendNotificationOutputDTO {
	subscribers := s.subscriberGateway.GetAll()
	rate, err := s.currencyGateway.GetCurrencyRate(data.From, data.To)
	if err != nil {
		return SendNotificationOutputDTO{err}
	}
	for _, subscriber := range subscribers {
		err = s.emailGateway.Send(subscriber.Email, rate)
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
		}
	}
	return SendNotificationOutputDTO{nil}
}
