package app

import (
	"go_service/internal/common"
	"log"
	"time"
)

type SchedulerTimeGateway interface {
	common.SchedulerReader
	common.SchedulerWriter
}

type RateMailer struct {
	emailGateway         common.EmailSender
	subscriberGateway    common.SubscriberListReader
	schedulerTimeGateway SchedulerTimeGateway
	currencyGateway      common.CurrencyReader
}

func InitRateMailer(
	es common.EmailSender,
	sr common.SubscriberListReader,
	sg SchedulerTimeGateway,
	cr common.CurrencyReader,
) RateMailer {
	return RateMailer{es, sr, sg, cr}
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
	subscribers := ms.subscriberGateway.GetAll()
	rate, err := ms.currencyGateway.GetUSDCurrencyRate()
	if err != nil {
		return err
	}
	for _, subscriber := range subscribers {
		err = ms.emailGateway.Send(subscriber.Email, rate)
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
		}
	}
	return nil
}

func (RateMailer) GetNextTime(lt *time.Time) time.Time {
	return lt.Add(24 * time.Hour)
}
