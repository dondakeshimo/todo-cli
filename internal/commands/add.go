package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	d, err := timestr.Validate(c.String("remind_time"))
	if err != nil {
		return err
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	h.AppendTask(&task.Task{
		Task:     c.String("task"),
		RemindTime: d,
		UUID: uu.String(),
	})

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
