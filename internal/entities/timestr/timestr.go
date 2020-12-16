package timestr

import (
	"fmt"
	"time"
)

const (
	layoutMin = "2006/1/2 15:4"
	layoutDay = "2006/1/2"
)

// Validate is a function that confirm time valid.
func Validate(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	_, errM := time.ParseInLocation(layoutMin, str, time.Local)
	_, errD := time.ParseInLocation(layoutDay, str, time.Local)

	if errM != nil && errD != nil {
		return "", fmt.Errorf("invalid time layout: [minutes layout]: %s, [day layout]: %s", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return str, nil
	} else if errM != nil && errD == nil {
		return str + " 00:00", nil
	}

	return "", fmt.Errorf("Parse failed for some reason")
}

// Parse is a function that parse time.
func Parse(str string) (time.Time, error) {
	tM, errM := time.ParseInLocation(layoutMin, str, time.Local)
	tD, errD := time.ParseInLocation(layoutDay, str, time.Local)

	if errM != nil && errD != nil {
		return time.Time{}, fmt.Errorf("invalid time layout: [minutes layout]: %s, [day layout]: %s", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return tM, nil
	} else if errM != nil && errD == nil {
		return tD, nil
	}

	return time.Time{}, fmt.Errorf("Parse failed for some reason")
}
