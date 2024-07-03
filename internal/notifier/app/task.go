package app

import (
	"go_service/internal/infrastructure/database/models"
	"go_service/internal/notifier/services"
	"log"
	"time"
)

type SchedulerTimeGateway interface {
	GetLastTime() *time.Time
	SetLastTime() error
}

type SubscriberGateway interface {
	GetAll() []models.Subscriber
}

type EmailGateway interface {
	Send(target string, rate float32) error
}

type CurrencyGateway interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

type RateMailer struct {
	container            InteractorFactory
	schedulerTimeGateway SchedulerTimeGateway
}

func NewRateMailer(
	container InteractorFactory,
	sg SchedulerTimeGateway,
) RateMailer {
	return RateMailer{container: container, schedulerTimeGateway: sg}
}

func (ms RateMailer) Run() {
	lastTime := ms.schedulerTimeGateway.GetLastTime()
	now := time.Now()

	if (lastTime != nil && lastTime.Day() < now.Day() && lastTime.Hour() >= now.Hour()) || lastTime == nil {
		ms.RunSending()
		lastTime = &now
	}

	for {
		time.Sleep(time.Until(ms.GetNextTime(lastTime)))
		ms.RunSending()
		lastTime = &now
	}
}

func (ms RateMailer) RunSending() {
	err := ms.SendRateToAll()
	if err != nil {
		log.Printf("Failed to send rate mail to all emails: %v\n", err)
	}
	err = ms.schedulerTimeGateway.SetLastTime()
	if err != nil {
		log.Printf("Failed to save last sending time: %v\n", err)
	}
}

func (ms RateMailer) SendRateToAll() error {
	interactor := ms.container.SendNotification()
	return interactor.Handle(services.SendNotificationInputDTO{From: "USD", To: "UAH"}).Err
}

func (RateMailer) GetNextTime(lt *time.Time) time.Time {
	return lt.Add(24 * time.Hour)
}
