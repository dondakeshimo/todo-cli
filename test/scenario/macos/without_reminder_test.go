package scenario_test

import (
	"os/exec"
	"testing"
)

func TestWithoutReminder(t *testing.T) {
	// TODO: refactor this awful spaghetti code
	got, err := exec.Command("todo", "l").Output()
	want := "ID  Task                                                RemindTime        reminder\n1   deleting or modifying this task is your first TODO  2099/01/01 00:00  \n"
	if err != nil {
		t.Fatalf("todo l has error: %#v", err.Error())
	}
	if string(got) != want {
		t.Fatalf("want %q, but %q", want, got)
	}

	err = exec.Command("todo", "a", "-t", "scenario test 1").Run()
	if err != nil {
		t.Fatalf("todo a 1 has error: %#v", err.Error())
	}

	err = exec.Command("todo", "a", "-t", "scenario test 2", "-r", "macos", "-d", "2099/01/01 00:00").Run()
	if err != nil {
		t.Fatalf("todo a 2 has error: %#v", err.Error())
	}

	err = exec.Command("todo", "a", "-t", "scenario test 3", "-r", "reminder invalid", "-d", "2099/01/01 00:00").Run()
	wantErr := "exit status 1"
	if err.Error() != wantErr {
		t.Fatalf("want %#v, but %#v", wantErr, err.Error())
	}

	err = exec.Command("todo", "a", "-t", "scenario test 4", "-r", "reminder invalid", "-d", "2099/01/hoge 00:00").Run()
	wantErr = "exit status 1"
	if err.Error() != wantErr {
		t.Fatalf("want %#v, but %#v", wantErr, err.Error())
	}

	err = exec.Command("todo", "a", "-z", "scenario test 5", "-r", "reminder invalid", "-d", "2099/01/hoge 00:00").Run()
	wantErr = "exit status 1"
	if err.Error() != wantErr {
		t.Fatalf("want %#v, but %#v", wantErr, err.Error())
	}

	got, err = exec.Command("todo", "l").Output()
	want = "ID  Task                                                RemindTime        reminder\n1   deleting or modifying this task is your first TODO  2099/01/01 00:00  \n2   scenario test 1                                                       \n3   scenario test 2                                     2099/01/01 00:00  macos\n"
	if err != nil {
		t.Fatalf("todo l has error: %#v", err.Error())
	}
	if string(got) != want {
		t.Fatalf("want %q, but %q", want, got)
	}

	err = exec.Command("todo", "m", "-t", "scenario test modify 1", "-i", "2", "-d", "2099/01/01 12:00").Run()
	if err != nil {
		t.Fatalf("todo m 1 has error: %#v", err.Error())
	}

	got, err = exec.Command("todo", "l").Output()
	want = "ID  Task                                                RemindTime        reminder\n1   deleting or modifying this task is your first TODO  2099/01/01 00:00  \n2   scenario test modify 1                              2099/01/01 12:00  \n3   scenario test 2                                     2099/01/01 00:00  macos\n"
	if err != nil {
		t.Fatalf("todo l has error: %#v", err.Error())
	}
	if string(got) != want {
		t.Fatalf("want %q, but %q", want, got)
	}

	err = exec.Command("todo", "c", "-i", "1", "-i", "3").Run()
	if err != nil {
		t.Fatalf("todo c 1 has error: %#v", err.Error())
	}

	err = exec.Command("todo", "c", "-i", "1").Run()
	if err != nil {
		t.Fatalf("todo c 2 has error: %#v", err.Error())
	}

	got, err = exec.Command("todo", "l").Output()
	want = "ID  Task  RemindTime  reminder\n"
	if err != nil {
		t.Fatalf("todo l has error: %#v", err.Error())
	}
	if string(got) != want {
		t.Fatalf("want %q, but %q", want, got)
	}
}
