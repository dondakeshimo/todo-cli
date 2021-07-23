package usecases

import (
	"github.com/dondakeshimo/todo-cli/pkg/domain/notifier"
	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
)

// Notify is a function that notify a task.
// It should be called internal only.
func Notify(uuid string) error {
	h, err := task.NewHandler(taskRepository)
	if err != nil {
		return err
	}

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
	if t.Reminder() == task.MacOS {
		n, err = notifier.NewOsascriptNotifier()
		if err != nil {
			return err
		}
	}
	if t.Reminder() == task.Slack {
		n, err = notifier.NewSlackNotifier(config.SlackWebhookURL, config.SlackMentionTo)
		if err != nil {
			return err
		}
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
