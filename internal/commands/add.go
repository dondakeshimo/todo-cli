package commands

import (
	"errors"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	// validation
	if c.String("task") == "" {
		return errors.New("task could not be empty")
	}

	d, err := timestr.Parse(c.String("deadline"))
	if err != nil {
		return err
	}

	h.AppendTask(&task.Task{
		Task:     c.String("task"),
		Deadline: d,
	})

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
