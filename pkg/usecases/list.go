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
	group      string
	remindTime string
	reminder   string
	priority   string
}

// List is a function that show task list.
func List(group string) error {
	h, err := task.NewHandler(taskRepository)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false) // for preventing header from being made upper case automatically
	table.SetColWidth(config.ColumnWidth)

	header := columns{
		id:         "ID",
		task:       "Task",
		group:      "Group",
		remindTime: "RemindTime",
		reminder:   "Reminder",
		priority:   "Priority",
	}
	table.SetHeader(buildRowAccordingToConfig(header))

	tasks := h.GetTasks()
	if group != "" {
		tasks = h.GetTasksInGroup(group)
	}

	for _, t := range tasks {
		row := columns{
			id:         strconv.Itoa(t.ID()),
			task:       t.Task(),
			group:      t.Group(),
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
	builtRow := []string{row.id, row.task}

	if !config.HideRemindTime {
		builtRow = append(builtRow, row.remindTime)
	}

	if !config.HideGroup {
		builtRow = append(builtRow, row.group)
	}

	if !config.HideReminder {
		builtRow = append(builtRow, row.reminder)
	}

	if !config.HidePriority {
		builtRow = append(builtRow, row.priority)
	}

	return builtRow
}
