package notifier

import (
	"fmt"
	"os/exec"
)

type OsascriptNotifier struct {}

func (on *OsascriptNotifier) Push (r *Request) error {
	c := fmt.Sprintf("display dialog \"%s\" buttons [\"done\"] with title \"%s\"", r.Contents, r.Title)

	if err := exec.Command("osascript", "-e", c).Run(); err != nil {
		return err
	}

	return nil
}
