package task_test

import (
	"testing"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
)

func TestSetReminder(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"TrueMacos", "macos", true},
		{"Fail", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := task.IsValidReminder(tt.in)

			if got != tt.want {
				t.Fatalf("want %v, but %v", tt.want, got)
			}
		})
	}
}
