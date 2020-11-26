package task

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
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

	sr := &scheduler.Request{
		DateTime: *ts,
		Command: fmt.Sprintf("todo notify --uuid %s", t.UUID),
	}

	if err := s.Register(sr); err != nil {
		return err
	}

	s.ClearExpired()

	return err
}

func IsValidReminder(r string) bool {
	allowReminders := []string{"macos", "slack"}
	for _, a := range allowReminders {
		if r == a {
			return true
		}
	}

	return false
}
