package notificator

import (
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

type OssascriptNotificator struct {
	Scheduler scheduler.Scheduler
}

func (on *OssascriptNotificator) Push () (*Response, error) {
	return nil, nil
}
