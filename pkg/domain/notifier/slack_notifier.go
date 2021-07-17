package notifier

import (
	"github.com/dondakeshimo/todo-cli/pkg/domain/scheduler"
)

// SlackNotifier is a struct that notify with slack.
type SlackNotifier struct {
	Scheduler scheduler.Scheduler
}

// Push is a function that push notification.
func (sn *SlackNotifier) Push(r *Request) error {
	return nil
}
