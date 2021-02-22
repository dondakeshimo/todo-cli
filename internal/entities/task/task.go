package task

import (
	"fmt"
	"os"

	"github.com/dondakeshimo/todo-cli/internal/values/reminder"
	"github.com/dondakeshimo/todo-cli/internal/values/remindtime"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

// Task is a struct that describe task.
// This struct is a value object.
type Task struct {
	ID         int
	Task       string
	RemindTime remindtime.RemindTime
	UUID       string
	Reminder   reminder.Reminder
}

// TaskJSON is a struct to write/read JSON.
type TaskJSON struct {
	ID         int    `json:"id"`
	Task       string `json:"task"`
	RemindTime string `json:"remindtime"`
	UUID       string `json:"uuid"`
	Reminder   string `json:"reminder"`
}

func (t Task) ToTaskJSON() TaskJSON {
	return TaskJSON{
		ID:         t.ID,
		Task:       t.Task,
		RemindTime: string(t.RemindTime),
		UUID:       t.UUID,
		Reminder:   string(t.Reminder),
	}
}

func (t Task) AlterID(id int) Task {
	return Task{
		ID:         id,
		Task:       t.Task,
		RemindTime: t.RemindTime,
		UUID:       t.UUID,
		Reminder:   t.Reminder,
	}
}

func (tj TaskJSON) ToTask() Task {
	return Task{
		ID:         tj.ID,
		Task:       tj.Task,
		RemindTime: remindtime.RemindTime(tj.RemindTime),
		UUID:       tj.UUID,
		Reminder:   reminder.Reminder(tj.Reminder),
	}
}

// SetReminder is a function that set a reminder of the task.
func (t Task) SetReminder(s scheduler.Scheduler) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	sr := scheduler.Request{
		ID:       t.UUID,
		DateTime: t.RemindTime.ToTime(),
		Command:  fmt.Sprintf("%s notify --uuid %s", exe, t.UUID),
	}

	if err := s.Register(sr); err != nil {
		return err
	}

	return err
}
