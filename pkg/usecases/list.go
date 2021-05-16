package usecases

import (
	"os"
	"strconv"

	"github.com/dondakeshimo/todo-cli/pkg/entities/task"
	"github.com/dondakeshimo/todo-cli/pkg/gateways/json"
	"github.com/olekukonko/tablewriter"
)

// List is a function that show task list.
func List() error {
	jc, err := json.NewClient()
	if err != nil {
		return err
	}

	h, err := task.NewHandler(jc)
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
