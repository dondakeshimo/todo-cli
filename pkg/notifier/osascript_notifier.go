package notifier

import (
	"fmt"
	"os/exec"
	"strings"
)

type OsascriptNotifier struct{}

func (on *OsascriptNotifier) Push(r *Request) (string, error) {
	out, err := exec.Command("osascript", "-e", buildScript(r)).Output()

	if err != nil {
		return "", err
	}

	o := strings.Replace(string(out), "button returned:", "", 1)
	o = strings.Trim(o, " ")
	o = strings.Trim(o, "\n")

	return o, nil
}

func buildScript(r *Request) string {
	const baseScript = "display dialog \"%s\" buttons [%s] with title \"%s\""

	as := make([]string, len(r.Answer))
	for i, a := range r.Answer {
		as[i] = "\"" + a + "\""
	}

	ans := strings.Join(as, ",")

	return fmt.Sprintf(baseScript, r.Contents, ans, r.Title)
}
