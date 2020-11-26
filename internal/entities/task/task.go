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

func (t *Task) SetReminder() error {
	ts, err := timestr.Parse(t.RemindTime)
	if err != nil {
		return err
	}

	sr := &scheduler.Request{
		DateTime: *ts,
		Command: fmt.Sprintf("todo notify --uuid %s", t.UUID),
	}

	// TODO: abstraction
	s := scheduler.NewLaunchdScheduler()

	if err := s.Register(sr); err != nil {
		return err
	}

	return err
}
