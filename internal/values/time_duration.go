package values

import (
	"fmt"
	"time"
	"strings"
)

type TimeDuration time.Duration

func NewTimeDuration(str string) (TimeDuration, error) {
	if isRelativeToTask(str) {
		td, err := NewTimeDurationRelativeToTask(str)
		if err != nil {
			return TimeDuration(0), err
		}
		return TimeDuration(td), nil
	}

	if isRelativeToNow(str) {
		td, err := NewTimeDurationRelativeToNow(str)
		if err != nil {
			return TimeDuration(0), err
		}
		return TimeDuration(td), nil
	}

	return TimeDuration(0), fmt.Errorf("could not convert to time.Duration: %s", str)
}

func IsValidTimeDuration(str string) bool {
	return isRelativeToTask(str) || isRelativeToNow(str)
}

func isRelativeToTask(str string) bool {
	return strings.HasPrefix(str, "task")
}

func isRelativeToNow(str string) bool {
	return strings.HasPrefix(str, "+") || strings.HasPrefix(str, "now")
}

func NewTimeDurationRelativeToTask(str string) (TimeDuration, error) {
	str = strings.Replace(str, "task", "", 1)
	td, err := time.ParseDuration(str)
	if err != nil {
		return TimeDuration(0), err
	}
	return TimeDuration(td), nil
}

func NewTimeDurationRelativeToNow(str string) (TimeDuration, error) {
	str = strings.Replace(str, "now", "", 1)
	td, err := time.ParseDuration(str)
	if err != nil {
		return TimeDuration(0), err
	}
	return TimeDuration(td), nil
}
