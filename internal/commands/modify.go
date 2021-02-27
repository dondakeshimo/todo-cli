package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/dondakeshimo/todo-cli/internal/values/reminder"
	"github.com/dondakeshimo/todo-cli/internal/values/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/urfave/cli/v2"
)

type modifyParams struct {
	id               int
	task             string
	isTask           bool
	remindTime       remindtime.RemindTime
	isRemindTime     bool
	relativeTime     remindtime.RelativeTime
	isRelativeTime   bool
	isRemoveReminder bool
	reminder         reminder.Reminder
	isReminder       bool
}

func newModifyParams(c *cli.Context) (*modifyParams, error) {
	p := new(modifyParams)

	p.id = c.Int("id") // required

	p.task = c.String("task")
	p.isTask = p.task != ""

	crt := c.String("remind_time")
	if crt == "" {
		p.isRemindTime = false
		p.isRelativeTime = false
	}

	if crt != "" && remindtime.IsValidRelativeTime(crt) {
		td, err := remindtime.NewRelativeTime(crt)
		if err != nil {
			return nil, err
		}
		p.relativeTime = td
		p.isRelativeTime = true
		p.isRemindTime = false
	}

	if crt != "" && !remindtime.IsValidRelativeTime(crt) {
		rt, err := remindtime.NewRemindTime(crt)
		if err != nil {
			return nil, err
		}
		p.remindTime = rt
		p.isRemindTime = true
		p.isRelativeTime = false
	}

	// NOTE: assert isRelativeTime and isRemindTime never be true
	if p.isRelativeTime && p.isRemindTime {
		return nil, fmt.Errorf("internal command error")
	}

	p.isRemoveReminder = false
	if c.Bool("remove_reminder") {
		p.isRemoveReminder = true
	}

	p.isReminder = false
	if !p.isRemoveReminder && c.String("reminder") != "" {
		rm, err := reminder.NewReminder(c.String("reminder"))
		if err != nil {
			return nil, err
		}
		p.reminder = rm
		p.isReminder = true
	}

	return p, nil
}

// Modify is a function that modify a task.
func Modify(c *cli.Context) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	p, err := newModifyParams(c)
	if err != nil {
		return err
	}

	t, err := h.GetTask(p.id)
	if err != nil {
		return err
	}

	newTask := t.Task()
	if p.isTask {
		newTask = p.task
	}

	newRemindTime := t.RemindTime()
	if p.isRelativeTime {
		nrt, err := t.RemindTime().AddTime(p.relativeTime)
		if err != nil {
			return err
		}
		newRemindTime = nrt
	}

	if p.isRemindTime {
		newRemindTime = p.remindTime
	}

	newReminder := t.Reminder()

	// construct scheduler when remove previous reminder or update reminder
	var s scheduler.Scheduler
	if (p.isRemoveReminder && t.Reminder() != "") || (!p.isRemoveReminder && p.isReminder) {
		var err error
		s, err = scheduler.NewScheduler()
		if err != nil {
			return err
		}
	}

	if p.isRemoveReminder && t.Reminder() != "" {
		if err := t.RemoveReminder(s); err != nil {
			fmt.Println("reminder had been removed for some reason.")
		}
		newReminder = ""
	}

	if !p.isRemoveReminder && p.isReminder {
		if t.Reminder() != "" {
			if err := t.RemoveReminder(s); err != nil {
				fmt.Println("previous reminder had been removed for some reason.")
			}
		}

		newReminder = p.reminder
	}

	nt := task.NewTask(t.ID(), newTask, newRemindTime, t.UUID(), newReminder)

	if !p.isRemoveReminder && p.isReminder {
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
