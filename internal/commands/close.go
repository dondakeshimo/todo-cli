package commands

import (
	"github.com/dondakeshimo/todo-cli/internal/usecases"
	"github.com/urfave/cli/v2"
)

// Close is a function that close a task or tasks.
func Close(c *cli.Context) error {
	ids := c.IntSlice("ids")

	return usecases.Close(ids)
}
