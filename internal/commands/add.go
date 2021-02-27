package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/dondakeshimo/todo-cli/internal/values/reminder"
	"github.com/dondakeshimo/todo-cli/internal/values/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

type addParams struct {
	task           string
	remindTime     remindtime.RemindTime
	isRemindTime   bool
	relativeTime   remindtime.RelativeTime
	isRelativeTime bool
	reminder       reminder.Reminder
	isReminder     bool
}

func newAddParams(c *cli.Context) (*addParams, error) {
	p := new(addParams)

	p.task = c.Args().Get(0)
	if p.task == "" {
		return nil, fmt.Errorf("`$ todo add` need an argument what represents a task")
	}

	crt := c.String("remind_time")
	if crt == "" {
		p.isRemindTime = false
		p.isRelativeTime = false
	}

	if crt != "" && remindtime.IsRelativeToNow(crt) {
		td, err := remindtime.NewRelativeTime(crt)
		if err != nil {
			return nil, err
		}
		p.relativeTime = td
		p.isRelativeTime = true
		p.isRemindTime = false
	}

	if crt != "" && !remindtime.IsRelativeToNow(crt) {
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

	p.isReminder = false
	if c.String("reminder") != "" {
		rm, err := reminder.NewReminder(c.String("reminder"))
		if err != nil {
			return nil, err
		}
		p.reminder = rm
		p.isReminder = true
	}

	return p, nil
}

// Add is a function that add a task (and reminder).
func Add(c *cli.Context) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	p, err := newAddParams(c)
	if err != nil {
		return err
	}

	rt := remindtime.RemindTime("")
	if p.isRelativeTime {
		nrt, err := rt.AddTime(p.relativeTime)
		if err != nil {
			return err
		}
		rt = nrt
	}

	if p.isRemindTime {
		rt = p.remindTime
	}

	rm := reminder.Reminder("")
	if p.isReminder {
		rm = p.reminder
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	nt := task.NewTask(0, p.task, rt, uu.String(), rm)

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
