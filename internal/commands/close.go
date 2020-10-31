package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/urfave/cli/v2"
)

func Close(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	h.Remove(c.Int("id"))

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
