package usecases

import (
	"os"
	"strconv"

	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
	"github.com/olekukonko/tablewriter"
)

type columns struct {
	id         string
	task       string
	remindTime string
	reminder   string
	priority   string
}

// List is a function that show task list.
func List() error {
	h, err := task.NewHandler(taskRepository)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false) // for preventing header from being made upper case automatically

	header := columns{
		id:         "ID",
		task:       "Task",
		remindTime: "RemindTime",
		reminder:   "Reminder",
		priority:   "Priority",
	}
	table.SetHeader(buildRowAccordingToConfig(header))

	for _, t := range h.GetTasks() {
		row := columns{
			id:         strconv.Itoa(t.ID()),
			task:       t.Task(),
			remindTime: string(t.RemindTime()),
			reminder:   string(t.Reminder()),
			priority:   strconv.Itoa(t.Priority()),
		}
		table.Append(buildRowAccordingToConfig(row))
	}

	table.Render()

	return nil
}

// buildColumnAccordingToSetting build a row of table according to hide-setting.
// this function define column order.
func buildRowAccordingToConfig(row columns) []string {
	builtRow := []string{row.id, row.task, row.remindTime}

	if !config.hideReminder {
		builtRow = append(builtRow, row.reminder)
	}

	if !config.hidePriority {
		builtRow = append(builtRow, row.priority)
	}

	return builtRow
}
