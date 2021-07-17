package commands

import (
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/urfave/cli/v2"
)

// List is a function that show task list.
func List(c *cli.Context) error {
	return usecases.List()
}
