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
	id         int
	task       string
	remindTime remindtime.RemindTime
	uuid       string
	reminder   reminder.Reminder
}

// NewTask is a constructor for Task.
func NewTask(i int, t string, rt remindtime.RemindTime, uuid string, rm reminder.Reminder) Task {
	return Task{
		id:         i,
		task:       t,
		remindTime: rt,
		uuid:       uuid,
		reminder:   rm,
	}
}

// ID is a getter for id.
func (t Task) ID() int {
	return t.id
}

// Task is a getter for task.
func (t Task) Task() string {
	return t.task
}

// RemindTime is a getter for remindTime.
func (t Task) RemindTime() remindtime.RemindTime {
	return t.remindTime
}

// UUID is a getter for uuid.
func (t Task) UUID() string {
	return t.uuid
}

// Reminder is a getter for reminder.
func (t Task) Reminder() reminder.Reminder {
	return t.reminder
}

func (t Task) alterID(id int) Task {
	return Task{
		id:         id,
		task:       t.task,
		remindTime: t.remindTime,
		uuid:       t.uuid,
		reminder:   t.reminder,
	}
}

// SetReminder is a function that set a reminder of the task.
func (t Task) SetReminder(s scheduler.Scheduler) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	sr := scheduler.Request{
		ID:       t.uuid,
		DateTime: t.remindTime.ToTime(),
		Command:  fmt.Sprintf("%s notify --uuid %s", exe, t.uuid),
	}

	if err := s.Register(sr); err != nil {
		return err
	}

	s.ClearExpired()
	return nil
}

// RemoveReminder is a function that remove a reminder of the task.
func (t Task) RemoveReminder(s scheduler.Scheduler) error {
	if err := s.RemoveWithID(t.uuid); err != nil {
		return err
	}

	s.ClearExpired()
	return nil
}
