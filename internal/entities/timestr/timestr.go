package timestr

import (
	"fmt"
	"time"
)

const (
	layoutMin = "2006/1/2 15:04"
	layoutDay = "2006/1/2"
)

// UnifyLayout is a function that validates input and returns layoutMin type object.
// Allowed layouts are
//     "2006/1/2 15:04"
//     "2006/1/2"
// Transfer "2006/1/2" to "2006/1/2 15:04".
func UnifyLayout(str string) (string, error) {
	tm, errM := time.ParseInLocation(layoutMin, str, time.Local)
	td, errD := time.ParseInLocation(layoutDay, str, time.Local)

	if errM != nil && errD != nil {
		return "", fmt.Errorf("invalid time layout: [minutes layout]: %s, [day layout]: %s", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return tm.Format(layoutMin), nil
	} else if errM != nil && errD == nil {
		return td.Format(layoutMin), nil
	}

	return "", fmt.Errorf("Parse failed for some reason")
}

// FormatTime is a function that formats time.
func FormatTime(t time.Time) string {
	return t.Format(layoutMin)
}

// ModifyTime is a function that add duration to base time.
func ModifyTime(duration string, base time.Time) (string, error) {
	rt, err := time.ParseDuration(duration)

	if err != nil {
		return "", err
	}

	return base.Add(rt).Format(layoutMin), nil
}

// Parse is a function that parse time.
// Allowed layouts are
//     "2006/1/2 15:04"
//     "2006/1/2"
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
