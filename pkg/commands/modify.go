package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/domain/reminder"
	"github.com/dondakeshimo/todo-cli/pkg/domain/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/urfave/cli/v2"
)

// Modify invoke usecases.Modify with parameter from cli.
func Modify(c *cli.Context) error {
	var r usecases.ModifyRequest

	r.ID = c.Int("id") // required

	r.Task = c.String("task")
	r.IsTask = r.Task != ""

	crt := c.String("remind_time")
	if crt == "" {
		r.IsRemindTime = false
		r.IsRelativeTime = false
	}

	if crt != "" && remindtime.IsValidRelativeTime(crt) {
		td, err := remindtime.NewRelativeTime(crt)
		if err != nil {
			return err
		}
		r.RelativeTime = td
		r.IsRelativeTime = true
		r.IsRemindTime = false
	}

	if crt != "" && !remindtime.IsValidRelativeTime(crt) {
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

	r.IsRemoveReminder = false
	if c.Bool("remove_reminder") {
		r.IsRemoveReminder = true
	}

	r.IsReminder = false
	if !r.IsRemoveReminder && c.String("reminder") != "" {
		rm, err := reminder.NewReminder(c.String("reminder"))
		if err != nil {
			return err
		}
		r.Reminder = rm
		r.IsReminder = true
	}

	r.IsPriority = c.IsSet("priority")
	if r.IsPriority {
		r.Priority = c.Int("priority")
	}

	return usecases.Modify(r)
}
