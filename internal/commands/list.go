package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/usecases"
	"github.com/urfave/cli/v2"
)

// List is a function that show task list.
func List(c *cli.Context) error {
	return usecases.List()
}
