package models

type SchedulerTime struct {
	Time int64
}

func (SchedulerTime) TableName() string {
	return "scheduler_time"
}
