package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/urfave/cli/v2"
)

// Close is a function that close a task or tasks.
func Close(c *cli.Context) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	ids := c.IntSlice("ids")

	for _, id := range ids {
		t, err := h.GetTask(id)
		if err != nil {
			return err
		}

		if t.Reminder() == "" {
			continue
		}

		// NOTE: ignore err message
		s, err := scheduler.NewScheduler()
		if err != nil {
			continue
		}

		// NOTE: ignore err message
		if err := t.RemoveReminder(s); err != nil {
			continue
		}
	}

	if err := h.RemoveTasks(ids); err != nil {
		return err
	}

	if err := h.Commit(); err != nil {
		return err
	}

	return nil
}
