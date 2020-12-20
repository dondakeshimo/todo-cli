package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
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

	d, err := timestr.UnifyLayout(c.String("remind_time"))
	if err != nil {
		return err
	}

	r := c.String("reminder")
	if r != "" && !task.IsValidReminder(r) {
		return fmt.Errorf("invalid reminder: %s", r)
	}

	preReminder := t.Reminder

	t.Task = c.String("task")
	t.RemindTime = d
	t.Reminder = r

	if err := h.Write(); err != nil {
		return err
	}

	// when do not remind, do early return
	if preReminder == "" && r == "" {
		return nil
	}

	s, err := scheduler.NewScheduler()
	if err != nil {
		return err
	}

	if preReminder != "" {
		if err := s.RemoveWithID(t.UUID); err != nil {
			fmt.Println("reminder is removed for some reason.")
		}
	}

	// when do not remind, do early return
	if r == "" {
		return nil
	}

	if err := t.SetReminder(s); err != nil {
		return err
	}

	s.ClearExpired()

	return nil
}
