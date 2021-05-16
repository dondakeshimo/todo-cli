package usecases

import (
	"github.com/dondakeshimo/todo-cli/pkg/entities/scheduler"
	"github.com/dondakeshimo/todo-cli/pkg/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/gateways/json"
)

// Close is a function that close a task or tasks.
func Close(ids []int) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	for _, id := range ids {
		t, err := h.GetTask(id)
		if err != nil {
			return err
		}

		if t.Reminder() == "" {
			continue
		}

		// NOTE: ignore err message
		s, err := scheduler.NewScheduler()
		if err != nil {
			continue
		}

		// NOTE: ignore err message
		if err := t.RemoveReminder(s); err != nil {
			continue
		}
	}

	if err := h.RemoveTasks(ids); err != nil {
		return err
	}

	if err := h.Commit(); err != nil {
		return err
	}

	return nil
}
