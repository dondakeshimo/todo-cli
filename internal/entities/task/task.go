package task

import (
	"fmt"
	"os"

	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

type Task struct {
	ID         int    `json:"id"`
	Task       string `json:"task"`
	RemindTime string `json:"remindtime"`
	UUID       string `json:"uuid"`
	Reminder   string `json:"reminder"`
}

func (t *Task) SetReminder(s scheduler.Scheduler) error {
	ts, err := timestr.Parse(t.RemindTime)
	if err != nil {
		return err
	}

	exe, err := os.Executable()
	if err != nil {
		return err
	}

	sr := &scheduler.Request{
		ID:       t.UUID,
		DateTime: *ts,
		Command:  fmt.Sprintf("%s notify --uuid %s", exe, t.UUID),
	}

	if err := s.Register(sr); err != nil {
		return err
	}

	return err
}

func IsValidReminder(r string) bool {
	// TODO: add slack
	allowReminders := []string{"macos"}
	for _, a := range allowReminders {
		if r == a {
			return true
		}
	}

	return false
}
