package task_test

import (
	"errors"
	"testing"

	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
)

func TestIsValidReminder(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      task.Reminder
		wantError bool
		err       error
	}{
		{"TrueMacos", "macos", task.Reminder("macos"), false, nil},
		{"HasErrorInvalid", "invalid", "", true, errors.New("invalid reminder: invalid")},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := task.NewReminder(tt.in)

			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}

			if tt.wantError && err.Error() != tt.err.Error() {
				t.Fatalf("want %#v, but %#v", tt.err.Error(), err.Error())
			}

			if !tt.wantError && got != tt.want {
				t.Fatalf("want %#v, but %#v", tt.want, got)
			}
		})
	}
}
