package commands

import (
	"fmt"

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
	configureCmd.Flags().String("slack_webhook_url", "", "slack webhookURL which can be gotten from Incoming Webhook")
	configureCmd.Flags().String("slack_mention_to", "", "configure the user name whom mention to in slack message")
	configureCmd.Flags().Bool("reset_config", false, "reset config to default")
	configureCmd.Flags().Bool("show_config", false, "show your config after setting up given change")
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

	if c.Flags().Changed("slack_webhook_url") {
		f, err := c.Flags().GetString("slack_webhook_url")
		if err != nil {
			return err
		}
		viper.Set("SlackWebhookURL", f)
	}

	if c.Flags().Changed("slack_mention_to") {
		f, err := c.Flags().GetString("slack_mention_to")
		if err != nil {
			return err
		}
		viper.Set("SlackMentionTo", f)
	}

	if c.Flags().Changed("reset_config") {
		f, err := c.Flags().GetBool("reset_config")
		if err != nil {
			return err
		}

		if f {
			viper.Set("HideReminder", usecases.DefaultConfig.HideReminder)
			viper.Set("HidePriority", usecases.DefaultConfig.HidePriority)
			viper.Set("TaskFilePath", usecases.DefaultConfig.TaskFilePath)
			viper.Set("SlackWebhookURL", usecases.DefaultConfig.SlackWebhookURL)
			viper.Set("SlackMentionTo", usecases.DefaultConfig.SlackMentionTo)
		}
	}

	if c.Flags().Changed("show_config") {
		f, err := c.Flags().GetBool("show_config")
		if err != nil {
			return err
		}

		if f {
			for k, v := range viper.AllSettings() {
				fmt.Printf("%s: %v\n", k, v)
			}
		}
	}

	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}
