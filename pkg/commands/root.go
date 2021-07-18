package commands

import (
	"log"

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
	cobra.OnInitialize(initConfig)

	jc, err := json.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	usecases.SetRepository(jc)
}

func initConfig() {
	c := usecases.FindConfigFile()
	viper.SetConfigName(c.ConfigName)
	viper.SetConfigType(c.ConfigType)
	viper.AddConfigPath(c.ConfigPath)

	viper.SetDefault("HideReminder", usecases.DefaultConfig.HideReminder)
	viper.SetDefault("HidePriority", usecases.DefaultConfig.HidePriority)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfig()
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

// Execute invoke cobra.Command.Execute.
func Execute() error {
	return rootCmd.Execute()
}
