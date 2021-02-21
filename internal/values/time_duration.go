package values

import (
	"fmt"
	"strings"
	"time"
)

// TimeDuration is time.Duration that had paresd from a certain layout string.
type TimeDuration time.Duration

// NewTimeDuration is a constructor for TimeDuration.
func NewTimeDuration(str string) (TimeDuration, error) {
	if isRelativeToTask(str) {
		td, err := newTimeDurationRelativeToTask(str)
		if err != nil {
			return TimeDuration(0), err
		}
		return TimeDuration(td), nil
	}

	if isRelativeToNow(str) {
		td, err := newTimeDurationRelativeToNow(str)
		if err != nil {
			return TimeDuration(0), err
		}
		return TimeDuration(td), nil
	}

	return TimeDuration(0), fmt.Errorf("could not convert to time.Duration: %s", str)
}

// IsValidTimeDuration is a logical function that confirm TimeDuration is constructed from str.
func IsValidTimeDuration(str string) bool {
	return isRelativeToTask(str) || isRelativeToNow(str)
}

func isRelativeToTask(str string) bool {
	return strings.HasPrefix(str, "task")
}

func isRelativeToNow(str string) bool {
	return strings.HasPrefix(str, "+") || strings.HasPrefix(str, "now")
}

func newTimeDurationRelativeToTask(str string) (TimeDuration, error) {
	str = strings.Replace(str, "task", "", 1)
	td, err := time.ParseDuration(str)
	if err != nil {
		return TimeDuration(0), err
	}
	return TimeDuration(td), nil
}

func newTimeDurationRelativeToNow(str string) (TimeDuration, error) {
	str = strings.Replace(str, "now", "", 1)
	td, err := time.ParseDuration(str)
	if err != nil {
		return TimeDuration(0), err
	}
	return TimeDuration(td), nil
}
