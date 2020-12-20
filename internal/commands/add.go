package commands

import (
	"fmt"
	"time"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

// Add is a function that add a task (and reminder).
func Add(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	rt := c.String("remind_time")

	if c.Bool("relative") {
		rt, err = timestr.TransformFromRelative(rt, time.Now())
		if err != nil {
			return err
		}
	}

	d, err := timestr.UnifyLayout(rt)
	if err != nil {
		return err
	}

	r := c.String("reminder")
	if r != "" && !task.IsValidReminder(r) {
		return fmt.Errorf("invalid reminder: %s", r)
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	t := &task.Task{
		Task:       c.String("task"),
		RemindTime: d,
		UUID:       uu.String(),
		Reminder:   r,
	}

	h.AppendTask(t)

	if err := h.Write(); err != nil {
		return err
	}

	// when do not remind, do early return
	if r == "" {
		return nil
	}

	s, err := scheduler.NewScheduler()
	if err != nil {
		return err
	}

	if err := t.SetReminder(s); err != nil {
		return err
	}

	s.ClearExpired()

	return nil
}
