package src

import (
	"go_service/src/common"
	"time"
)

type SchedulerTimeGateway interface {
	common.SchedulerReader
	common.SchedulerWriter
}

type RateMailer struct {
	Es common.EmailSender
	Sr common.SubscriberListReader
	Sg SchedulerTimeGateway
	Cr common.CurrencyReader
}

func (ms RateMailer) Run() {
	lastTime := ms.Sg.GetLastTime()
	now := time.Now()

	if (lastTime != nil && lastTime.Day() < now.Day() && lastTime.Hour() >= now.Hour()) || lastTime == nil {
		err := ms.SendRateToAll()
		if err != nil {
			return
		}
		err = ms.Sg.SetLastTime()
		if err != nil {
			return
		}
		lastTime = &now
	}

	for {
		time.Sleep(time.Until(ms.GetNextTime(lastTime)))
		err := ms.SendRateToAll()
		if err != nil {
			return
		}
		err = ms.Sg.SetLastTime()
		if err != nil {
			return
		}
		lastTime = &now
	}

}

func (ms RateMailer) SendRateToAll() error {
	subscribers := ms.Sr.GetAll()
	rate, err := ms.Cr.GetUSDCurrencyRate()
	if err != nil {
		return err
	}
	for _, subscriber := range subscribers {
		go ms.Es.Send(subscriber.Email, rate)
	}
	return nil
}

func (RateMailer) GetNextTime(lt *time.Time) time.Time {
	return lt.Add(24 * time.Hour)
}
