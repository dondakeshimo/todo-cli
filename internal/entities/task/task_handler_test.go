package task_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
)

// TODO: UpdateTask, FindTaskWithUUID, AppendTask

func TestGetTask(t *testing.T) {
	tasks := []task.Task{}
	tasks = append(tasks, task.NewTask(1, "", "", "uuid1", "", 1))
	tasks = append(tasks, task.NewTask(2, "", "", "uuid2", "", 2))
	tasks = append(tasks, task.NewTask(3, "", "", "uuid3", "", 3))

	handler := task.Handler{}

	for _, ts := range tasks {
		handler.AppendTask(ts)
	}

	tests := []struct {
		name      string
		id        int
		want      task.Task
		wantError bool
		err       error
	}{
		{
			name:      "Success",
			id:        2,
			want:      tasks[1],
			wantError: false,
			err:       nil,
		},
		{
			name:      "HasErrorOverRange",
			id:        4,
			want:      task.Task{},
			wantError: true,
			err:       errors.New("no exist id: 4"),
		},
		{
			name:      "FailMinus",
			id:        -1,
			want:      task.Task{},
			wantError: true,
			err:       errors.New("no exist id: -1"),
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetTask(tt.id)

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

func TestGetTasks(t *testing.T) {
	tasks := []task.Task{}
	tasks = append(tasks, task.NewTask(1, "", "", "uuid1", "", 1))
	tasks = append(tasks, task.NewTask(2, "", "", "uuid2", "", 2))
	tasks = append(tasks, task.NewTask(3, "", "", "uuid3", "", 3))

	handler := task.Handler{}

	for _, ts := range tasks {
		handler.AppendTask(ts)
	}

	tests := []struct {
		name string
		want []task.Task
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
	tasks := []task.Task{}
	tasks = append(tasks, task.NewTask(1, "", "", "uuid1", "", 1))
	tasks = append(tasks, task.NewTask(2, "", "", "uuid2", "", 2))
	tasks = append(tasks, task.NewTask(3, "", "", "uuid3", "", 3))

	tasksSuccess1 := []task.Task{}
	tasksSuccess1 = append(tasksSuccess1, task.NewTask(1, "", "", "uuid2", "", 2))
	tasksSuccess1 = append(tasksSuccess1, task.NewTask(2, "", "", "uuid3", "", 3))

	tasksSuccess2 := []task.Task{}
	tasksSuccess2 = append(tasksSuccess2, task.NewTask(1, "", "", "uuid1", "", 1))
	tasksSuccess2 = append(tasksSuccess2, task.NewTask(2, "", "", "uuid3", "", 3))

	tests := []struct {
		name      string
		id        int
		want      []task.Task
		wantError bool
		err       error
	}{
		{
			name:      "Success1",
			id:        1,
			want:      tasksSuccess1,
			wantError: false,
			err:       nil,
		},
		{
			name:      "Success2",
			id:        2,
			want:      tasksSuccess2,
			wantError: false,
			err:       nil,
		},
		{
			name:      "FailOverRange",
			id:        4,
			want:      tasks,
			wantError: true,
			err:       errors.New("invalid id [4]"),
		},
		{
			name:      "FailMinus",
			id:        -1,
			want:      tasks,
			wantError: true,
			err:       errors.New("invalid id [-1]"),
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tasks := []task.Task{}
			tasks = append(tasks, task.NewTask(1, "", "", "uuid1", "", 1))
			tasks = append(tasks, task.NewTask(2, "", "", "uuid2", "", 2))
			tasks = append(tasks, task.NewTask(3, "", "", "uuid3", "", 3))

			handler := task.Handler{}
			for _, ts := range tasks {
				handler.AppendTask(ts)
			}

			err := handler.RemoveTask(tt.id)
			ts := handler.GetTasks()

			if !reflect.DeepEqual(ts, tt.want) {
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
					if ts[i].ID() != tt.want[i].ID() || ts[i].UUID() != tt.want[i].UUID() {
						t.Fatalf("want %#v, but %#v", tt.want, ts)
					}
				}
			}
		})
	}
}

func TestRemoveTasks(t *testing.T) {
	tasks := []task.Task{}
	tasks = append(tasks, task.NewTask(1, "", "", "uuid1", "", 1))
	tasks = append(tasks, task.NewTask(2, "", "", "uuid2", "", 2))
	tasks = append(tasks, task.NewTask(3, "", "", "uuid3", "", 3))

	tasksSuccess1 := []task.Task{}
	tasksSuccess1 = append(tasksSuccess1, task.NewTask(1, "", "", "uuid2", "", 2))

	tasksSuccess2 := []task.Task{}
	tasksSuccess2 = append(tasksSuccess2, task.NewTask(1, "", "", "uuid3", "", 3))

	tests := []struct {
		name      string
		ids       []int
		want      []task.Task
		wantError bool
		err       error
	}{
		{
			name:      "Success",
			ids:       []int{1, 3},
			want:      tasksSuccess1,
			wantError: false,
			err:       nil,
		},
		{
			name:      "SuccessDescending",
			ids:       []int{2, 1},
			want:      tasksSuccess2,
			wantError: false,
			err:       nil,
		},
		{
			name:      "HasContinueOverRange",
			ids:       []int{1, 2, 4},
			want:      tasks,
			wantError: true,
			err:       errors.New("no task with id [1 2 4]"),
		},
		{
			name:      "HasContinueMinus",
			ids:       []int{1, -1},
			want:      tasks,
			wantError: true,
			err:       errors.New("not natural value is invalid [-1 1]"),
		},
	}

	for _, tt := range tests {
		tt := tt // set local scope for parallel test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tasks := []task.Task{}
			tasks = append(tasks, task.NewTask(1, "", "", "uuid1", "", 1))
			tasks = append(tasks, task.NewTask(2, "", "", "uuid2", "", 2))
			tasks = append(tasks, task.NewTask(3, "", "", "uuid3", "", 3))

			handler := task.Handler{}
			for _, ts := range tasks {
				handler.AppendTask(ts)
			}

			err := handler.RemoveTasks(tt.ids)
			ts := handler.GetTasks()

			if !reflect.DeepEqual(ts, tt.want) {
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
					if ts[i].ID() != tt.want[i].ID() || ts[i].UUID() != tt.want[i].UUID() {
						t.Fatalf("want %#v, but %#v", tt.want, ts)
					}
				}
			}
		})
	}
}
