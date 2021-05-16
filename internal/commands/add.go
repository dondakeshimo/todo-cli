package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/internal/entities/reminder"
	"github.com/dondakeshimo/todo-cli/internal/entities/remindtime"
	"github.com/dondakeshimo/todo-cli/internal/usecases"
	"github.com/urfave/cli/v2"
)

// Add invoke usecases.Add with parameter from cli.
func Add(c *cli.Context) error {
    var r usecases.AddRequest

	r.Task = c.Args().Get(0)
	if r.Task == "" {
		return fmt.Errorf("`$ todo add` need an argument what represents a task")
	}

	crt := c.String("remind_time")
	if crt == "" {
		r.IsRemindTime = false
		r.IsRelativeTime = false
	}

	if crt != "" && remindtime.IsRelativeToNow(crt) {
		td, err := remindtime.NewRelativeTime(crt)
		if err != nil {
			return err
		}
		r.RelativeTime = td
		r.IsRelativeTime = true
		r.IsRemindTime = false
	}

	if crt != "" && !remindtime.IsRelativeToNow(crt) {
		rt, err := remindtime.NewRemindTime(crt)
		if err != nil {
			return err
		}
		r.RemindTime = rt
		r.IsRemindTime = true
		r.IsRelativeTime = false
	}

	// NOTE: assert isRelativeTime and isRemindTime never be true
	if r.IsRelativeTime && r.IsRemindTime {
		return fmt.Errorf("internal command error")
	}

	r.IsReminder = false
	if c.String("reminder") != "" {
		rm, err := reminder.NewReminder(c.String("reminder"))
		if err != nil {
			return err
		}
		r.Reminder = rm
		r.IsReminder = true
	}

	r.Priority = c.Int("priority")

    return usecases.Add(r)
}
