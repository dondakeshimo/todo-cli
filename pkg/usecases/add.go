package usecases

import (
	"github.com/dondakeshimo/todo-cli/pkg/domain/scheduler"
	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
	"github.com/google/uuid"
)

// AddRequest is a request parameter to invoke Add.
type AddRequest struct {
	Task           string
	Group          string
	RemindTime     task.RemindTime
	IsRemindTime   bool
	RelativeTime   task.RelativeTime
	IsRelativeTime bool
	Reminder       task.Reminder
	IsReminder     bool
	Priority       int
}

// Add is a function that add a task (and reminder).
func Add(r AddRequest) error {
	h, err := task.NewHandler(taskRepository)
	if err != nil {
		return err
	}

	rt := task.RemindTime("")
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

	rm := task.Reminder("")
	if r.IsReminder {
		rm = r.Reminder
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	nt := task.NewTask(0, r.Task, r.Group, rt, uu.String(), rm, r.Priority)

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
