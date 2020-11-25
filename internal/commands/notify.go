package commands

import (
	"fmt"
	"errors"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/notifier"
	"github.com/urfave/cli/v2"
)

func Notify(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	id := c.Int("id")
	t := h.GetTask(id)
	if t == nil {
		return errors.New("invalid id")
	}

	r := notifier.Request{
		Title: "todo",
		Contents: t.Task,
	}
	n := notifier.OsascriptNotifier{}
	fmt.Printf("r: %v", r)


	if err := n.Push(&r); err != nil {
		return err
	}

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
