package commands

import (
	"strconv"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/writer"
	"github.com/urfave/cli/v2"
)

// List is a function that show task list.
func List(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	w := writer.NewTSVWriter()
	defer w.Flush()

	header := []string{"ID", "Task", "RemindTime", "reminder"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, t := range h.GetTasks() {
		if err := w.Write([]string{strconv.Itoa(t.ID), t.Task, t.RemindTime, t.Reminder}); err != nil {
			return err
		}
	}

	return nil
}
