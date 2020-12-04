package task_test

import (
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
