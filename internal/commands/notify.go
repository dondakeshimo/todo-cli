package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/notifier"
	"github.com/urfave/cli/v2"
)

func Notify(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	uuid := c.String("uuid")
	t := h.FindTaskWithUUID(uuid)
	if t == nil {
		return fmt.Errorf("not found uuid: %s", uuid)
	}

	r := notifier.Request{
		Title:    "todo",
		Contents: t.Task,
	}

	var n notifier.Notifier
	if t.Reminder == "macos" {
		n = &notifier.OsascriptNotifier{}
	}

	reply, err := n.Push(&r)
	if err != nil {
		return err
	}

	fmt.Println(reply == "done")

	if reply == "done" {
		h.RemoveTask(t.ID)
	}

	if err := h.Write(); err != nil {
		return err
	}

	return nil
}
