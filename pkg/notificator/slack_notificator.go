package notificator

import (
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

type SlackNotificator struct {
	Scheduler scheduler.Scheduler
}

func (sn *SlackNotificator) Push (r *Request) error {
	return nil
}
