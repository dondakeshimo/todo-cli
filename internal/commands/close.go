package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/urfave/cli/v2"
)

func Close(c *cli.Context) error {
	th, err := task.NewTaskHandler()
	if err != nil {
		return err
	}

	th.Remove(c.Int("id"))

	if err := th.Write(); err != nil {
		return err
	}

	return nil
}
