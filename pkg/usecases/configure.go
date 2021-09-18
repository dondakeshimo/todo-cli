package usecases

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config is an application configuration.
type Config struct {
	HideReminder    bool
	HideRemindTime  bool
	HidePriority    bool
	HideGroup       bool
	TaskFilePath    string
	SlackWebhookURL string
	SlackMentionTo  string
	ColumnWidth     int
}

// ConfigFile is information of config file location.
type ConfigFile struct {
	ConfigName string
	ConfigType string
	ConfigPath string
}

// DefaultConfig is a default config.
var DefaultConfig = Config{
	HideReminder:    false,
	HideRemindTime:  false,
	HidePriority:    false,
	HideGroup:       false,
	TaskFilePath:    filepath.Join(findDataDir(), "todo.json"),
	SlackWebhookURL: "",
	SlackMentionTo:  "",
	ColumnWidth:     30,
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

// ValidateTaskFilePath judges that a given path is valid for TaskFilePath.
func ValidateTaskFilePath(path string) error {
	// only allow absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if path != absPath {
		return fmt.Errorf("TaskFilePath only support an absolute path")
	}

	// only allow JSON file
	fInfo, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			// it is invalid that the file is exist and occurs an error
			return err
		}
	}
	if fInfo != nil {
		if fInfo.IsDir() {
			return fmt.Errorf("%s is a directory", path)
		}
	}

	return nil
}
