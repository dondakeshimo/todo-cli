package commands

import (
	"fmt"

	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
)

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Notify a task (basicaly be used by system)",
	RunE:  notifyHandler,
}

func init() {
	rootCmd.AddCommand(notifyCmd)

	notifyCmd.Flags().String("uuid", "", "task's UUID")

	if err := notifyCmd.MarkFlagRequired("uuid"); err != nil {
		// NOTE: err is set when "uuid" is not found in flags.
		//       so, this block never work.
		fmt.Println(err)
	}
}

// notifyHandler invoke usecases.Notify with parameter from cli.
func notifyHandler(c *cobra.Command, args []string) error {
	uuid, err := c.Flags().GetString("uuid")
	if err != nil {
		return err
	}

	return usecases.Notify(uuid)
}
