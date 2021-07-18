package usecases

import (
	"github.com/dondakeshimo/todo-cli/pkg/domain/scheduler"
	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
)

// Close is a function that close a task or tasks.
func Close(ids []int) error {
	h, err := task.NewHandler(taskRepository)
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
