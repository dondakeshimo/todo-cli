package usecases

import (
	"os"
	"path/filepath"
)

// configStruct is an application configuration.
type Config struct {
	HideReminder bool
	HidePriority bool
}

type ConfigFile struct {
	ConfigName string
	ConfigType string
	ConfigPath string
}

var DefaultConfig = Config{
	HideReminder: false,
	HidePriority: false,
}

// config is a protected member in usecases, which is readable from the other usecases.
var config Config

// SetConfig is a setter to config.
func SetConfig(c Config) {
	config = c
}

// FindConfigFile return ConfigFile in which ConfigPath is set according to XDG_DATA_HOME.
func FindConfigFile() ConfigFile {
	const (
		defaultConfigName = "config"
		defaultConfigType = "yaml"
	)

	configPath := findDataDir()

	return ConfigFile{
		ConfigName: defaultConfigName,
		ConfigType: defaultConfigType,
		ConfigPath: configPath,
	}
}

func findDataDir() string {
	const (
		xdg_data_home   = "XDG_DATA_HOME"
		appDir          = "todo"
		defaultDataHome = ".local/share/"
	)

	var homeDir, _ = os.UserHomeDir()
	configPath := filepath.Join(homeDir, defaultDataHome, appDir)

	if dataHome := os.Getenv(xdg_data_home); dataHome != "" {
		configPath = filepath.Join(dataHome, appDir)
	}

	return configPath
}
