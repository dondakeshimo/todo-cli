package commands

import (
	"errors"
	"time"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
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

	t := h.TaskList.Tasks[id-1]

	// validation
	if task := c.String("task"); task != "" {
		t.Task = task
	}

	if d := c.String("deadline"); d != "" {
		layout := "2006/01/02 15:04"
		_, err := time.Parse(layout, c.String("deadline"))
		if err != nil {
			return err
		}
		t.Deadline = d
	}

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
