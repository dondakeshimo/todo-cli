package timestr_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		want      string
		wantError bool
		err       error
	}{
		{"SuccessMinutes", "2020/12/04 23:29", "2020/12/04 23:29", false, nil},
		{"SuccessDay", "2020/12/04", "2020/12/04 00:00", false, nil},
		{"HasErrorInvalidLayout", "invalid layout", "", true, errors.New("invalid layout: [parsing time \"invalid layout\" as \"2006/01/02 15:04\": cannot parse \"invalid layout\" as \"2006\", parsing time \"invalid layout\" as \"2006/01/02\": cannot parse \"invalid layout\" as \"2006\"]")},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := timestr.Validate(tt.in)

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
		{"SuccessMinutes", "2020/12/04 23:29", time.Date(2020, 12, 4, 23, 29, 0, 0, time.Local), false, nil},
		{"SuccessMinutes", "2020/12/04", time.Date(2020, 12, 4, 0, 0, 0, 0, time.Local), false, nil},
		{"HasErrorInvalidLayout", "invalid layout", time.Time{}, true, errors.New("invalid layout: [parsing time \"invalid layout\" as \"2006/01/02 15:04\": cannot parse \"invalid layout\" as \"2006\", parsing time \"invalid layout\" as \"2006/01/02\": cannot parse \"invalid layout\" as \"2006\"]")},
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
