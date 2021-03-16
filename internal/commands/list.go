package commands

import (
	"os"
	"strconv"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/gateways/json"
	"github.com/olekukonko/tablewriter"
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false) // for preventing header from being made upper case automatically

	table.SetHeader([]string{"ID", "Task", "RemindTime", "reminder", "Priority"})

	for _, t := range h.GetTasks() {
		table.Append([]string{strconv.Itoa(t.ID()), t.Task(), string(t.RemindTime()), string(t.Reminder()), strconv.Itoa(t.Priority())})
	}

	table.Render()

	return nil
}
