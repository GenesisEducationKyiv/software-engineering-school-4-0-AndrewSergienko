package services

import (
	"log"
)

type Subscriber struct {
	Email string
}

type SubscriberGateway interface {
	GetAll() ([]string, error)
}

type EmailGateway interface {
	Send(target string, rate float32) error
}

type CurrencyRateGateway interface {
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
	emailGateway      EmailGateway
	subscriberGateway SubscriberGateway
	currencyGateway   CurrencyRateGateway
}

func NewSendNotification(
	emailGateway EmailGateway,
	subscriberGateway SubscriberGateway,
	currencyGateway CurrencyRateGateway,
) *SendNotification {
	return &SendNotification{
		emailGateway:      emailGateway,
		subscriberGateway: subscriberGateway,
		currencyGateway:   currencyGateway,
	}
}

func (s *SendNotification) Handle(data SendNotificationInputDTO) SendNotificationOutputDTO {
	var rate float32

	subscribers, err := s.subscriberGateway.GetAll()
	if err != nil {
		return SendNotificationOutputDTO{err}
	}

	rate, err = s.currencyGateway.GetCurrencyRate(data.From, data.To)
	if err != nil {
		return SendNotificationOutputDTO{err}
	}
	for _, subscriber := range subscribers {
		err = s.emailGateway.Send(subscriber, rate)
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
		}
	}
	return SendNotificationOutputDTO{nil}
}
