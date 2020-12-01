package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/urfave/cli/v2"
)

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

		s.RemoveWithID(t.UUID)
	}

	h.RemoveTasks(ids)

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
