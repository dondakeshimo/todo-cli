package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/urfave/cli/v2"
)

// Close is a function that close a task or tasks.
func Close(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	ids := c.IntSlice("ids")

	for _, id := range ids {
		t := h.GetTask(id)
		if t.Reminder == "" {
			continue
		}

		// NOTE: ignore err message
		s, err := scheduler.NewScheduler()
		if err != nil {
			continue
		}

		// NOTE: ignore err message
		if err := s.RemoveWithID(t.UUID); err != nil {
			continue
		}
	}

	h.RemoveTasks(ids)

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
