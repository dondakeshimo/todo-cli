package scheduler

import (
	"time"
)

type Scheduler interface {
	Register(*Request) error
	ClearExpired()
	RemoveWithID(string)
}

type Request struct {
	ID       string
	DateTime time.Time
	Command  string
}
