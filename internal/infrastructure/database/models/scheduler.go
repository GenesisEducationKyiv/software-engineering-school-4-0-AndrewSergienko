package models

type ScheduleTime struct {
	Time int64
}

func (ScheduleTime) TableName() string {
	return "scheduler_time"
}
