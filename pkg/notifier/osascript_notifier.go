package notifier

import (
	"fmt"
	"os/exec"
	"strings"
)

type OsascriptNotifier struct{}

func (on *OsascriptNotifier) Push(r *Request) (string, error) {
	c := fmt.Sprintf("display dialog \"%s\" buttons [\"skip\",\"done\"] with title \"%s\"", r.Contents, r.Title)

	out, err := exec.Command("osascript", "-e", c).Output()

	if err != nil {
		return "", err
	}

	o := strings.Replace(string(out), "button returned:", "", 1)
	o = strings.Trim(o, " ")
	o = strings.Trim(o, "\n")

	return o, nil
}
