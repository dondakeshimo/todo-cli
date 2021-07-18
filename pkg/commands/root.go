package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage Your TODO",
}

// init is a Main Component, which injects dependencies.
func init() {
	jc, err := json.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	usecases.SetRepository(jc)
}

// Execute invoke cobra.Command.Execute.
func Execute() error {
	return rootCmd.Execute()
}
