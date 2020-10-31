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
	// validation
	if id == 0 {
		return errors.New("id could not be empty")
	}

	t := h.GetTask(id)
	if t == nil {
		return errors.New("invalid id")
	}

	// validation
	if task := c.String("task"); task != "" {
		t.Task = task
	}


	d, err := timestr.Parse(c.String("deadline"))
	if err != nil {
		return err
	}
	t.Deadline = d

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
