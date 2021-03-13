// +build scenario

package scenario_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestWithoutReminder(t *testing.T) {
	const binaryName = "todo"
	const binaryDir = "../../"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working directory: %s", err.Error())
	}

	bin := filepath.Join(dir, binaryDir, binaryName)

	scenarios := []struct {
		name      string
		command   []string
		hasOutput bool
		want      string
		wantError bool
		err       string
	}{
		{
			name:      "list initial sample",
			command:   []string{"l"},
			hasOutput: true,
			want:      "ID  Task                                                RemindTime      reminder  Priority\n1   deleting or modifying this task is your first TODO  2099/1/1 00:00            0\n",
			wantError: false,
			err:       "",
		},
		{
			name:      "add task 1",
			command:   []string{"a", "scenario test 1"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "add task 2",
			command:   []string{"a", "-r", "macos", "-d", "2099/01/01", "scenario test 2"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "not add task 3",
			command:   []string{"a", "-r", "invalid", "-d", "2099/01/01", "scenario test 3"},
			hasOutput: false,
			want:      "",
			wantError: true,
			err:       "exit status 1",
		},
		{
			name:      "not add task 4",
			command:   []string{"a", "-r", "macos", "-d", "2099/01/hoge", "scenario test 4"},
			hasOutput: false,
			want:      "",
			wantError: true,
			err:       "exit status 1",
		},
		{
			name:      "not add task 5",
			command:   []string{"a"},
			hasOutput: false,
			want:      "",
			wantError: true,
			err:       "exit status 1",
		},
		{
			name:      "add task 6 with priority",
			command:   []string{"a", "-p", "50", "scenario test 6"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list added tasks",
			command:   []string{"l"},
			hasOutput: true,
			want:      "ID  Task                                                RemindTime      reminder  Priority\n1   deleting or modifying this task is your first TODO  2099/1/1 00:00            0\n2   scenario test 6                                                               50\n3   scenario test 1                                                               100\n4   scenario test 2                                     2099/1/1 00:00  macos     100\n",
			wantError: false,
			err:       "",
		},
		{
			name:      "modify task 1",
			command:   []string{"m", "-t", "scenario test modified 1", "-i", "3", "-d", "2099/01/01 12:00"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list modified tasks",
			command:   []string{"l"},
			hasOutput: true,
			want:      "ID  Task                                                RemindTime      reminder  Priority\n1   deleting or modifying this task is your first TODO  2099/1/1 00:00            0\n2   scenario test 6                                                               50\n3   scenario test modified 1                            2099/1/1 12:00            100\n4   scenario test 2                                     2099/1/1 00:00  macos     100\n",
			wantError: false,
			err:       "",
		},
		{
			name:      "clear tasks",
			command:   []string{"c", "-i=1", "-i=3", "-i=4"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "clear task 2",
			command:   []string{"c", "-i=1"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list modified tasks",
			command:   []string{"l"},
			hasOutput: true,
			want:      "ID  Task  RemindTime  reminder  Priority\n",
			wantError: false,
			err:       "",
		},
	}

	for i, sc := range scenarios {
		t.Run(fmt.Sprintf("[%d] %s", i, sc.name), func(t *testing.T) {
			var got []byte
			var err error

			if sc.hasOutput {
				got, err = exec.Command(bin, sc.command...).Output()
			} else {
				err = exec.Command(bin, sc.command...).Run()
			}

			if !sc.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err.Error())
			}

			if sc.wantError && err.Error() != sc.err {
				fmt.Println(11)
				t.Fatalf("want %#v, but %#v", sc.err, err.Error())
			}

			if sc.hasOutput {
				if string(got) != sc.want {
					t.Fatalf("want %#v, but %#v", sc.want, string(got))
				}
			}
		})
	}
}
