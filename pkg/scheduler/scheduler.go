package scheduler

import (
	"time"
)

type Scheduler interface {
	Register(*Request) error
}

type Request struct {
	DateTime time.Time
	Command  string
}
