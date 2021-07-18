package commands

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dondakeshimo/todo-cli/pkg/gateways/json"
	"github.com/dondakeshimo/todo-cli/pkg/usecases"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage Your TODO",
}

type configFile struct {
	configName string
	configType string
	configPath string
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
	c := findConfigFile()
	viper.SetConfigName(c.configName)
	viper.SetConfigType(c.configType)
	viper.AddConfigPath(c.configPath)

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

// FindConfigFile return ConfigFile in which ConfigPath is set according to XDG_DATA_HOME.
func findConfigFile() configFile {
	const (
		xdg_data_home     = "XDG_DATA_HOME"
		appDir            = "todo"
		defaultConfigName = "config.yaml"
		defaultConfigType = "yaml"
		defaultDataHome   = ".local/share/"
	)

	var homeDir, _ = os.UserHomeDir()
	configPath := filepath.Join(homeDir, defaultDataHome, appDir)

	if dataHome := os.Getenv(xdg_data_home); dataHome != "" {
		configPath = filepath.Join(dataHome, appDir)
	}

	return configFile{
		configName: defaultConfigName,
		configType: defaultConfigType,
		configPath: configPath,
	}
}

// Execute invoke cobra.Command.Execute.
func Execute() error {
	return rootCmd.Execute()
}
