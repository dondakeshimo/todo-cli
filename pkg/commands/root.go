package commands

import (
	"log"
	"os"

	"github.com/dondakeshimo/todo-cli/pkg/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage Your TODO",
}

// init is a Main Component, which injects dependencies.
func init() {
	cobra.OnInitialize(initConfig, injectDependencies)
}

func initConfig() {
	c := usecases.FindConfigFile()
	viper.SetConfigName(c.ConfigName)
	viper.SetConfigType(c.ConfigType)
	viper.AddConfigPath(c.ConfigPath)

	viper.SetDefault("HideReminder", usecases.DefaultConfig.HideReminder)
	viper.SetDefault("HidePriority", usecases.DefaultConfig.HidePriority)
	viper.SetDefault("TaskFilePath", usecases.DefaultConfig.TaskFilePath)
	viper.SetDefault("SlackWebhookURL", usecases.DefaultConfig.SlackWebhookURL)
	viper.SetDefault("SlackMentionTo", usecases.DefaultConfig.SlackMentionTo)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// if not exist ConfigPath, make directory
			if _, err := os.Stat(c.ConfigPath); err != nil {
				if err := os.MkdirAll(c.ConfigPath, os.ModePerm); err != nil {
					log.Fatalln(err)
				}
			}

			if err := viper.SafeWriteConfig(); err != nil {
				log.Fatalln(err)
			}
		} else {
			log.Fatalln(err)
		}
	}

	var config usecases.Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}
	usecases.SetConfig(config)
}

func injectDependencies() {
	jc, err := json.NewClient(usecases.GetTaskFilePath())
	if err != nil {
		log.Fatalln(err)
	}
	usecases.SetRepository(jc)
}

// Execute invoke cobra.Command.Execute.
func Execute() error {
	return rootCmd.Execute()
}
