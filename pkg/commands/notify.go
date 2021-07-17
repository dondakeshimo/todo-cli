package commands

import (
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/urfave/cli/v2"
)

// Notify is a function that notify a task.
// It should be called internal only.
func Notify(c *cli.Context) error {
	uuid := c.String("uuid")

	return usecases.Notify(uuid)
}
