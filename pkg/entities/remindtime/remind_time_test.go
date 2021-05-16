package remindtime_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dondakeshimo/todo-cli/pkg/entities/remindtime"
)

func TestNewRemindTime(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      remindtime.RemindTime
		wantError bool
		err       error
	}{
		{"SuccessMinutesZeroPadding", "2020/12/04 23:29", remindtime.RemindTime("2020/12/4 23:29"), false, nil},
		{"SuccessDay", "2020/1/4", remindtime.RemindTime("2020/1/4 00:00"), false, nil},
		{"SuccessDayZeroPadding", "2020/12/04", remindtime.RemindTime("2020/12/4 00:00"), false, nil},
		{"HasErrorInvalidLayout", "invalid layout", "", true, errors.New("invalid time layout: [minutes layout]: parsing time \"invalid layout\" as \"2006/1/2 15:04\": cannot parse \"invalid layout\" as \"2006\", [day layout]: parsing time \"invalid layout\" as \"2006/1/2\": cannot parse \"invalid layout\" as \"2006\"")},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := remindtime.NewRemindTime(tt.in)

			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}

			if tt.wantError && err.Error() != tt.err.Error() {
				t.Fatalf("want %#v, but %#v", tt.err.Error(), err.Error())
			}

			if !tt.wantError && got != tt.want {
				t.Fatalf("want %q, but %q", tt.want, got)
			}
		})
	}
}

func TestAddTime(t *testing.T) {
	tests := []struct {
		name         string
		relativeTime remindtime.RelativeTime
		remindTime   remindtime.RemindTime
		want         remindtime.RemindTime
		wantError    bool
		err          error
	}{
		{
			name:         "SuccessTaskBase",
			relativeTime: remindtime.RelativeTime{2 * time.Minute, true},
			remindTime:   remindtime.RemindTime("2020/12/31 23:58"),
			want:         remindtime.RemindTime("2021/1/1 00:00"),
			wantError:    false,
			err:          nil,
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.remindTime.AddTime(tt.relativeTime)

			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}

			if tt.wantError && err.Error() != tt.err.Error() {
				t.Fatalf("want %#v, but %#v", tt.err.Error(), err.Error())
			}

			if !tt.wantError && got != tt.want {
				t.Fatalf("want %q, but %q", tt.want, got)
			}
		})
	}
}
