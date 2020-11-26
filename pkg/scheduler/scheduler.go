package scheduler

import (
	"time"
)

type Scheduler interface {
	Register(*Request) error
	ClearExpired()
}

type Request struct {
	DateTime time.Time
	Command  string
}
