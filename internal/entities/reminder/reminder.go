package reminder

import (
	"fmt"
)

// Reminder is a Value Object.
type Reminder string

// NewReminder is a constructor for Rminder.
func NewReminder(str string) (Reminder, error) {
	if !isValidReminder(str) {
		return "", fmt.Errorf("invalid reminder: %s", str)
	}

	return Reminder(str), nil
}

// IsValidReminder is a function that judge a reminder valid.
func isValidReminder(str string) bool {
	allowReminders := [1]string{"macos"}
	for _, a := range allowReminders {
		if str == a {
			return true
		}
	}

	return false
}
