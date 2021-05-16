package usecases

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/reminder"
	"github.com/dondakeshimo/todo-cli/internal/entities/remindtime"
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/google/uuid"
)

type AddRequest struct {
	Task           string
	RemindTime     remindtime.RemindTime
	IsRemindTime   bool
	RelativeTime   remindtime.RelativeTime
	IsRelativeTime bool
	Reminder       reminder.Reminder
	IsReminder     bool
	Priority       int
}

// Add is a function that add a task (and reminder).
func Add(r AddRequest) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	rt := remindtime.RemindTime("")
	if r.IsRelativeTime {
		nrt, err := rt.AddTime(r.RelativeTime)
		if err != nil {
			return err
		}
		rt = nrt
	}

	if r.IsRemindTime {
		rt = r.RemindTime
	}

	rm := reminder.Reminder("")
	if r.IsReminder {
		rm = r.Reminder
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	nt := task.NewTask(0, r.Task, rt, uu.String(), rm, r.Priority)

	h.AppendTask(nt)

	if err := h.Commit(); err != nil {
		return err
	}

	// when do not remind, do early return
	if nt.Reminder() == "" {
		return nil
	}

	s, err := scheduler.NewScheduler()
	if err != nil {
		return err
	}

	if err := nt.SetReminder(s); err != nil {
		return err
	}

	return nil
}
