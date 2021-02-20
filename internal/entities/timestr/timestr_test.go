package timestr_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
)

func TestUnifyLayout(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      string
		wantError bool
		err       error
	}{
		{"SuccessMinutesZeroPadding", "2020/12/04 23:29", "2020/12/4 23:29", false, nil},
		{"SuccessDay", "2020/1/4", "2020/1/4 00:00", false, nil},
		{"SuccessDayZeroPadding", "2020/12/04", "2020/12/4 00:00", false, nil},
		{"HasErrorInvalidLayout", "invalid layout", "", true, errors.New("invalid time layout: [minutes layout]: parsing time \"invalid layout\" as \"2006/1/2 15:04\": cannot parse \"invalid layout\" as \"2006\", [day layout]: parsing time \"invalid layout\" as \"2006/1/2\": cannot parse \"invalid layout\" as \"2006\"")},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := timestr.UnifyLayout(tt.in)

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

func TestModifyTime(t *testing.T) {
	tests := []struct {
		name      string
		duration  string
		base      string
		want      string
		wantError bool
		err       error
	}{
		{"Success", "1h1m", "2020/1/4 02:09", "2020/1/4 03:10", false, nil},
		{"HasErrorInvalidDuration", "+1D1h1m", "2020/1/4 02:09", "", true, errors.New("time: unknown unit \"D\" in duration \"+1D1h1m\"")},
		{"HasErrorInvalidBasetime", "1h", "invalid layout", "", true, errors.New("parsing time \"invalid layout\" as \"2006/1/2 15:04\": cannot parse \"invalid layout\" as \"2006\"")},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := timestr.ModifyTime(tt.duration, tt.base)

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

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      time.Time
		wantError bool
		err       error
	}{
		{"SuccessMinutesZeroPadding", "2020/12/04 23:29", time.Date(2020, 12, 4, 23, 29, 0, 0, time.Local), false, nil},
		{"SuccessDay", "2020/1/4", time.Date(2020, 1, 4, 0, 0, 0, 0, time.Local), false, nil},
		{"SuccessDayZeroPadding", "2020/12/04", time.Date(2020, 12, 4, 0, 0, 0, 0, time.Local), false, nil},
		{"HasErrorInvalidLayout", "invalid layout", time.Time{}, true, errors.New("invalid time layout: [minutes layout]: parsing time \"invalid layout\" as \"2006/1/2 15:04\": cannot parse \"invalid layout\" as \"2006\", [day layout]: parsing time \"invalid layout\" as \"2006/1/2\": cannot parse \"invalid layout\" as \"2006\"")},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := timestr.Parse(tt.in)

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
