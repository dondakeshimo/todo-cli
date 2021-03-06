package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/notifier"
	"github.com/urfave/cli/v2"
)

// Notify is a function that notify a task.
// It should be called internal only.
func Notify(c *cli.Context) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	uuid := c.String("uuid")
	t, err := h.FindTaskWithUUID(uuid)
	if err != nil {
		return err
	}

	r := notifier.Request{
		Title:    "todo",
		Contents: t.Task(),
		Answer:   []string{"skip", "done"},
	}

	var n notifier.Notifier
	if t.Reminder() == "macos" {
		n = &notifier.OsascriptNotifier{}
	}

	reply, err := n.Push(r)
	if err != nil {
		return err
	}

	if reply == "done" {
		if err := h.RemoveTask(t.ID()); err != nil {
			return err
		}
	}

	if err := h.Commit(); err != nil {
		return err
	}

	return nil
}
