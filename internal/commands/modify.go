package commands

import (
	"errors"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/urfave/cli/v2"
)

func Modify(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	id := c.Int("id")
	t := h.GetTask(id)
	if t == nil {
		return errors.New("invalid id")
	}

	d, err := timestr.Parse(c.String("deadline"))
	if err != nil {
		return err
	}

	t.Task = c.String("task")
	t.Deadline = d

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
