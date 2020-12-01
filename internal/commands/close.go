package commands

import (
	"runtime"

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

	for _, i := range ids {
		t := h.GetTask(i)
		if t.Reminder == "" {
			continue
		}

		var s scheduler.Scheduler
		if runtime.GOOS == "darwin" {
			s = scheduler.NewLaunchdScheduler()
		} else {
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
