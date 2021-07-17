package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:     "close",
	Short:   "Close tasks",
	Aliases: []string{"c"},
	RunE:    closeHandler,
}

func init() {
	rootCmd.AddCommand(closeCmd)
	closeCmd.Flags().IntSliceP("ids", "i", []int{}, "task's IDs  (ex. $ todo c -i 2,3,5 )")
}

// closeHandler invoke usecases.Close with parameter from cli.
func closeHandler(c *cobra.Command, args []string) error {
	ids, err := c.Flags().GetIntSlice("ids")
	if err != nil {
		return err
	}

	if len(ids) < 1 {
		return fmt.Errorf("ids to close are needed")
	}

	return usecases.Close(ids)
}
