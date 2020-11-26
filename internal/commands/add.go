package commands

import (
	"fmt"
	"runtime"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	d, err := timestr.Validate(c.String("remind_time"))
	if err != nil {
		return err
	}

	r := c.String("reminder")
	if r != "" && !task.IsValidReminder(r){
		return fmt.Errorf("invalid reminder: %s", r)
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	t := &task.Task{
		Task:       c.String("task"),
		RemindTime: d,
		UUID:       uu.String(),
		Reminder:   r,
	}

	h.AppendTask(t)

	if err := h.Write(); err != nil {
		return err
	}

	// when do not remind, do early return
	if r == "" {
		return nil
	}

	var s scheduler.Scheduler
	// TODO: when adjusting the other os, add condition
	if runtime.GOOS == "darwin" {
		s = scheduler.NewLaunchdScheduler()
	} else {
		return nil
	}

	if err := t.SetReminder(s); err != nil {
		return err
	}

	return nil
}
