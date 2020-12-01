package scheduler

import (
	"fmt"
	"runtime"
	"time"
)

// Scheduler is a interface that schedule command.
type Scheduler interface {
	Register(*Request) error
	ClearExpired()
	RemoveWithID(string) error
}

// Request is a struct that is passed to Scheduler.
type Request struct {
	ID       string
	DateTime time.Time
	Command  string
}

// NewScheduler is a constructor that is make new appropriate Scheduler.
func NewScheduler() (Scheduler, error) {
	if runtime.GOOS == "darwin" {
		return NewLaunchdScheduler(), nil
	}

	return nil, fmt.Errorf("no appropriate scheduler with [%s]", runtime.GOOS)
}
