package usecases

import (
	"os"
	"path/filepath"
)

// Config is an application configuration.
type Config struct {
	HideReminder bool
	HidePriority bool
	TaskFilePath string
}

// ConfigFile is information of config file location.
type ConfigFile struct {
	ConfigName string
	ConfigType string
	ConfigPath string
}

// DefaultConfig is a default config.
var DefaultConfig = Config{
	HideReminder: false,
	HidePriority: false,
	TaskFilePath: filepath.Join(findDataDir(), "todo.json"),
}

// config is a protected member in usecases, which is readable from the other usecases.
var config Config

// SetConfig is a setter to config.
func SetConfig(c Config) {
	config = c
}

// GetTaskFilePath is a getter config.TaskFilePath.
func GetTaskFilePath() string {
	return config.TaskFilePath
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
		xdgDataHome     = "XDG_DATA_HOME"
		appDir          = "todo"
		defaultDataHome = ".local/share/"
	)

	var homeDir, _ = os.UserHomeDir()
	configPath := filepath.Join(homeDir, defaultDataHome, appDir)

	if dataHome := os.Getenv(xdgDataHome); dataHome != "" {
		configPath = filepath.Join(dataHome, appDir)
	}

	return configPath
}
