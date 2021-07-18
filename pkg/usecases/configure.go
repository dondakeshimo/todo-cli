package usecases

// configStruct is an application configuration.
type configStruct struct {
	hideReminder bool
	hidePriority bool
}

// config is a protected member in usecases, which is readable from the other usecases.
var config configStruct

// SetConfig is a setter to config.
func SetConfig(hideReminder bool, hidePriority bool) {
	config.hideReminder = hideReminder
	config.hidePriority = hidePriority
}
