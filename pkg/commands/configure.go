package commands

import (
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// List is a function that show task list.
var configureCmd = &cobra.Command{
	Use:     "configure",
	Short:   "Configure your todo-cli",
	Aliases: []string{"conf", "config"},
	RunE:    configureHandler,
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().Bool("hide_reminder", false, "hide a reminder column when show a list")
	configureCmd.Flags().Bool("hide_priority", false, "hide a priority column when show a list")
	configureCmd.Flags().String("task_file_path", "", "the absolute path of your task json file. if not exist, create new directories and a file at the given path.")
}

// configureHandler overwrites viper config and writes out.
func configureHandler(c *cobra.Command, args []string) error {
	if c.Flags().Changed("hide_reminder") {
		f, err := c.Flags().GetBool("hide_reminder")
		if err != nil {
			return err
		}
		viper.Set("HideReminder", f)
	}

	if c.Flags().Changed("hide_priority") {
		f, err := c.Flags().GetBool("hide_priority")
		if err != nil {
			return err
		}
		viper.Set("HidePriority", f)
	}

	if c.Flags().Changed("task_file_path") {
		f, err := c.Flags().GetString("task_file_path")
		if err != nil {
			return err
		}
		if err := usecases.ValidateTaskFilePath(f); err != nil {
			return err
		}
		viper.Set("TaskFilePath", f)
	}

	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}
