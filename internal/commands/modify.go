package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/urfave/cli/v2"
)

// Modify is a function that modify a task.
func Modify(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	id := c.Int("id")
	t := h.GetTask(id)
	if t == nil {
		return fmt.Errorf("invalid id: %d", id)
	}

	if st := c.String("task"); st != "" {
		t.Task = st
	}

	rt, err := arrangeRemindTime(c.String("remind_time"), t.RemindTime)
	if err != nil {
		return err
	}
	t.RemindTime = rt

	if c.Bool("remove-reminder") && t.Reminder != "" {
		s, err := scheduler.NewScheduler()
		if err != nil {
			return err
		}

		if err := s.RemoveWithID(t.UUID); err != nil {
			fmt.Println("reminder had been removed for some reason.")
		}

		t.Reminder = ""
		s.ClearExpired()
	}

	if !c.Bool("remove-reminder") && c.String("reminder") != "" {
		rm := c.String("reminder")

		if !task.IsValidReminder(rm) {
			return fmt.Errorf("invalid reminder: %s", rm)
		}

		s, err := scheduler.NewScheduler()
		if err != nil {
			return err
		}

		if t.Reminder != "" {
			if err := s.RemoveWithID(t.UUID); err != nil {
				fmt.Println("previous reminder had been removed for some reason.")
			}
		}

		t.Reminder = rm

		if err := t.SetReminder(s); err != nil {
			return err
		}

		s.ClearExpired()
	}

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
