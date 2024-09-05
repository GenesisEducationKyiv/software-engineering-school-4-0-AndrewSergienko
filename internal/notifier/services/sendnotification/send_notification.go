package sendnotification

import (
	"fmt"
	"go_service/internal/notifier/infrastructure/database/models"
	"log/slog"
)

type SubscriberGateway interface {
	GetAll() []models.Subscriber
}

type EmailGateway interface {
	Send(target string, rate float32) error
}

type CurrencyRateGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type InputData struct {
	From string
	To   string
}

type OutputData struct {
	Err       error
	ErrEmails []string
}

type SendNotification struct {
	emailGateway      EmailGateway
	subscriberGateway SubscriberGateway
	currencyGateway   CurrencyRateGateway
}

func New(
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

func (s *SendNotification) Handle(data InputData) OutputData {
	var rate float32

	subscribers := s.subscriberGateway.GetAll()
	if subscribers == nil {
		return OutputData{nil, make([]string, 0)}
	}

	rate, err := s.currencyGateway.GetCurrencyRate(data.From, data.To)
	if err != nil {
		return OutputData{err, make([]string, 0)}
	}

	errEmails := make([]string, 0)
	for _, subscriber := range subscribers {
		err = s.emailGateway.Send(subscriber.Email, rate)
		if err != nil {
			errEmails = append(errEmails, subscriber.Email)
			slog.Warn(fmt.Sprintf("Failed to send email: %v\n", err))
		}
	}
	return OutputData{nil, errEmails}
}
