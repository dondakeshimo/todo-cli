package reminder

import (
	"fmt"
)

// Reminder is a Value Object.
type Reminder string

const (
	MacOS = Reminder("macos")
	Slack = Reminder("slack")
)

// NewReminder is a constructor for Rminder.
func NewReminder(str string) (Reminder, error) {
	if !isValidReminder(str) {
		return "", fmt.Errorf("invalid reminder: %s", str)
	}

	return Reminder(str), nil
}

// IsValidReminder is a function that judge a reminder valid.
func isValidReminder(str string) bool {
	allowReminders := [2]Reminder{MacOS, Slack}
	for _, a := range allowReminders {
		if str == string(a) {
			return true
		}
	}

	return false
}
