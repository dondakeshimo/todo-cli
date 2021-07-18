package usecases

import (
	"os"
	"strconv"

	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
	"github.com/olekukonko/tablewriter"
)

// List is a function that show task list.
func List() error {
	h, err := task.NewHandler(taskRepository)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false) // for preventing header from being made upper case automatically

	table.SetHeader([]string{"ID", "Task", "RemindTime", "Reminder", "Priority"})

	for _, t := range h.GetTasks() {
		table.Append([]string{strconv.Itoa(t.ID()), t.Task(), string(t.RemindTime()), string(t.Reminder()), strconv.Itoa(t.Priority())})
	}

	table.Render()

	return nil
}
