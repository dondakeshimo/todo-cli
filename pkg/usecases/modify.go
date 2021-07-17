package usecases

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/domain/reminder"
	"github.com/dondakeshimo/todo-cli/pkg/domain/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/domain/scheduler"
	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
	"github.com/dondakeshimo/todo-cli/pkg/gateways/json"
)

// ModifyRequest is a request parameter to invoke Modify.
type ModifyRequest struct {
	ID               int
	Task             string
	IsTask           bool
	RemindTime       remindtime.RemindTime
	IsRemindTime     bool
	RelativeTime     remindtime.RelativeTime
	IsRelativeTime   bool
	IsRemoveReminder bool
	Reminder         reminder.Reminder
	IsReminder       bool
	Priority         int
	IsPriority       bool
}

// Modify is a function that modify a task.
func Modify(r ModifyRequest) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	t, err := h.GetTask(r.ID)
	if err != nil {
		return err
	}

	newTask := t.Task()
	if r.IsTask {
		newTask = r.Task
	}

	newRemindTime := t.RemindTime()
	if r.IsRelativeTime {
		nrt, err := t.RemindTime().AddTime(r.RelativeTime)
		if err != nil {
			return err
		}
		newRemindTime = nrt
	}

	if r.IsRemindTime {
		newRemindTime = r.RemindTime
	}

	newReminder := t.Reminder()

	// construct scheduler when remove previous reminder or update reminder
	var s scheduler.Scheduler
	if (r.IsRemoveReminder && t.Reminder() != "") || (!r.IsRemoveReminder && r.IsReminder) {
		var err error
		s, err = scheduler.NewScheduler()
		if err != nil {
			return err
		}
	}

	if r.IsRemoveReminder && t.Reminder() != "" {
		if err := t.RemoveReminder(s); err != nil {
			fmt.Println("reminder had been removed for some reason.")
		}
		newReminder = ""
	}

	if !r.IsRemoveReminder && r.IsReminder {
		if t.Reminder() != "" {
			if err := t.RemoveReminder(s); err != nil {
				fmt.Println("previous reminder had been removed for some reason.")
			}
		}

		newReminder = r.Reminder
	}

	newPriority := t.Priority()
	if r.IsPriority {
		newPriority = r.Priority
	}

	nt := task.NewTask(t.ID(), newTask, newRemindTime, t.UUID(), newReminder, newPriority)

	if !r.IsRemoveReminder && r.IsReminder {
		if err := nt.SetReminder(s); err != nil {
			return err
		}
	}

	if err := h.UpdateTask(t.ID(), nt); err != nil {
		return err
	}

	if err := h.Commit(); err != nil {
		return err
	}

	return nil
}
