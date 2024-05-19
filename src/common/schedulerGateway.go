package common

import "time"

type SchedulerReader interface {
	GetLastTime() *time.Time
}

type SchedulerWriter interface {
	SetLastTime() error
}
