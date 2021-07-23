package commands

import (
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
)

// List is a function that show task list.
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List tasks",
	Aliases: []string{"l"},
	RunE:    listHandler,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("group", "g", "", "group you want to check")
}

// listHandler invoke usecases.List
func listHandler(c *cobra.Command, args []string) error {
	group, err := c.Flags().GetString("group")
	if err != nil {
		return err
	}

	return usecases.List(group)
}
