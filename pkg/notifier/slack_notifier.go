package notifier

import (
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

type SlackNotifier struct {
	Scheduler scheduler.Scheduler
}

func (sn *SlackNotifier) Push (r *Request) error {
	return nil
}
