package usecases

// configStruct is an application configuration.
type Config struct {
	HideReminder bool
	HidePriority bool
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
