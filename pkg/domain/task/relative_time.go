package task

import (
	"fmt"
	"strings"
	"time"
)

// RelativeTime is time.Duration that had paresd from a certain layout string.
type RelativeTime struct {
	RelativeTime time.Duration
	IsTaskBase   bool
}

// NewRelativeTime is a constructor for RelativeTime.
func NewRelativeTime(str string) (RelativeTime, error) {
	if IsRelativeToTask(str) {
		rt, err := newRelativeTimeRelativeToTask(str)
		if err != nil {
			return rt, err
		}
		return rt, nil
	}

	if IsRelativeToNow(str) {
		rt, err := newRelativeTimeRelativeToNow(str)
		if err != nil {
			return rt, err
		}
		return rt, nil
	}

	return RelativeTime{}, fmt.Errorf("could not convert to time.Duration: %s", str)
}

// IsValidRelativeTime is a logical function that confirm RelativeTime is constructed from str.
func IsValidRelativeTime(str string) bool {
	return IsRelativeToTask(str) || IsRelativeToNow(str)
}

// IsRelativeToTask is a logical function that confirm valid relativetime to task.
func IsRelativeToTask(str string) bool {
	return strings.HasPrefix(str, "task")
}

// IsRelativeToNow is a logical function that confirm valid relativetime to now.
func IsRelativeToNow(str string) bool {
	return strings.HasPrefix(str, "+") || strings.HasPrefix(str, "now")
}

func newRelativeTimeRelativeToTask(str string) (RelativeTime, error) {
	str = strings.Replace(str, "task", "", 1)
	td, err := time.ParseDuration(str)
	if err != nil {
		return RelativeTime{}, err
	}
	return RelativeTime{RelativeTime: td, IsTaskBase: true}, nil
}

func newRelativeTimeRelativeToNow(str string) (RelativeTime, error) {
	str = strings.Replace(str, "now", "", 1)
	td, err := time.ParseDuration(str)
	if err != nil {
		return RelativeTime{}, err
	}
	return RelativeTime{RelativeTime: td, IsTaskBase: false}, nil
}
