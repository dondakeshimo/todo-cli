package task_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
	"github.com/dondakeshimo/todo-cli/pkg/scheduler"
)

func TestSetReminder(t *testing.T) {
	exe, err := os.Executable()
	if err != nil {
		t.Fatalf("failed to get executable: %#v", err)
	}

	const invalidTime = "2020/12/05 00:26 invalid"
	_, parseErr := timestr.Parse(invalidTime)

	tests := []struct {
		name           string
		task           task.Task
		request        scheduler.Request
		buildScheduler func(m *scheduler.MockScheduler, r scheduler.Request)
		wantError      bool
		err            error
	}{
		{
			name: "Success",
			task: task.Task{
				UUID:       "uuid",
				RemindTime: "2020/12/05 00:26",
			},
			request: scheduler.Request{
				ID:       "uuid",
				DateTime: time.Date(2020, 12, 5, 00, 26, 00, 00, time.Local),
				Command:  fmt.Sprintf("%s notify --uuid uuid", exe),
			},
			buildScheduler: func(m *scheduler.MockScheduler, r scheduler.Request) {
				m.
					EXPECT().
					Register(r).
					Return(nil)
			},
			wantError: false,
			err:       nil,
		},
		{
			name: "HasErrorTimeParse",
			task: task.Task{
				UUID:       "uuid",
				RemindTime: invalidTime,
			},
			request: scheduler.Request{},
			buildScheduler: func(m *scheduler.MockScheduler, r scheduler.Request) {
				m.
					EXPECT().
					Register(gomock.Any()).
					Times(0)
			},
			wantError: true,
			err:       parseErr,
		},
		{
			name: "HasErrorRegister",
			task: task.Task{
				UUID:       "uuid",
				RemindTime: "2020/12/05 00:26",
			},
			request: scheduler.Request{
				ID:       "uuid",
				DateTime: time.Date(2020, 12, 5, 00, 26, 00, 00, time.Local),
				Command:  fmt.Sprintf("%s notify --uuid uuid", exe),
			},
			buildScheduler: func(m *scheduler.MockScheduler, r scheduler.Request) {
				m.
					EXPECT().
					Register(r).
					Return(errors.New("error test"))
			},
			wantError: true,
			err:       errors.New("error test"),
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := scheduler.NewMockScheduler(ctrl)
			tt.buildScheduler(s, tt.request)

			err := tt.task.SetReminder(s)

			if tt.wantError && err.Error() != tt.err.Error() {
				t.Fatalf("want %#v, but %#v", tt.err.Error(), err.Error())
			}
		})
	}

}

func TestIsValidReminder(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"TrueMacos", "macos", true},
		{"Fail", "invalid", false},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := task.IsValidReminder(tt.in)

			if got != tt.want {
				t.Fatalf("want %v, but %v", tt.want, got)
			}
		})
	}
}
