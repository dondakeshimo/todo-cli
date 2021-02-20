package commands

import (
	"strings"
	"time"

	"github.com/dondakeshimo/todo-cli/internal/entities/timestr"
)

func arrangeRemindTime(rt string, preRt string) (string, error) {
	if rt == "" {
		return preRt, nil
	}

	// when input layout is "task+1h3m" or "task-1h3m"
	if preRt != "" && strings.HasPrefix(rt, "task") {
		rt = strings.Replace(rt, "task", "", 1)
		pt, err := timestr.Parse(preRt)
		if err != nil {
			return "", err
		}
		rt, err = timestr.ModifyTime(rt, pt)
		if err != nil {
			return "", err
		}
	}

	// when input layout is "+1h3m" or "now+1h3m"
	if strings.HasPrefix(rt, "+") || strings.HasPrefix(rt, "now+") {
		rt = strings.Replace(rt, "now", "", 1)
		var err error
		rt, err = timestr.ModifyTime(rt, time.Now())
		if err != nil {
			return "", err
		}
	}

	return timestr.UnifyLayout(rt)
}
