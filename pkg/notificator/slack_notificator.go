package notificator

import (
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

type SlackNotificator struct {
	Scheduler scheduler.Scheduler
}

func (sn *SlackNotificator) Push () (*Response, error) {
	return nil, nil
}
