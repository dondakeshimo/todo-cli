package commands

import (
	"sort"
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
	if err := w.Write(header); err != nil {
		return err
	}

	tasks := h.GetTasks()
	sort.SliceStable(tasks, func(i, j int) bool {
		switch c.String("order") {
		case "priority":
			return tasks[i].Priority() < tasks[j].Priority()
		case "id":
			return tasks[i].ID() < tasks[j].ID()
		default: // same as the case of "priority"
			return tasks[i].Priority() < tasks[j].Priority()
		}
	})
	for _, t := range tasks {
		if err := w.Write([]string{strconv.Itoa(t.ID()), t.Task(), string(t.RemindTime()), string(t.Reminder()), strconv.Itoa(t.Priority())}); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}
