package scheduler

import (
	"fmt"
	"runtime"
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

func NewScheduler() (Scheduler, error) {
	if runtime.GOOS == "darwin" {
		return NewLaunchdScheduler(), nil
	}

	return nil, fmt.Errorf("no appropriate scheduler with [%s]", runtime.GOOS)
}
