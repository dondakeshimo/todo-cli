package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/domain/reminder"
	"github.com/dondakeshimo/todo-cli/pkg/domain/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:     "modify",
	Short:   "Modify a task",
	Aliases: []string{"m"},
	RunE:    modifyHandler,
}

func init() {
	rootCmd.AddCommand(modifyCmd)

	modifyCmd.Flags().IntP("id", "i", -1, "task's ID")
	modifyCmd.Flags().StringP("task", "t", "", "task contents")
	modifyCmd.Flags().StringP("remind_time", "d", "", "remind_time (2021/3/3 03:03, 2021/3/3, +2h3m, task-4h15m)")
	modifyCmd.Flags().StringP("reminder", "r", "", "choose reminder from [macos]")
	modifyCmd.Flags().Bool("remove_reminder", false, "remove reminder. this option overrides reminder option")
	modifyCmd.Flags().IntP("priority", "p", 100, "task's priority. Lower number means high priority.")

	if err := modifyCmd.MarkFlagRequired("id"); err != nil {
		// NOTE: err is set when "id" is not found in flags.
		//       so, this block never work.
		fmt.Println(err)
	}
}

// modifyHandler invoke usecases.Modify with parameter from cli.
func modifyHandler(c *cobra.Command, args []string) error {
	var r usecases.ModifyRequest

	id, err := c.Flags().GetInt("id") // required
	if err != nil {
		return err
	}
	r.ID = id

	r.Task, err = c.Flags().GetString("task")
	if err != nil {
		return err
	}
	r.IsTask = r.Task != ""

	crt, err := c.Flags().GetString("remind_time")
	if err != nil {
		return err
	}
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
	crr, err := c.Flags().GetBool("remove_reminder")
	if err != nil {
		return err
	}
	if crr {
		r.IsRemoveReminder = true
	}

	r.IsReminder = false
	cr, err := c.Flags().GetString("reminder")
	if err != nil {
		return err
	}
	if !r.IsRemoveReminder && cr != "" {
		rm, err := reminder.NewReminder(cr)
		if err != nil {
			return err
		}
		r.Reminder = rm
		r.IsReminder = true
	}

	r.IsPriority = c.Flags().Changed("priority")
	if r.IsPriority {
		r.Priority, err = c.Flags().GetInt("priority")
		if err != nil {
			return err
		}
	}

	return usecases.Modify(r)
}
