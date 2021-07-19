package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/domain/reminder"
	"github.com/dondakeshimo/todo-cli/pkg/domain/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a task",
	Aliases: []string{"a"},
	RunE:    addHandler,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("remind_time", "d", "", "remind_time (2021/3/3 03:03, 2021/3/3, +2h3m, task-4h15m)")
	addCmd.Flags().StringP("reminder", "r", "", "choose reminder from [macos, slack]")
	addCmd.Flags().IntP("priority", "p", 100, "task's priority. Lower number means high priority.")
}

// addHandler invoke usecases.Add with parameter from cli.
func addHandler(c *cobra.Command, args []string) error {
	var r usecases.AddRequest

	if len(args) != 1 {
		return fmt.Errorf("`$ todo add` need an argument what represents a task")
	}

	r.Task = args[0]
	if r.Task == "" {
		return fmt.Errorf("`$ todo add` need an argument what represents a task")
	}

	crt, err := c.Flags().GetString("remind_time")
	if err != nil {
		return err
	}

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
	cr, err := c.Flags().GetString("reminder")
	if err != nil {
		return err
	}

	if cr != "" {
		rm, err := reminder.NewReminder(cr)
		if err != nil {
			return err
		}
		r.Reminder = rm
		r.IsReminder = true
	}

	r.Priority, err = c.Flags().GetInt("priority")
	if err != nil {
		return err
	}

	return usecases.Add(r)
}
