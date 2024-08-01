package app

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go_service/internal/notifier/infrastructure/metrics"
	"go_service/internal/notifier/services/sendnotification"
	"log/slog"
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

func (rn RateNotifier) Run() *cron.Cron {
	interactor := rn.container.SendNotification()

	handler := func() {
		slog.Info("Start sending notifications")
		err := interactor.Handle(sendnotification.InputData{From: "USD", To: "UAH"}).Err
		if err != nil {
			metrics.EmailsSentLastTime.WithLabelValues("error").Set(float64(time.Now().Unix()))
			slog.Warn(fmt.Sprintf("Failed to send email. Error: %s", err))
		} else {
			slog.Info("Email sent successfully")
			metrics.EmailsSentLastTime.WithLabelValues("success").Set(float64(time.Now().Unix()))
			err := rn.schedulerGateway.SetLastTime(time.Now())
			if err != nil {
				slog.Warn(fmt.Sprintf("Failed to set last time. Error: %s", err))
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
		return nil
	}
	rn.cronObj.Start()
	slog.Info("Cron started")

	return rn.cronObj
}
