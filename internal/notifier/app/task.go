package app

import (
	"github.com/robfig/cron/v3"
	"go_service/internal/notifier/services/sendnotification"
	"log"
	"time"
)

type SchedulerGateway interface {
	GetLastTime() *time.Time
	SetLastTime(lastTime time.Time) error
}

type RateNotifier struct {
	container        InteractorFactory
	schedulerGateway SchedulerGateway
	cronObj          *cron.Cron
}

func NewRateNotifier(
	container InteractorFactory,
	schedulerGateway SchedulerGateway,
) RateNotifier {
	return RateNotifier{container: container, schedulerGateway: schedulerGateway, cronObj: cron.New()}
}

func (rn RateNotifier) Run() {
	interactor := rn.container.SendNotification()

	handler := func() {
		if interactor.Handle(sendnotification.InputData{From: "USD", To: "UAH"}).Err != nil {
			log.Printf("EMAIL SEND ERROR")
		} else {
			log.Printf("EMAIL SEND SUCCESS")
			err := rn.schedulerGateway.SetLastTime(time.Now())
			if err != nil {
				return
			}
		}
	}

	// Check if notification was sent today and send it if not
	now := time.Now()
	lastTime := rn.schedulerGateway.GetLastTime()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	sendTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location())

	if lastTime == nil || lastTime.Before(midnight) && now.After(sendTime) {
		handler()
	}

	// Setup cron notifier job
	_, err := rn.cronObj.AddFunc("0 9 * * *", handler)
	if err != nil {
		return
	}
	rn.cronObj.Start()
	log.Printf("CRON STARTED")
}
