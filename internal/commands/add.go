package commands

import (
	"errors"
	"time"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
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

	if c.String("deadline") != "" {
		layout := "2006/01/02 15:04"
		_, err := time.Parse(layout, c.String("deadline"))
		if err != nil {
			return err
		}
	}

	tl := h.TaskList
	tl.Tasks = append(tl.Tasks, &task.Task{
		ID:       len(tl.Tasks) + 1,
		Task:     c.String("task"),
		Deadline: c.String("deadline"),
	})

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
