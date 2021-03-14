package commands

import (
	"strconv"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/writer"
	"github.com/urfave/cli/v2"
)

// List is a function that show task list.
func List(c *cli.Context) error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
	if err != nil {
		return err
	}

	w := writer.NewTSVWriter()

	header := []string{"ID", "Task", "RemindTime", "reminder", "Priority"}
	if err := w.Header(header); err != nil {
		return err
	}

	for _, t := range h.GetTasks() {
		if err := w.Write([]string{strconv.Itoa(t.ID()), t.Task(), string(t.RemindTime()), string(t.Reminder()), strconv.Itoa(t.Priority())}); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}
