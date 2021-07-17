package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage Your TODO",
}

// Execute invoke cobra.Command.Execute.
func Execute() error {
	return rootCmd.Execute()
}
