package timestr

import (
	"fmt"
	"time"
)

const (
	layoutMin = "2006/01/02 15:04"
	layoutDay = "2006/01/02"
)

func Parse(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	_, errM := time.Parse(layoutMin, str)
	_, errD := time.Parse(layoutDay, str)

	if errM != nil && errD != nil {
		return "", fmt.Errorf("invalid layout: [%s, %s]", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return str, nil
	} else if errM != nil && errD == nil {
		return str + " 00:00", nil
	}

	return "", fmt.Errorf("Parse failed for some reason")
}