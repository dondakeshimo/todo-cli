package task_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
)

func TestGetTask(t *testing.T) {
	tasks := []*task.Task{
		{ID: 1, UUID: "uuid1"},
		{ID: 2, UUID: "uuid2"},
		{ID: 3, UUID: "uuid3"},
	}

	handler := task.Handler{}

	for _, ts := range tasks {
		handler.AppendTask(ts)
	}

	tests := []struct {
		name string
		id   int
		want *task.Task
	}{
		{
			name: "Success",
			id:   2,
			want: tasks[1],
		},
		{
			name: "FailOverRange",
			id:   4,
			want: nil,
		},
		{
			name: "FailMinus",
			id:   -1,
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := handler.GetTask(tt.id)

			if got != tt.want {
				t.Fatalf("want %#v, but %#v", tt.want, got)
			}
		})
	}
}

func TestGetTasks(t *testing.T) {
	tasks := []*task.Task{
		{ID: 1, UUID: "uuid1"},
		{ID: 2, UUID: "uuid2"},
		{ID: 3, UUID: "uuid3"},
	}

	handler := task.Handler{}

	for _, ts := range tasks {
		handler.AppendTask(ts)
	}

	tests := []struct {
		name string
		want []*task.Task
	}{
		{
			name: "Success",
			want: tasks,
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := handler.GetTasks()

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("want %#v, but %#v", tt.want, got)
			}
		})
	}
}

func TestRemoveTask(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want []*task.Task
	}{
		{
			name: "Success",
			id:   2,
			want: []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid3"},
			},
		},
		{
			name: "FailOverRange",
			id:   4,
			want: []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid2"},
				{ID: 3, UUID: "uuid3"},
			},
		},
		{
			name: "FailMinus",
			id:   -1,
			want: []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid2"},
				{ID: 3, UUID: "uuid3"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tasks := []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid2"},
				{ID: 3, UUID: "uuid3"},
			}

			handler := task.Handler{}
			for _, ts := range tasks {
				handler.AppendTask(ts)
			}

			handler.RemoveTask(tt.id)
			ts := handler.GetTasks()

			if len(ts) != len(tt.want) {
				t.Fatalf("want %#v, but %#v", tt.want, ts)
			}

			for i := 0; i < len(tt.want); i++ {
				if ts[i].ID != tt.want[i].ID || ts[i].UUID != tt.want[i].UUID {
					t.Fatalf("want %#v, but %#v", tt.want, ts)
				}
			}
		})
	}
}

func TestRemoveTasks(t *testing.T) {
	tests := []struct {
		name string
		ids  []int
		want []*task.Task
		wantError bool
		err error
	}{
		{
			name: "Success",
			ids:  []int{1, 3},
			want: []*task.Task{
				{ID: 1, UUID: "uuid2"},
			},
			wantError: false,
			err: nil,
		},
		{
			name: "SuccessDescending",
			ids:  []int{2, 1},
			want: []*task.Task{
				{ID: 1, UUID: "uuid3"},
			},
			wantError: false,
			err: nil,
		},
		{
			name: "HasContinueOverRange",
			ids:  []int{1, 2, 4},
			want: []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid2"},
				{ID: 3, UUID: "uuid3"},
			},
			wantError: true,
			err: errors.New("no task with id [1 2 4]"),
		},
		{
			name: "HasContinueMinus",
			ids:  []int{1, -1},
			want: []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid2"},
				{ID: 3, UUID: "uuid3"},
			},
			wantError: true,
			err: errors.New("not natural value is invalid [-1 1]"),
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tasks := []*task.Task{
				{ID: 1, UUID: "uuid1"},
				{ID: 2, UUID: "uuid2"},
				{ID: 3, UUID: "uuid3"},
			}

			handler := task.Handler{}
			for _, ts := range tasks {
				handler.AppendTask(ts)
			}

			err := handler.RemoveTasks(tt.ids)
			ts := handler.GetTasks()

			if len(ts) != len(tt.want) {
				t.Fatalf("want %#v, but %#v", tt.want, ts)
			}

			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err.Error())
			}

			if tt.wantError && err.Error() != tt.err.Error() {
				t.Fatalf("want %#v, but %#v", tt.err.Error(), err.Error())
			}

			if !tt.wantError {
				for i := 0; i < len(tt.want); i++ {
					if ts[i].ID != tt.want[i].ID || ts[i].UUID != tt.want[i].UUID {
						t.Fatalf("want %#v, but %#v", tt.want, ts)
					}
				}
			}
		})
	}
}
