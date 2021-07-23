package task_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dondakeshimo/todo-cli/pkg/domain/remindtime"
)

func TestIsValidRelativeTime(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"TrueRelativeToTask", "task-1h1m", true},
		{"TrueRelativeToNow", "now+1h1m", true},
		{"TrueRelativeToNowOmmision", "+1h1m", true},
		{"FalseEmpty", "", false},
		{"FalseInvalid", "1h1m", false},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := remindtime.IsValidRelativeTime(tt.str)

			if got != tt.want {
				t.Fatalf("want %v, but %v", tt.want, got)
			}
		})
	}
}

func TestNewRelativeTime(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		want      remindtime.RelativeTime
		wantError bool
		err       error
	}{
		{
			name:      "SuccessTaskBase",
			str:       "task-1h1m",
			want:      remindtime.RelativeTime{-(time.Hour + time.Minute), true},
			wantError: false,
			err:       nil,
		},
		{
			name:      "SuccessTaskNow",
			str:       "now+1h1m",
			want:      remindtime.RelativeTime{time.Hour + time.Minute, false},
			wantError: false,
			err:       nil,
		},
		{
			name:      "HasErrorNoBase",
			str:       "1h1m",
			want:      remindtime.RelativeTime{},
			wantError: true,
			err:       errors.New("could not convert to time.Duration: 1h1m"),
		},
		{
			name:      "HasErrorParseTimeDuration",
			str:       "now+invalid",
			want:      remindtime.RelativeTime{},
			wantError: true,
			err:       errors.New("time: invalid duration \"+invalid\""),
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := remindtime.NewRelativeTime(tt.str)
			t.Logf("got: %#v, err: %#v", got, err)

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
