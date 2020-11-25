package commands

import (
	"fmt"
	"errors"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/urfave/cli/v2"
)

func Remind(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	id := c.Int("id")
	t := h.GetTask(id)
	if t == nil {
		return errors.New("invalid id")
	}

	ts, err := timestr.Parse(t.RemindTime)
	if err != nil {
		return err
	}

	sr := &scheduler.Request{
		DateTime: *ts,
		Command: fmt.Sprintf("todo notify --id %d", id),
	}

	s := scheduler.NewLaunchdScheduler()

	if err := s.Register(sr); err != nil {
		return err
	}

	return nil
}
