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
			want: `+----+--------------------------------+----------------+-------+----------+----------+
| ID |              Task              |   RemindTime   | Group | Reminder | Priority |
+----+--------------------------------+----------------+-------+----------+----------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |       |          |        0 |
|    | task is your first TODO        |                |       |          |          |
+----+--------------------------------+----------------+-------+----------+----------+
`,
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
			name:      "add task 7 in Japanese",
			command:   []string{"a", "シナリオテスト 7"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "add task 8 with group",
			command:   []string{"a", "シナリオテスト 8", "-g", "scenario"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list added tasks",
			command:   []string{"l"},
			hasOutput: true,
            want: `+----+--------------------------------+----------------+----------+----------+----------+
| ID |              Task              |   RemindTime   |  Group   | Reminder | Priority |
+----+--------------------------------+----------------+----------+----------+----------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |          |          |        0 |
|    | task is your first TODO        |                |          |          |          |
|  2 | scenario test 6                |                |          |          |       50 |
|  3 | scenario test 1                |                |          |          |      100 |
|  4 | scenario test 2                | 2099/1/1 00:00 |          | macos    |      100 |
|  5 | シナリオテスト 7               |                |          |          |      100 |
|  6 | シナリオテスト 8               |                | scenario |          |      100 |
+----+--------------------------------+----------------+----------+----------+----------+
`,
			wantError: false,
			err:       "",
		},
		{
			name:      "list filtered tasks",
			command:   []string{"l", "-g", "scenario"},
			hasOutput: true,
            want: `+----+------------------+------------+----------+----------+----------+
| ID |       Task       | RemindTime |  Group   | Reminder | Priority |
+----+------------------+------------+----------+----------+----------+
|  6 | シナリオテスト 8 |            | scenario |          |      100 |
+----+------------------+------------+----------+----------+----------+
`,
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
            want: `+----+--------------------------------+----------------+----------+----------+----------+
| ID |              Task              |   RemindTime   |  Group   | Reminder | Priority |
+----+--------------------------------+----------------+----------+----------+----------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |          |          |        0 |
|    | task is your first TODO        |                |          |          |          |
|  2 | scenario test 6                |                |          |          |       50 |
|  3 | scenario test modified 1       | 2099/1/1 12:00 |          |          |      100 |
|  4 | scenario test 2                | 2099/1/1 00:00 |          | macos    |      100 |
|  5 | シナリオテスト 7               |                |          |          |      100 |
|  6 | シナリオテスト 8               |                | scenario |          |      100 |
+----+--------------------------------+----------------+----------+----------+----------+
`,
			wantError: false,
			err:       "",
		},
		{
			name:      "config Hide",
			command:   []string{"conf", "--hide_group=true", "--hide_reminder=true", "--hide_priority=true"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list hided tasks",
			command:   []string{"l"},
			hasOutput: true,
			want: `+----+--------------------------------+----------------+
| ID |              Task              |   RemindTime   |
+----+--------------------------------+----------------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |
|    | task is your first TODO        |                |
|  2 | scenario test 6                |                |
|  3 | scenario test modified 1       | 2099/1/1 12:00 |
|  4 | scenario test 2                | 2099/1/1 00:00 |
|  5 | シナリオテスト 7               |                |
|  6 | シナリオテスト 8               |                |
+----+--------------------------------+----------------+
`,
			wantError: false,
			err:       "",
		},
		{
			name:      "reset config",
			command:   []string{"conf", "--reset_config=true"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list not hided tasks",
			command:   []string{"l"},
			hasOutput: true,
            want: `+----+--------------------------------+----------------+----------+----------+----------+
| ID |              Task              |   RemindTime   |  Group   | Reminder | Priority |
+----+--------------------------------+----------------+----------+----------+----------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |          |          |        0 |
|    | task is your first TODO        |                |          |          |          |
|  2 | scenario test 6                |                |          |          |       50 |
|  3 | scenario test modified 1       | 2099/1/1 12:00 |          |          |      100 |
|  4 | scenario test 2                | 2099/1/1 00:00 |          | macos    |      100 |
|  5 | シナリオテスト 7               |                |          |          |      100 |
|  6 | シナリオテスト 8               |                | scenario |          |      100 |
+----+--------------------------------+----------------+----------+----------+----------+
`,
			wantError: false,
			err:       "",
		},
		{
			name:      "clear tasks",
			command:   []string{"c", "-i=1,3,5"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "clear task 2",
			command:   []string{"c", "-i=1", "-i=2,3"},
			hasOutput: false,
			want:      "",
			wantError: false,
			err:       "",
		},
		{
			name:      "list modified tasks",
			command:   []string{"l"},
			hasOutput: true,
            want: `+----+------+------------+-------+----------+----------+
| ID | Task | RemindTime | Group | Reminder | Priority |
+----+------+------------+-------+----------+----------+
+----+------+------------+-------+----------+----------+
`,
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
				t.Fatalf("want %#v, but %#v", sc.err, err.Error())
			}

			if sc.hasOutput {
				if string(got) != sc.want {
                    t.Log("want")
                    t.Log(sc.want)
                    t.Log("got")
                    t.Fatal(string(got))
				}
			}
		})
	}
}
