package scheduler

import (
	"fmt"
	"runtime"
	"time"
)

// Scheduler is a interface that schedule command.
//go:generate mockgen -destination=./mock_scheduler.go -package=scheduler -self_package=github.com/dondakeshimo/todo-cli/pkg/domain/scheduler . Scheduler
type Scheduler interface {
	Register(Request) error
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

	if runtime.GOOS == "linux" {
		return NewCronScheduler(), nil
	}

	return nil, fmt.Errorf("no appropriate scheduler with [%s]", runtime.GOOS)
}

// isExpired is a function that judge plist file is enough old.
func isExpired(t time.Time) bool {
	deadline := time.Now().Add(-time.Duration(1) * time.Minute).Unix()
	return t.Unix() < deadline
}
