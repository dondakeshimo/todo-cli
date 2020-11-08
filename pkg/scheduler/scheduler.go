package scheduler

import (
	"time"
)

type Scheduler interface {
	Register() error
}

type Request struct {
	DateTime time.Time
	Command string
}
